package client

import (
	"github.com/rancher/norman/types"
)

const (
	MacvlanSubnetType                 = "macvlanSubnet"
	MacvlanSubnetFieldAnnotations     = "annotations"
	MacvlanSubnetFieldCIDR            = "cidr"
	MacvlanSubnetFieldCreated         = "created"
	MacvlanSubnetFieldCreatorID       = "creatorId"
	MacvlanSubnetFieldLabels          = "labels"
	MacvlanSubnetFieldMaster          = "master"
	MacvlanSubnetFieldMode            = "mode"
	MacvlanSubnetFieldName            = "name"
	MacvlanSubnetFieldNamespaceId     = "namespaceId"
	MacvlanSubnetFieldOwnerReferences = "ownerReferences"
	MacvlanSubnetFieldRemoved         = "removed"
	MacvlanSubnetFieldUUID            = "uuid"
)

type MacvlanSubnet struct {
	types.Resource
	Annotations     map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	CIDR            string            `json:"cidr,omitempty" yaml:"cidr,omitempty"`
	Created         string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Labels          map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Master          string            `json:"master,omitempty" yaml:"master,omitempty"`
	Mode            string            `json:"mode,omitempty" yaml:"mode,omitempty"`
	Name            string            `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId     string            `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed         string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	UUID            string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type MacvlanSubnetCollection struct {
	types.Collection
	Data   []MacvlanSubnet `json:"data,omitempty"`
	client *MacvlanSubnetClient
}

type MacvlanSubnetClient struct {
	apiClient *Client
}

type MacvlanSubnetOperations interface {
	List(opts *types.ListOpts) (*MacvlanSubnetCollection, error)
	Create(opts *MacvlanSubnet) (*MacvlanSubnet, error)
	Update(existing *MacvlanSubnet, updates interface{}) (*MacvlanSubnet, error)
	Replace(existing *MacvlanSubnet) (*MacvlanSubnet, error)
	ByID(id string) (*MacvlanSubnet, error)
	Delete(container *MacvlanSubnet) error
}

func newMacvlanSubnetClient(apiClient *Client) *MacvlanSubnetClient {
	return &MacvlanSubnetClient{
		apiClient: apiClient,
	}
}

func (c *MacvlanSubnetClient) Create(container *MacvlanSubnet) (*MacvlanSubnet, error) {
	resp := &MacvlanSubnet{}
	err := c.apiClient.Ops.DoCreate(MacvlanSubnetType, container, resp)
	return resp, err
}

func (c *MacvlanSubnetClient) Update(existing *MacvlanSubnet, updates interface{}) (*MacvlanSubnet, error) {
	resp := &MacvlanSubnet{}
	err := c.apiClient.Ops.DoUpdate(MacvlanSubnetType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *MacvlanSubnetClient) Replace(obj *MacvlanSubnet) (*MacvlanSubnet, error) {
	resp := &MacvlanSubnet{}
	err := c.apiClient.Ops.DoReplace(MacvlanSubnetType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *MacvlanSubnetClient) List(opts *types.ListOpts) (*MacvlanSubnetCollection, error) {
	resp := &MacvlanSubnetCollection{}
	err := c.apiClient.Ops.DoList(MacvlanSubnetType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *MacvlanSubnetCollection) Next() (*MacvlanSubnetCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &MacvlanSubnetCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *MacvlanSubnetClient) ByID(id string) (*MacvlanSubnet, error) {
	resp := &MacvlanSubnet{}
	err := c.apiClient.Ops.DoByID(MacvlanSubnetType, id, resp)
	return resp, err
}

func (c *MacvlanSubnetClient) Delete(container *MacvlanSubnet) error {
	return c.apiClient.Ops.DoResourceDelete(MacvlanSubnetType, &container.Resource)
}
