package mapper

import (
	"fmt"
	"sort"
	"strings"

	"regexp"

	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/convert"
	"github.com/rancher/norman/types/values"
	"k8s.io/api/core/v1"
	v1types "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	exprRegexp = regexp.MustCompile("^(.*?)\\s*(=|!=|<|>| in | notin )\\s*(.*)$")
)

type SchedulingMapper struct {
}

func (s SchedulingMapper) FromInternal(data map[string]interface{}) {
	defer func() {
		delete(data, "nodeSelector")
		delete(data, "affinity")
	}()

	var requireAll []string
	for key, value := range convert.ToMapInterface(data["nodeSelector"]) {
		if value == "" {
			requireAll = append(requireAll, key)
		} else {
			requireAll = append(requireAll, fmt.Sprintf("%s = %s", key, value))
		}
	}

	if len(requireAll) > 0 {
		values.PutValue(data, requireAll, "scheduling", "node", "requireAll")
	}

	v, ok := data["affinity"]
	if !ok || v == nil {
		return
	}

	affinity := &v1.Affinity{}
	if err := convert.ToObj(v, affinity); err != nil {
		return
	}

	if affinity.NodeAffinity != nil {
		s.nodeAffinity(data, affinity.NodeAffinity)
	}

	if affinity.PodAffinity != nil {
		s.podAffinity(data, affinity.PodAffinity)
	}

	if affinity.PodAntiAffinity != nil {
		s.podAntiAffinity(data, affinity.PodAntiAffinity)
	}
}

func (s SchedulingMapper) nodeAffinity(data map[string]interface{}, nodeAffinity *v1.NodeAffinity) {
	var requireAll []string
	var requireAny []string
	var preferred []string

	if nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution != nil {
		for _, term := range nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms {
			exprs := NodeSelectorTermToStrings(term)
			if len(exprs) == 0 {
				continue
			}
			if len(requireAny) > 0 {
				// Once any is set all new terms go to any
				requireAny = append(requireAny, strings.Join(exprs, " && "))
			} else if len(requireAll) > 0 {
				// If all is already set, we actually need to move everything to any
				requireAny = append(requireAny, strings.Join(requireAll, " && "))
				requireAny = append(requireAny, strings.Join(exprs, " && "))
				requireAll = []string{}
			} else {
				// The first term is considered all
				requireAll = exprs
			}
		}
	}

	if nodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution != nil {
		sortPreferred(nodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution)
		for _, term := range nodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
			exprs := NodeSelectorTermToStrings(term.Preference)
			preferred = append(preferred, strings.Join(exprs, " && "))
		}
	}

	if len(requireAll) > 0 {
		values.PutValue(data, requireAll, "scheduling", "node", "requireAll")
	}
	if len(requireAny) > 0 {
		values.PutValue(data, requireAny, "scheduling", "node", "requireAny")
	}
	if len(preferred) > 0 {
		values.PutValue(data, preferred, "scheduling", "node", "preferred")
	}
}

func (s SchedulingMapper) podAffinity(data map[string]interface{}, podAffinity *v1.PodAffinity) {
	podAffinityPut(data, podAffinity.RequiredDuringSchedulingIgnoredDuringExecution, podAffinity.PreferredDuringSchedulingIgnoredDuringExecution, "podAffinity")
}

func (s SchedulingMapper) podAntiAffinity(data map[string]interface{}, podAntiAffinity *v1.PodAntiAffinity) {
	podAffinityPut(data, podAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution, podAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution, "podAntiAffinity")
}

func podAffinityPut(data map[string]interface{}, podAffinityTerms []v1.PodAffinityTerm, weightedPodAffinityTerms []v1.WeightedPodAffinityTerm, affinityType string) {
	var required []interface{}
	var preferred []interface{}

	if podAffinityTerms != nil {
		for _, term := range podAffinityTerms {
			expressions := PodAffinityTermToStrings(term)
			required = append(required, map[string]interface{}{
				"expressions": expressions,
				"namespaces":  term.Namespaces,
				"topology":    term.TopologyKey,
			})
		}
	}

	if weightedPodAffinityTerms != nil {
		for _, weightTerm := range weightedPodAffinityTerms {
			term := weightTerm.PodAffinityTerm
			expressions := PodAffinityTermToStrings(term)
			preferred = append(preferred, map[string]interface{}{
				"expressions": expressions,
				"namespaces":  term.Namespaces,
				"topology":    term.TopologyKey,
			})
		}
	}

	if len(required) > 0 {
		values.PutValue(data, required, "scheduling", affinityType, "required")
	}

	if len(preferred) > 0 {
		values.PutValue(data, preferred, "scheduling", affinityType, "preferred")
	}
}

func PodAffinityTermToStrings(term v1.PodAffinityTerm) []string {
	var exprs []string

	if term.LabelSelector != nil {
		if len(term.LabelSelector.MatchExpressions) > 0 {
			for _, expr := range term.LabelSelector.MatchExpressions {
				nextExpr := ""
				switch expr.Operator {
				case v1types.LabelSelectorOpIn:
					if len(expr.Values) > 1 {
						nextExpr = fmt.Sprintf("%s in (%s)", expr.Key, strings.Join(expr.Values, ", "))
					} else if len(expr.Values) == 1 {
						nextExpr = fmt.Sprintf("%s = %s", expr.Key, expr.Values[0])
					}
				case v1types.LabelSelectorOpNotIn:
					if len(expr.Values) > 1 {
						nextExpr = fmt.Sprintf("%s notin (%s)", expr.Key, strings.Join(expr.Values, ", "))
					} else if len(expr.Values) == 1 {
						nextExpr = fmt.Sprintf("%s != %s", expr.Key, expr.Values[0])
					}
				case v1types.LabelSelectorOpExists:
					nextExpr = expr.Key
				case v1types.LabelSelectorOpDoesNotExist:
					nextExpr = "!" + expr.Key
				}
				if nextExpr != "" {
					exprs = append(exprs, nextExpr)
				}
			}
		}

		if term.LabelSelector.MatchLabels != nil {
			for k, v := range term.LabelSelector.MatchLabels {
				exprs = append(exprs, fmt.Sprintf("%s = %s", k, v))
			}
		}
	}

	return exprs
}

func sortPreferred(terms []v1.PreferredSchedulingTerm) {
	sort.Slice(terms, func(i, j int) bool {
		return terms[i].Weight > terms[j].Weight
	})
}

func NodeSelectorTermToStrings(term v1.NodeSelectorTerm) []string {
	exprs := []string{}

	for _, expr := range term.MatchExpressions {
		nextExpr := ""
		switch expr.Operator {
		case v1.NodeSelectorOpIn:
			if len(expr.Values) > 1 {
				nextExpr = fmt.Sprintf("%s in (%s)", expr.Key, strings.Join(expr.Values, ", "))
			} else if len(expr.Values) == 1 {
				nextExpr = fmt.Sprintf("%s = %s", expr.Key, expr.Values[0])
			}
		case v1.NodeSelectorOpNotIn:
			if len(expr.Values) > 1 {
				nextExpr = fmt.Sprintf("%s notin (%s)", expr.Key, strings.Join(expr.Values, ", "))
			} else if len(expr.Values) == 1 {
				nextExpr = fmt.Sprintf("%s != %s", expr.Key, expr.Values[0])
			}
		case v1.NodeSelectorOpExists:
			nextExpr = expr.Key
		case v1.NodeSelectorOpDoesNotExist:
			nextExpr = "!" + expr.Key
		case v1.NodeSelectorOpGt:
			if len(expr.Values) == 1 {
				nextExpr = fmt.Sprintf("%s > %s", expr.Key, expr.Values[0])
			}
		case v1.NodeSelectorOpLt:
			if len(expr.Values) == 1 {
				nextExpr = fmt.Sprintf("%s < %s", expr.Key, expr.Values[0])
			}
		}

		if nextExpr != "" {
			exprs = append(exprs, nextExpr)
		}
	}

	return exprs
}

func StringsToNodeSelectorTerm(exprs []string) []v1.NodeSelectorTerm {
	result := []v1.NodeSelectorTerm{}

	for _, inter := range exprs {
		term := v1.NodeSelectorTerm{}

		for _, expr := range strings.Split(inter, "&&") {
			groups := exprRegexp.FindStringSubmatch(expr)
			selectorRequirement := v1.NodeSelectorRequirement{}

			if groups == nil {
				if strings.HasPrefix(expr, "!") {
					selectorRequirement.Key = strings.TrimSpace(expr[1:])
					selectorRequirement.Operator = v1.NodeSelectorOpDoesNotExist
				} else {
					selectorRequirement.Key = strings.TrimSpace(expr)
					selectorRequirement.Operator = v1.NodeSelectorOpExists
				}
			} else {
				selectorRequirement.Key = strings.TrimSpace(groups[1])
				selectorRequirement.Values = convert.ToValuesSlice(groups[3])
				op := strings.TrimSpace(groups[2])
				switch op {
				case "=":
					selectorRequirement.Operator = v1.NodeSelectorOpIn
				case "!=":
					selectorRequirement.Operator = v1.NodeSelectorOpNotIn
				case "notin":
					selectorRequirement.Operator = v1.NodeSelectorOpNotIn
				case "in":
					selectorRequirement.Operator = v1.NodeSelectorOpIn
				case "<":
					selectorRequirement.Operator = v1.NodeSelectorOpLt
				case ">":
					selectorRequirement.Operator = v1.NodeSelectorOpGt
				}
			}

			term.MatchExpressions = append(term.MatchExpressions, selectorRequirement)
		}

		result = append(result, term)
	}

	return result
}

func (s SchedulingMapper) ToInternal(data map[string]interface{}) error {
	defer func() {
		delete(data, "scheduling")
	}()

	nodeName := convert.ToString(values.GetValueN(data, "scheduling", "node", "nodeId"))
	if nodeName != "" {
		data["nodeName"] = nodeName
	}

	requireAllV := values.GetValueN(data, "scheduling", "node", "requireAll")
	requireAnyV := values.GetValueN(data, "scheduling", "node", "requireAny")
	preferredV := values.GetValueN(data, "scheduling", "node", "preferred")
	podAffinityRequiredV := values.GetValueN(data, "scheduling", "podAffinity", "required")
	podAffinityPreferredV := values.GetValueN(data, "scheduling", "podAffinity", "preferred")
	podAntiAffinityRequiredV := values.GetValueN(data, "scheduling", "podAntiAffinity", "required")
	podAntiAffinityPreferredV := values.GetValueN(data, "scheduling", "podAntiAffinity", "preferred")

	if requireAllV == nil && requireAnyV == nil && preferredV == nil && podAffinityRequiredV == nil && podAffinityPreferredV == nil && podAntiAffinityRequiredV == nil && podAntiAffinityPreferredV == nil {
		return nil
	}

	requireAll := convert.ToStringSlice(requireAllV)
	requireAny := convert.ToStringSlice(requireAnyV)
	preferred := convert.ToStringSlice(preferredV)
	podAffinityRequired := convert.ToMapSlice(podAffinityRequiredV)
	podAffinityPreferred := convert.ToMapSlice(podAffinityPreferredV)
	podAntiAffinityRequired := convert.ToMapSlice(podAntiAffinityRequiredV)
	podAntiAffinityPreferred := convert.ToMapSlice(podAntiAffinityPreferredV)

	affinityNil := true
	if len(requireAll) == 0 && len(requireAny) == 0 && len(preferred) == 0 {
		values.PutValue(data, nil, "affinity", "nodeAffinity")
	} else {
		affinityNil = false
	}

	if len(podAffinityRequired) == 0 && len(podAffinityPreferred) == 0 {
		values.PutValue(data, nil, "affinity", "podAffinity")
	} else {
		affinityNil = false
	}

	if len(podAntiAffinityRequired) == 0 && len(podAntiAffinityPreferred) == 0 {
		values.PutValue(data, nil, "affinity", "podAntiAffinity")
	} else {
		affinityNil = false
	}

	if affinityNil {
		return nil
	}

	nodeAffinity := v1.NodeAffinity{}
	podAffinity := v1.PodAffinity{}
	podAntiAffinity := v1.PodAntiAffinity{}

	if len(requireAll) > 0 {
		nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution = &v1.NodeSelector{
			NodeSelectorTerms: []v1.NodeSelectorTerm{
				AggregateTerms(StringsToNodeSelectorTerm(requireAll)),
			},
		}
	}

	if len(requireAny) > 0 {
		if nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution == nil {
			nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution = &v1.NodeSelector{}
		}
		nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms = append(nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms, StringsToNodeSelectorTerm(requireAny)...)
	}

	if len(preferred) > 0 {
		count := int32(100)
		for _, term := range StringsToNodeSelectorTerm(preferred) {
			nodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution = append(
				nodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution, v1.PreferredSchedulingTerm{
					Weight:     count,
					Preference: term,
				})
			count--
		}
	}

	if len(podAffinityRequired) > 0 {
		for _, required := range podAffinityRequired {
			podAffinity.RequiredDuringSchedulingIgnoredDuringExecution = append(podAffinity.RequiredDuringSchedulingIgnoredDuringExecution, GetPodAffinityTerm(required))
		}
	}

	if len(podAffinityPreferred) > 0 {
		for _, preferred := range podAffinityPreferred {
			podAffinity.PreferredDuringSchedulingIgnoredDuringExecution = append(podAffinity.PreferredDuringSchedulingIgnoredDuringExecution, GetWeightedPodAffinityTerm(preferred))
		}
	}

	if len(podAntiAffinityRequired) > 0 {
		for _, required := range podAntiAffinityRequired {
			podAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution = append(podAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution, GetPodAffinityTerm(required))
		}
	}

	if len(podAntiAffinityPreferred) > 0 {
		for _, preferred := range podAntiAffinityPreferred {
			podAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution = append(podAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution, GetWeightedPodAffinityTerm(preferred))
		}
	}

	affinity, _ := convert.EncodeToMap(&v1.Affinity{
		NodeAffinity:    &nodeAffinity,
		PodAffinity:     &podAffinity,
		PodAntiAffinity: &podAntiAffinity,
	})

	if nodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution == nil {
		values.PutValue(affinity, nil, "nodeAffinity", "preferredDuringSchedulingIgnoredDuringExecution")
	}

	if nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution == nil {
		values.PutValue(affinity, nil, "nodeAffinity", "requiredDuringSchedulingIgnoredDuringExecution")
	}

	if podAffinity.PreferredDuringSchedulingIgnoredDuringExecution == nil {
		values.PutValue(affinity, nil, "podAffinity", "preferredDuringSchedulingIgnoredDuringExecution")
	}

	if podAffinity.RequiredDuringSchedulingIgnoredDuringExecution == nil {
		values.PutValue(affinity, nil, "podAffinity", "requiredDuringSchedulingIgnoredDuringExecution")
	}

	if podAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution == nil {
		values.PutValue(affinity, nil, "podAntiAffinity", "preferredDuringSchedulingIgnoredDuringExecution")
	}

	if podAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution == nil {
		values.PutValue(affinity, nil, "podAntiAffinity", "requiredDuringSchedulingIgnoredDuringExecution")
	}

	data["affinity"] = affinity

	return nil
}

func GetWeightedPodAffinityTerm(podAffinityTerm interface{}) v1.WeightedPodAffinityTerm {
	return v1.WeightedPodAffinityTerm{
		Weight:          int32(100),
		PodAffinityTerm: GetPodAffinityTerm(podAffinityTerm),
	}
}

func GetPodAffinityTerm(podAffinityTerm interface{}) v1.PodAffinityTerm {
	expressions := convert.ToStringSlice(values.GetValueN(convert.ToMapInterface(podAffinityTerm), "expressions"))
	namespaces := convert.ToStringSlice(values.GetValueN(convert.ToMapInterface(podAffinityTerm), "namespaces"))
	topology := convert.ToString(values.GetValueN(convert.ToMapInterface(podAffinityTerm), "topology"))

	var term v1.PodAffinityTerm
	if expressions != nil && len(expressions) > 0 {
		term.LabelSelector = &v1types.LabelSelector{
			MatchExpressions: StringsToLabelSelectorTerm(expressions),
		}
	}

	if len(namespaces) > 0 {
		term.Namespaces = namespaces
	}

	if topology == "" {
		topology = "kubernetes.io/hostname"
	}

	term.TopologyKey = topology

	return term
}

func StringsToLabelSelectorTerm(exprs []string) []v1types.LabelSelectorRequirement {
	var result []v1types.LabelSelectorRequirement

	for _, inter := range exprs {
		for _, expr := range strings.Split(inter, "&&") {
			term := v1types.LabelSelectorRequirement{}
			groups := exprRegexp.FindStringSubmatch(expr)
			if groups == nil {
				if strings.HasPrefix(expr, "!") {
					term.Key = strings.TrimSpace(expr[1:])
					term.Operator = v1types.LabelSelectorOpDoesNotExist
				} else {
					term.Key = strings.TrimSpace(expr)
					term.Operator = v1types.LabelSelectorOpExists
				}
			} else {
				term.Key = strings.TrimSpace(groups[1])
				term.Values = convert.ToValuesSlice(groups[3])
				op := strings.TrimSpace(groups[2])
				switch op {
				case "=":
					term.Operator = v1types.LabelSelectorOpIn
				case "!=":
					term.Operator = v1types.LabelSelectorOpNotIn
				case "notin":
					term.Operator = v1types.LabelSelectorOpNotIn
				case "in":
					term.Operator = v1types.LabelSelectorOpIn
				}
			}
			result = append(result, term)
		}
	}
	return result
}

func AggregateTerms(terms []v1.NodeSelectorTerm) v1.NodeSelectorTerm {
	result := v1.NodeSelectorTerm{}
	for _, term := range terms {
		result.MatchExpressions = append(result.MatchExpressions, term.MatchExpressions...)
	}
	return result
}

func (s SchedulingMapper) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	delete(schema.ResourceFields, "nodeSelector")
	delete(schema.ResourceFields, "affinity")
	return nil
}
