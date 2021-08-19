package client

const (
	SchedulingType                   = "scheduling"
	SchedulingFieldNode              = "node"
	SchedulingFieldPodAffinity       = "podAffinity"
	SchedulingFieldPodAntiAffinity   = "podAntiAffinity"
	SchedulingFieldPriority          = "priority"
	SchedulingFieldPriorityClassName = "priorityClassName"
	SchedulingFieldScheduler         = "scheduler"
	SchedulingFieldTolerate          = "tolerate"
)

type Scheduling struct {
	Node              *NodeScheduling  `json:"node,omitempty" yaml:"node,omitempty"`
	PodAffinity       *PodAffinityRule `json:"podAffinity,omitempty" yaml:"podAffinity,omitempty"`
	PodAntiAffinity   *PodAffinityRule `json:"podAntiAffinity,omitempty" yaml:"podAntiAffinity,omitempty"`
	Priority          *int64           `json:"priority,omitempty" yaml:"priority,omitempty"`
	PriorityClassName string           `json:"priorityClassName,omitempty" yaml:"priorityClassName,omitempty"`
	Scheduler         string           `json:"scheduler,omitempty" yaml:"scheduler,omitempty"`
	Tolerate          []Toleration     `json:"tolerate,omitempty" yaml:"tolerate,omitempty"`
}
