package client

import (
	"github.com/rancher/norman/types"
)

const (
	ExternalDNSType                   = "externalDNS"
	ExternalDNSFieldAnnotations       = "annotations"
	ExternalDNSFieldCreated           = "created"
	ExternalDNSFieldCreatorID         = "creatorId"
	ExternalDNSFieldDNSRecordType     = "dnsRecordType"
	ExternalDNSFieldDomainName        = "domainName"
	ExternalDNSFieldLabels            = "labels"
	ExternalDNSFieldLoadBalanceMethod = "loadBalanceMethod"
	ExternalDNSFieldName              = "name"
	ExternalDNSFieldNamespaceId       = "namespaceId"
	ExternalDNSFieldOwnerReferences   = "ownerReferences"
	ExternalDNSFieldPools             = "pools"
	ExternalDNSFieldProjectID         = "projectId"
	ExternalDNSFieldRemoved           = "removed"
	ExternalDNSFieldUUID              = "uuid"
)

type ExternalDNS struct {
	types.Resource
	Annotations       map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created           string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID         string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	DNSRecordType     string            `json:"dnsRecordType,omitempty" yaml:"dnsRecordType,omitempty"`
	DomainName        string            `json:"domainName,omitempty" yaml:"domainName,omitempty"`
	Labels            map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	LoadBalanceMethod string            `json:"loadBalanceMethod,omitempty" yaml:"loadBalanceMethod,omitempty"`
	Name              string            `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId       string            `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`
	OwnerReferences   []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Pools             []DNSPool         `json:"pools,omitempty" yaml:"pools,omitempty"`
	ProjectID         string            `json:"projectId,omitempty" yaml:"projectId,omitempty"`
	Removed           string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	UUID              string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ExternalDNSCollection struct {
	types.Collection
	Data   []ExternalDNS `json:"data,omitempty"`
	client *ExternalDNSClient
}

type ExternalDNSClient struct {
	apiClient *Client
}

type ExternalDNSOperations interface {
	List(opts *types.ListOpts) (*ExternalDNSCollection, error)
	ListAll(opts *types.ListOpts) (*ExternalDNSCollection, error)
	Create(opts *ExternalDNS) (*ExternalDNS, error)
	Update(existing *ExternalDNS, updates interface{}) (*ExternalDNS, error)
	Replace(existing *ExternalDNS) (*ExternalDNS, error)
	ByID(id string) (*ExternalDNS, error)
	Delete(container *ExternalDNS) error
}

func newExternalDNSClient(apiClient *Client) *ExternalDNSClient {
	return &ExternalDNSClient{
		apiClient: apiClient,
	}
}

func (c *ExternalDNSClient) Create(container *ExternalDNS) (*ExternalDNS, error) {
	resp := &ExternalDNS{}
	err := c.apiClient.Ops.DoCreate(ExternalDNSType, container, resp)
	return resp, err
}

func (c *ExternalDNSClient) Update(existing *ExternalDNS, updates interface{}) (*ExternalDNS, error) {
	resp := &ExternalDNS{}
	err := c.apiClient.Ops.DoUpdate(ExternalDNSType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ExternalDNSClient) Replace(obj *ExternalDNS) (*ExternalDNS, error) {
	resp := &ExternalDNS{}
	err := c.apiClient.Ops.DoReplace(ExternalDNSType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *ExternalDNSClient) List(opts *types.ListOpts) (*ExternalDNSCollection, error) {
	resp := &ExternalDNSCollection{}
	err := c.apiClient.Ops.DoList(ExternalDNSType, opts, resp)
	resp.client = c
	return resp, err
}

func (c *ExternalDNSClient) ListAll(opts *types.ListOpts) (*ExternalDNSCollection, error) {
	resp := &ExternalDNSCollection{}
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

func (cc *ExternalDNSCollection) Next() (*ExternalDNSCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ExternalDNSCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ExternalDNSClient) ByID(id string) (*ExternalDNS, error) {
	resp := &ExternalDNS{}
	err := c.apiClient.Ops.DoByID(ExternalDNSType, id, resp)
	return resp, err
}

func (c *ExternalDNSClient) Delete(container *ExternalDNS) error {
	return c.apiClient.Ops.DoResourceDelete(ExternalDNSType, &container.Resource)
}
