package client

import (
	"github.com/rancher/norman/types"
)

const (
	TransportServerType                      = "transportServer"
	TransportServerFieldAnnotations          = "annotations"
	TransportServerFieldCreated              = "created"
	TransportServerFieldCreatorID            = "creatorId"
	TransportServerFieldLabels               = "labels"
	TransportServerFieldMode                 = "mode"
	TransportServerFieldName                 = "name"
	TransportServerFieldNamespaceId          = "namespaceId"
	TransportServerFieldOwnerReferences      = "ownerReferences"
	TransportServerFieldPool                 = "pool"
	TransportServerFieldProjectID            = "projectId"
	TransportServerFieldRemoved              = "removed"
	TransportServerFieldSNAT                 = "snat"
	TransportServerFieldUUID                 = "uuid"
	TransportServerFieldVirtualServerAddress = "virtualServerAddress"
	TransportServerFieldVirtualServerName    = "virtualServerName"
	TransportServerFieldVirtualServerPort    = "virtualServerPort"
)

type TransportServer struct {
	types.Resource
	Annotations          map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created              string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID            string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Labels               map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Mode                 string            `json:"mode,omitempty" yaml:"mode,omitempty"`
	Name                 string            `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId          string            `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`
	OwnerReferences      []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Pool                 *Pool             `json:"pool,omitempty" yaml:"pool,omitempty"`
	ProjectID            string            `json:"projectId,omitempty" yaml:"projectId,omitempty"`
	Removed              string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	SNAT                 string            `json:"snat,omitempty" yaml:"snat,omitempty"`
	UUID                 string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	VirtualServerAddress string            `json:"virtualServerAddress,omitempty" yaml:"virtualServerAddress,omitempty"`
	VirtualServerName    string            `json:"virtualServerName,omitempty" yaml:"virtualServerName,omitempty"`
	VirtualServerPort    int64             `json:"virtualServerPort,omitempty" yaml:"virtualServerPort,omitempty"`
}

type TransportServerCollection struct {
	types.Collection
	Data   []TransportServer `json:"data,omitempty"`
	client *TransportServerClient
}

type TransportServerClient struct {
	apiClient *Client
}

type TransportServerOperations interface {
	List(opts *types.ListOpts) (*TransportServerCollection, error)
	ListAll(opts *types.ListOpts) (*TransportServerCollection, error)
	Create(opts *TransportServer) (*TransportServer, error)
	Update(existing *TransportServer, updates interface{}) (*TransportServer, error)
	Replace(existing *TransportServer) (*TransportServer, error)
	ByID(id string) (*TransportServer, error)
	Delete(container *TransportServer) error
}

func newTransportServerClient(apiClient *Client) *TransportServerClient {
	return &TransportServerClient{
		apiClient: apiClient,
	}
}

func (c *TransportServerClient) Create(container *TransportServer) (*TransportServer, error) {
	resp := &TransportServer{}
	err := c.apiClient.Ops.DoCreate(TransportServerType, container, resp)
	return resp, err
}

func (c *TransportServerClient) Update(existing *TransportServer, updates interface{}) (*TransportServer, error) {
	resp := &TransportServer{}
	err := c.apiClient.Ops.DoUpdate(TransportServerType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *TransportServerClient) Replace(obj *TransportServer) (*TransportServer, error) {
	resp := &TransportServer{}
	err := c.apiClient.Ops.DoReplace(TransportServerType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *TransportServerClient) List(opts *types.ListOpts) (*TransportServerCollection, error) {
	resp := &TransportServerCollection{}
	err := c.apiClient.Ops.DoList(TransportServerType, opts, resp)
	resp.client = c
	return resp, err
}

func (c *TransportServerClient) ListAll(opts *types.ListOpts) (*TransportServerCollection, error) {
	resp := &TransportServerCollection{}
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

func (cc *TransportServerCollection) Next() (*TransportServerCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &TransportServerCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *TransportServerClient) ByID(id string) (*TransportServer, error) {
	resp := &TransportServer{}
	err := c.apiClient.Ops.DoByID(TransportServerType, id, resp)
	return resp, err
}

func (c *TransportServerClient) Delete(container *TransportServer) error {
	return c.apiClient.Ops.DoResourceDelete(TransportServerType, &container.Resource)
}
