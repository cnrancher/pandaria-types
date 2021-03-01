package client

import (
	"github.com/rancher/norman/types"
)

const (
	TLSProfileType                 = "tlsProfile"
	TLSProfileFieldAnnotations     = "annotations"
	TLSProfileFieldCreated         = "created"
	TLSProfileFieldCreatorID       = "creatorId"
	TLSProfileFieldHosts           = "hosts"
	TLSProfileFieldLabels          = "labels"
	TLSProfileFieldName            = "name"
	TLSProfileFieldNamespaceId     = "namespaceId"
	TLSProfileFieldOwnerReferences = "ownerReferences"
	TLSProfileFieldProjectID       = "projectId"
	TLSProfileFieldRemoved         = "removed"
	TLSProfileFieldTLS             = "tls"
	TLSProfileFieldUUID            = "uuid"
)

type TLSProfile struct {
	types.Resource
	Annotations     map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created         string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Hosts           []string          `json:"hosts,omitempty" yaml:"hosts,omitempty"`
	Labels          map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name            string            `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId     string            `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	ProjectID       string            `json:"projectId,omitempty" yaml:"projectId,omitempty"`
	Removed         string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	TLS             *TLS              `json:"tls,omitempty" yaml:"tls,omitempty"`
	UUID            string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type TLSProfileCollection struct {
	types.Collection
	Data   []TLSProfile `json:"data,omitempty"`
	client *TLSProfileClient
}

type TLSProfileClient struct {
	apiClient *Client
}

type TLSProfileOperations interface {
	List(opts *types.ListOpts) (*TLSProfileCollection, error)
	ListAll(opts *types.ListOpts) (*TLSProfileCollection, error)
	Create(opts *TLSProfile) (*TLSProfile, error)
	Update(existing *TLSProfile, updates interface{}) (*TLSProfile, error)
	Replace(existing *TLSProfile) (*TLSProfile, error)
	ByID(id string) (*TLSProfile, error)
	Delete(container *TLSProfile) error
}

func newTLSProfileClient(apiClient *Client) *TLSProfileClient {
	return &TLSProfileClient{
		apiClient: apiClient,
	}
}

func (c *TLSProfileClient) Create(container *TLSProfile) (*TLSProfile, error) {
	resp := &TLSProfile{}
	err := c.apiClient.Ops.DoCreate(TLSProfileType, container, resp)
	return resp, err
}

func (c *TLSProfileClient) Update(existing *TLSProfile, updates interface{}) (*TLSProfile, error) {
	resp := &TLSProfile{}
	err := c.apiClient.Ops.DoUpdate(TLSProfileType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *TLSProfileClient) Replace(obj *TLSProfile) (*TLSProfile, error) {
	resp := &TLSProfile{}
	err := c.apiClient.Ops.DoReplace(TLSProfileType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *TLSProfileClient) List(opts *types.ListOpts) (*TLSProfileCollection, error) {
	resp := &TLSProfileCollection{}
	err := c.apiClient.Ops.DoList(TLSProfileType, opts, resp)
	resp.client = c
	return resp, err
}

func (c *TLSProfileClient) ListAll(opts *types.ListOpts) (*TLSProfileCollection, error) {
	resp := &TLSProfileCollection{}
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

func (cc *TLSProfileCollection) Next() (*TLSProfileCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &TLSProfileCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *TLSProfileClient) ByID(id string) (*TLSProfile, error) {
	resp := &TLSProfile{}
	err := c.apiClient.Ops.DoByID(TLSProfileType, id, resp)
	return resp, err
}

func (c *TLSProfileClient) Delete(container *TLSProfile) error {
	return c.apiClient.Ops.DoResourceDelete(TLSProfileType, &container.Resource)
}
