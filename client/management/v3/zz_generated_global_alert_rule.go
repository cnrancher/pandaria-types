package client

import (
	"github.com/rancher/norman/types"
)

const (
	GlobalAlertRuleType                       = "globalAlertRule"
	GlobalAlertRuleFieldAlertState            = "alertState"
	GlobalAlertRuleFieldAnnotations           = "annotations"
	GlobalAlertRuleFieldCreated               = "created"
	GlobalAlertRuleFieldCreatorID             = "creatorId"
	GlobalAlertRuleFieldGroupID               = "groupId"
	GlobalAlertRuleFieldGroupIntervalSeconds  = "groupIntervalSeconds"
	GlobalAlertRuleFieldGroupWaitSeconds      = "groupWaitSeconds"
	GlobalAlertRuleFieldInherited             = "inherited"
	GlobalAlertRuleFieldLabels                = "labels"
	GlobalAlertRuleFieldMetricRule            = "metricRule"
	GlobalAlertRuleFieldName                  = "name"
	GlobalAlertRuleFieldOwnerReferences       = "ownerReferences"
	GlobalAlertRuleFieldRemoved               = "removed"
	GlobalAlertRuleFieldRepeatIntervalSeconds = "repeatIntervalSeconds"
	GlobalAlertRuleFieldSeverity              = "severity"
	GlobalAlertRuleFieldState                 = "state"
	GlobalAlertRuleFieldTransitioning         = "transitioning"
	GlobalAlertRuleFieldTransitioningMessage  = "transitioningMessage"
	GlobalAlertRuleFieldUUID                  = "uuid"
)

type GlobalAlertRule struct {
	types.Resource
	AlertState            string            `json:"alertState,omitempty" yaml:"alertState,omitempty"`
	Annotations           map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created               string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID             string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	GroupID               string            `json:"groupId,omitempty" yaml:"groupId,omitempty"`
	GroupIntervalSeconds  int64             `json:"groupIntervalSeconds,omitempty" yaml:"groupIntervalSeconds,omitempty"`
	GroupWaitSeconds      int64             `json:"groupWaitSeconds,omitempty" yaml:"groupWaitSeconds,omitempty"`
	Inherited             *bool             `json:"inherited,omitempty" yaml:"inherited,omitempty"`
	Labels                map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	MetricRule            *MetricRule       `json:"metricRule,omitempty" yaml:"metricRule,omitempty"`
	Name                  string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences       []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed               string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	RepeatIntervalSeconds int64             `json:"repeatIntervalSeconds,omitempty" yaml:"repeatIntervalSeconds,omitempty"`
	Severity              string            `json:"severity,omitempty" yaml:"severity,omitempty"`
	State                 string            `json:"state,omitempty" yaml:"state,omitempty"`
	Transitioning         string            `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage  string            `json:"transitioningMessage,omitempty" yaml:"transitioningMessage,omitempty"`
	UUID                  string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type GlobalAlertRuleCollection struct {
	types.Collection
	Data   []GlobalAlertRule `json:"data,omitempty"`
	client *GlobalAlertRuleClient
}

type GlobalAlertRuleClient struct {
	apiClient *Client
}

type GlobalAlertRuleOperations interface {
	List(opts *types.ListOpts) (*GlobalAlertRuleCollection, error)
	Create(opts *GlobalAlertRule) (*GlobalAlertRule, error)
	Update(existing *GlobalAlertRule, updates interface{}) (*GlobalAlertRule, error)
	Replace(existing *GlobalAlertRule) (*GlobalAlertRule, error)
	ByID(id string) (*GlobalAlertRule, error)
	Delete(container *GlobalAlertRule) error

	ActionActivate(resource *GlobalAlertRule) error

	ActionDeactivate(resource *GlobalAlertRule) error

	ActionMute(resource *GlobalAlertRule) error

	ActionUnmute(resource *GlobalAlertRule) error
}

func newGlobalAlertRuleClient(apiClient *Client) *GlobalAlertRuleClient {
	return &GlobalAlertRuleClient{
		apiClient: apiClient,
	}
}

func (c *GlobalAlertRuleClient) Create(container *GlobalAlertRule) (*GlobalAlertRule, error) {
	resp := &GlobalAlertRule{}
	err := c.apiClient.Ops.DoCreate(GlobalAlertRuleType, container, resp)
	return resp, err
}

func (c *GlobalAlertRuleClient) Update(existing *GlobalAlertRule, updates interface{}) (*GlobalAlertRule, error) {
	resp := &GlobalAlertRule{}
	err := c.apiClient.Ops.DoUpdate(GlobalAlertRuleType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *GlobalAlertRuleClient) Replace(obj *GlobalAlertRule) (*GlobalAlertRule, error) {
	resp := &GlobalAlertRule{}
	err := c.apiClient.Ops.DoReplace(GlobalAlertRuleType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *GlobalAlertRuleClient) List(opts *types.ListOpts) (*GlobalAlertRuleCollection, error) {
	resp := &GlobalAlertRuleCollection{}
	err := c.apiClient.Ops.DoList(GlobalAlertRuleType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *GlobalAlertRuleCollection) Next() (*GlobalAlertRuleCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &GlobalAlertRuleCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *GlobalAlertRuleClient) ByID(id string) (*GlobalAlertRule, error) {
	resp := &GlobalAlertRule{}
	err := c.apiClient.Ops.DoByID(GlobalAlertRuleType, id, resp)
	return resp, err
}

func (c *GlobalAlertRuleClient) Delete(container *GlobalAlertRule) error {
	return c.apiClient.Ops.DoResourceDelete(GlobalAlertRuleType, &container.Resource)
}

func (c *GlobalAlertRuleClient) ActionActivate(resource *GlobalAlertRule) error {
	err := c.apiClient.Ops.DoAction(GlobalAlertRuleType, "activate", &resource.Resource, nil, nil)
	return err
}

func (c *GlobalAlertRuleClient) ActionDeactivate(resource *GlobalAlertRule) error {
	err := c.apiClient.Ops.DoAction(GlobalAlertRuleType, "deactivate", &resource.Resource, nil, nil)
	return err
}

func (c *GlobalAlertRuleClient) ActionMute(resource *GlobalAlertRule) error {
	err := c.apiClient.Ops.DoAction(GlobalAlertRuleType, "mute", &resource.Resource, nil, nil)
	return err
}

func (c *GlobalAlertRuleClient) ActionUnmute(resource *GlobalAlertRule) error {
	err := c.apiClient.Ops.DoAction(GlobalAlertRuleType, "unmute", &resource.Resource, nil, nil)
	return err
}
