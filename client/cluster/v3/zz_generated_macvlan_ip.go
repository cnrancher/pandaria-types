package client

import (
	"github.com/rancher/norman/types"
)

const (
	MacvlanIPType                 = "macvlanIP"
	MacvlanIPFieldAnnotations     = "annotations"
	MacvlanIPFieldCIDR            = "cidr"
	MacvlanIPFieldCreated         = "created"
	MacvlanIPFieldCreatorID       = "creatorId"
	MacvlanIPFieldLabels          = "labels"
	MacvlanIPFieldMAC             = "mac"
	MacvlanIPFieldName            = "name"
	MacvlanIPFieldNamespaceId     = "namespaceId"
	MacvlanIPFieldOwnerReferences = "ownerReferences"
	MacvlanIPFieldPodID           = "podId"
	MacvlanIPFieldRemoved         = "removed"
	MacvlanIPFieldUUID            = "uuid"
	MacvlanIPFieldVLAN            = "vlan"
)

type MacvlanIP struct {
	types.Resource
	Annotations     map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	CIDR            string            `json:"cidr,omitempty" yaml:"cidr,omitempty"`
	Created         string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Labels          map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	MAC             string            `json:"mac,omitempty" yaml:"mac,omitempty"`
	Name            string            `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId     string            `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	PodID           string            `json:"podId,omitempty" yaml:"podId,omitempty"`
	Removed         string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	UUID            string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	VLAN            string            `json:"vlan,omitempty" yaml:"vlan,omitempty"`
}

type MacvlanIPCollection struct {
	types.Collection
	Data   []MacvlanIP `json:"data,omitempty"`
	client *MacvlanIPClient
}

type MacvlanIPClient struct {
	apiClient *Client
}

type MacvlanIPOperations interface {
	List(opts *types.ListOpts) (*MacvlanIPCollection, error)
	Create(opts *MacvlanIP) (*MacvlanIP, error)
	Update(existing *MacvlanIP, updates interface{}) (*MacvlanIP, error)
	Replace(existing *MacvlanIP) (*MacvlanIP, error)
	ByID(id string) (*MacvlanIP, error)
	Delete(container *MacvlanIP) error
}

func newMacvlanIPClient(apiClient *Client) *MacvlanIPClient {
	return &MacvlanIPClient{
		apiClient: apiClient,
	}
}

func (c *MacvlanIPClient) Create(container *MacvlanIP) (*MacvlanIP, error) {
	resp := &MacvlanIP{}
	err := c.apiClient.Ops.DoCreate(MacvlanIPType, container, resp)
	return resp, err
}

func (c *MacvlanIPClient) Update(existing *MacvlanIP, updates interface{}) (*MacvlanIP, error) {
	resp := &MacvlanIP{}
	err := c.apiClient.Ops.DoUpdate(MacvlanIPType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *MacvlanIPClient) Replace(obj *MacvlanIP) (*MacvlanIP, error) {
	resp := &MacvlanIP{}
	err := c.apiClient.Ops.DoReplace(MacvlanIPType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *MacvlanIPClient) List(opts *types.ListOpts) (*MacvlanIPCollection, error) {
	resp := &MacvlanIPCollection{}
	err := c.apiClient.Ops.DoList(MacvlanIPType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *MacvlanIPCollection) Next() (*MacvlanIPCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &MacvlanIPCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *MacvlanIPClient) ByID(id string) (*MacvlanIP, error) {
	resp := &MacvlanIP{}
	err := c.apiClient.Ops.DoByID(MacvlanIPType, id, resp)
	return resp, err
}

func (c *MacvlanIPClient) Delete(container *MacvlanIP) error {
	return c.apiClient.Ops.DoResourceDelete(MacvlanIPType, &container.Resource)
}
