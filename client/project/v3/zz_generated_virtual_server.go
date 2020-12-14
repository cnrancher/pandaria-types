package client

import (
	"github.com/rancher/norman/types"
)

const (
	VirtualServerType                        = "virtualServer"
	VirtualServerFieldAnnotations            = "annotations"
	VirtualServerFieldCreated                = "created"
	VirtualServerFieldCreatorID              = "creatorId"
	VirtualServerFieldHTTPTraffic            = "httpTraffic"
	VirtualServerFieldHost                   = "host"
	VirtualServerFieldLabels                 = "labels"
	VirtualServerFieldName                   = "name"
	VirtualServerFieldNamespaceId            = "namespaceId"
	VirtualServerFieldOwnerReferences        = "ownerReferences"
	VirtualServerFieldPools                  = "pools"
	VirtualServerFieldProjectID              = "projectId"
	VirtualServerFieldRemoved                = "removed"
	VirtualServerFieldRewriteAppRoot         = "rewriteAppRoot"
	VirtualServerFieldSNAT                   = "snat"
	VirtualServerFieldTLSProfileName         = "tlsProfileName"
	VirtualServerFieldUUID                   = "uuid"
	VirtualServerFieldVirtualServerAddress   = "virtualServerAddress"
	VirtualServerFieldVirtualServerHTTPPort  = "virtualServerHTTPPort"
	VirtualServerFieldVirtualServerHTTPSPort = "virtualServerHTTPSPort"
	VirtualServerFieldVirtualServerName      = "virtualServerName"
	VirtualServerFieldWAF                    = "waf"
)

type VirtualServer struct {
	types.Resource
	Annotations            map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created                string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID              string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	HTTPTraffic            string            `json:"httpTraffic,omitempty" yaml:"httpTraffic,omitempty"`
	Host                   string            `json:"host,omitempty" yaml:"host,omitempty"`
	Labels                 map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                   string            `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId            string            `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`
	OwnerReferences        []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Pools                  []Pool            `json:"pools,omitempty" yaml:"pools,omitempty"`
	ProjectID              string            `json:"projectId,omitempty" yaml:"projectId,omitempty"`
	Removed                string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	RewriteAppRoot         string            `json:"rewriteAppRoot,omitempty" yaml:"rewriteAppRoot,omitempty"`
	SNAT                   string            `json:"snat,omitempty" yaml:"snat,omitempty"`
	TLSProfileName         string            `json:"tlsProfileName,omitempty" yaml:"tlsProfileName,omitempty"`
	UUID                   string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	VirtualServerAddress   string            `json:"virtualServerAddress,omitempty" yaml:"virtualServerAddress,omitempty"`
	VirtualServerHTTPPort  int64             `json:"virtualServerHTTPPort,omitempty" yaml:"virtualServerHTTPPort,omitempty"`
	VirtualServerHTTPSPort int64             `json:"virtualServerHTTPSPort,omitempty" yaml:"virtualServerHTTPSPort,omitempty"`
	VirtualServerName      string            `json:"virtualServerName,omitempty" yaml:"virtualServerName,omitempty"`
	WAF                    string            `json:"waf,omitempty" yaml:"waf,omitempty"`
}

type VirtualServerCollection struct {
	types.Collection
	Data   []VirtualServer `json:"data,omitempty"`
	client *VirtualServerClient
}

type VirtualServerClient struct {
	apiClient *Client
}

type VirtualServerOperations interface {
	List(opts *types.ListOpts) (*VirtualServerCollection, error)
	ListAll(opts *types.ListOpts) (*VirtualServerCollection, error)
	Create(opts *VirtualServer) (*VirtualServer, error)
	Update(existing *VirtualServer, updates interface{}) (*VirtualServer, error)
	Replace(existing *VirtualServer) (*VirtualServer, error)
	ByID(id string) (*VirtualServer, error)
	Delete(container *VirtualServer) error
}

func newVirtualServerClient(apiClient *Client) *VirtualServerClient {
	return &VirtualServerClient{
		apiClient: apiClient,
	}
}

func (c *VirtualServerClient) Create(container *VirtualServer) (*VirtualServer, error) {
	resp := &VirtualServer{}
	err := c.apiClient.Ops.DoCreate(VirtualServerType, container, resp)
	return resp, err
}

func (c *VirtualServerClient) Update(existing *VirtualServer, updates interface{}) (*VirtualServer, error) {
	resp := &VirtualServer{}
	err := c.apiClient.Ops.DoUpdate(VirtualServerType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *VirtualServerClient) Replace(obj *VirtualServer) (*VirtualServer, error) {
	resp := &VirtualServer{}
	err := c.apiClient.Ops.DoReplace(VirtualServerType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *VirtualServerClient) List(opts *types.ListOpts) (*VirtualServerCollection, error) {
	resp := &VirtualServerCollection{}
	err := c.apiClient.Ops.DoList(VirtualServerType, opts, resp)
	resp.client = c
	return resp, err
}

func (c *VirtualServerClient) ListAll(opts *types.ListOpts) (*VirtualServerCollection, error) {
	resp := &VirtualServerCollection{}
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

func (cc *VirtualServerCollection) Next() (*VirtualServerCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &VirtualServerCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *VirtualServerClient) ByID(id string) (*VirtualServer, error) {
	resp := &VirtualServer{}
	err := c.apiClient.Ops.DoByID(VirtualServerType, id, resp)
	return resp, err
}

func (c *VirtualServerClient) Delete(container *VirtualServer) error {
	return c.apiClient.Ops.DoResourceDelete(VirtualServerType, &container.Resource)
}
