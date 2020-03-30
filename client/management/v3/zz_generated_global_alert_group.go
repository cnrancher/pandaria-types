package client

import (
	"github.com/rancher/norman/types"
)

const (
	GlobalAlertGroupType                       = "globalAlertGroup"
	GlobalAlertGroupFieldAlertState            = "alertState"
	GlobalAlertGroupFieldAnnotations           = "annotations"
	GlobalAlertGroupFieldCreated               = "created"
	GlobalAlertGroupFieldCreatorID             = "creatorId"
	GlobalAlertGroupFieldDescription           = "description"
	GlobalAlertGroupFieldGroupIntervalSeconds  = "groupIntervalSeconds"
	GlobalAlertGroupFieldGroupWaitSeconds      = "groupWaitSeconds"
	GlobalAlertGroupFieldLabels                = "labels"
	GlobalAlertGroupFieldName                  = "name"
	GlobalAlertGroupFieldOwnerReferences       = "ownerReferences"
	GlobalAlertGroupFieldRecipients            = "recipients"
	GlobalAlertGroupFieldRemoved               = "removed"
	GlobalAlertGroupFieldRepeatIntervalSeconds = "repeatIntervalSeconds"
	GlobalAlertGroupFieldState                 = "state"
	GlobalAlertGroupFieldTransitioning         = "transitioning"
	GlobalAlertGroupFieldTransitioningMessage  = "transitioningMessage"
	GlobalAlertGroupFieldUUID                  = "uuid"
)

type GlobalAlertGroup struct {
	types.Resource
	AlertState            string            `json:"alertState,omitempty" yaml:"alertState,omitempty"`
	Annotations           map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created               string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID             string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Description           string            `json:"description,omitempty" yaml:"description,omitempty"`
	GroupIntervalSeconds  int64             `json:"groupIntervalSeconds,omitempty" yaml:"groupIntervalSeconds,omitempty"`
	GroupWaitSeconds      int64             `json:"groupWaitSeconds,omitempty" yaml:"groupWaitSeconds,omitempty"`
	Labels                map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                  string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences       []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Recipients            []Recipient       `json:"recipients,omitempty" yaml:"recipients,omitempty"`
	Removed               string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	RepeatIntervalSeconds int64             `json:"repeatIntervalSeconds,omitempty" yaml:"repeatIntervalSeconds,omitempty"`
	State                 string            `json:"state,omitempty" yaml:"state,omitempty"`
	Transitioning         string            `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage  string            `json:"transitioningMessage,omitempty" yaml:"transitioningMessage,omitempty"`
	UUID                  string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type GlobalAlertGroupCollection struct {
	types.Collection
	Data   []GlobalAlertGroup `json:"data,omitempty"`
	client *GlobalAlertGroupClient
}

type GlobalAlertGroupClient struct {
	apiClient *Client
}

type GlobalAlertGroupOperations interface {
	List(opts *types.ListOpts) (*GlobalAlertGroupCollection, error)
	Create(opts *GlobalAlertGroup) (*GlobalAlertGroup, error)
	Update(existing *GlobalAlertGroup, updates interface{}) (*GlobalAlertGroup, error)
	Replace(existing *GlobalAlertGroup) (*GlobalAlertGroup, error)
	ByID(id string) (*GlobalAlertGroup, error)
	Delete(container *GlobalAlertGroup) error
}

func newGlobalAlertGroupClient(apiClient *Client) *GlobalAlertGroupClient {
	return &GlobalAlertGroupClient{
		apiClient: apiClient,
	}
}

func (c *GlobalAlertGroupClient) Create(container *GlobalAlertGroup) (*GlobalAlertGroup, error) {
	resp := &GlobalAlertGroup{}
	err := c.apiClient.Ops.DoCreate(GlobalAlertGroupType, container, resp)
	return resp, err
}

func (c *GlobalAlertGroupClient) Update(existing *GlobalAlertGroup, updates interface{}) (*GlobalAlertGroup, error) {
	resp := &GlobalAlertGroup{}
	err := c.apiClient.Ops.DoUpdate(GlobalAlertGroupType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *GlobalAlertGroupClient) Replace(obj *GlobalAlertGroup) (*GlobalAlertGroup, error) {
	resp := &GlobalAlertGroup{}
	err := c.apiClient.Ops.DoReplace(GlobalAlertGroupType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *GlobalAlertGroupClient) List(opts *types.ListOpts) (*GlobalAlertGroupCollection, error) {
	resp := &GlobalAlertGroupCollection{}
	err := c.apiClient.Ops.DoList(GlobalAlertGroupType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *GlobalAlertGroupCollection) Next() (*GlobalAlertGroupCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &GlobalAlertGroupCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *GlobalAlertGroupClient) ByID(id string) (*GlobalAlertGroup, error) {
	resp := &GlobalAlertGroup{}
	err := c.apiClient.Ops.DoByID(GlobalAlertGroupType, id, resp)
	return resp, err
}

func (c *GlobalAlertGroupClient) Delete(container *GlobalAlertGroup) error {
	return c.apiClient.Ops.DoResourceDelete(GlobalAlertGroupType, &container.Resource)
}
