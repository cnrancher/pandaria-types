package client

import (
	"github.com/rancher/norman/types"
)

const (
	SensitiveFilterType                 = "sensitiveFilter"
	SensitiveFilterFieldAnnotations     = "annotations"
	SensitiveFilterFieldCreated         = "created"
	SensitiveFilterFieldCreatorID       = "creatorId"
	SensitiveFilterFieldFilters         = "filters"
	SensitiveFilterFieldLabels          = "labels"
	SensitiveFilterFieldName            = "name"
	SensitiveFilterFieldOwnerReferences = "ownerReferences"
	SensitiveFilterFieldRemoved         = "removed"
	SensitiveFilterFieldUUID            = "uuid"
)

type SensitiveFilter struct {
	types.Resource
	Annotations     map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created         string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Filters         []Filter          `json:"filters,omitempty" yaml:"filters,omitempty"`
	Labels          map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name            string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed         string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	UUID            string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type SensitiveFilterCollection struct {
	types.Collection
	Data   []SensitiveFilter `json:"data,omitempty"`
	client *SensitiveFilterClient
}

type SensitiveFilterClient struct {
	apiClient *Client
}

type SensitiveFilterOperations interface {
	List(opts *types.ListOpts) (*SensitiveFilterCollection, error)
	ListAll(opts *types.ListOpts) (*SensitiveFilterCollection, error)
	Create(opts *SensitiveFilter) (*SensitiveFilter, error)
	Update(existing *SensitiveFilter, updates interface{}) (*SensitiveFilter, error)
	Replace(existing *SensitiveFilter) (*SensitiveFilter, error)
	ByID(id string) (*SensitiveFilter, error)
	Delete(container *SensitiveFilter) error
}

func newSensitiveFilterClient(apiClient *Client) *SensitiveFilterClient {
	return &SensitiveFilterClient{
		apiClient: apiClient,
	}
}

func (c *SensitiveFilterClient) Create(container *SensitiveFilter) (*SensitiveFilter, error) {
	resp := &SensitiveFilter{}
	err := c.apiClient.Ops.DoCreate(SensitiveFilterType, container, resp)
	return resp, err
}

func (c *SensitiveFilterClient) Update(existing *SensitiveFilter, updates interface{}) (*SensitiveFilter, error) {
	resp := &SensitiveFilter{}
	err := c.apiClient.Ops.DoUpdate(SensitiveFilterType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *SensitiveFilterClient) Replace(obj *SensitiveFilter) (*SensitiveFilter, error) {
	resp := &SensitiveFilter{}
	err := c.apiClient.Ops.DoReplace(SensitiveFilterType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *SensitiveFilterClient) List(opts *types.ListOpts) (*SensitiveFilterCollection, error) {
	resp := &SensitiveFilterCollection{}
	err := c.apiClient.Ops.DoList(SensitiveFilterType, opts, resp)
	resp.client = c
	return resp, err
}

func (c *SensitiveFilterClient) ListAll(opts *types.ListOpts) (*SensitiveFilterCollection, error) {
	resp := &SensitiveFilterCollection{}
	resp, err := c.List(opts)
	if err != nil {
		return resp, err
	}
	data := resp.Data
	for next, err := resp.Next(); next != nil && err == nil; next, err = next.Next() {
		data = append(data, next.Data...)
		resp = next
		resp.Data = data
	}
	if err != nil {
		return resp, err
	}
	return resp, err
}

func (cc *SensitiveFilterCollection) Next() (*SensitiveFilterCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &SensitiveFilterCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *SensitiveFilterClient) ByID(id string) (*SensitiveFilter, error) {
	resp := &SensitiveFilter{}
	err := c.apiClient.Ops.DoByID(SensitiveFilterType, id, resp)
	return resp, err
}

func (c *SensitiveFilterClient) Delete(container *SensitiveFilter) error {
	return c.apiClient.Ops.DoResourceDelete(SensitiveFilterType, &container.Resource)
}
