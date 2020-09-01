package client

import (
	"github.com/rancher/norman/types"
)

const (
	CloneAppType                           = "cloneApp"
	CloneAppFieldAnnotations               = "annotations"
	CloneAppFieldCertificateList           = "certificateList"
	CloneAppFieldConfigMapList             = "configMapList"
	CloneAppFieldCreated                   = "created"
	CloneAppFieldCreatorID                 = "creatorId"
	CloneAppFieldDockerCredentialList      = "credentialList"
	CloneAppFieldIngressList               = "ingressList"
	CloneAppFieldLabels                    = "labels"
	CloneAppFieldName                      = "name"
	CloneAppFieldOwnerReferences           = "ownerReferences"
	CloneAppFieldPersistentVolumeClaimList = "pvcList"
	CloneAppFieldRemoved                   = "removed"
	CloneAppFieldSecretList                = "secretList"
	CloneAppFieldTarget                    = "target"
	CloneAppFieldUUID                      = "uuid"
	CloneAppFieldWorkload                  = "workload"
)

type CloneApp struct {
	types.Resource
	Annotations               map[string]string        `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	CertificateList           []map[string]interface{} `json:"certificateList,omitempty" yaml:"certificateList,omitempty"`
	ConfigMapList             []ConfigMap              `json:"configMapList,omitempty" yaml:"configMapList,omitempty"`
	Created                   string                   `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID                 string                   `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	DockerCredentialList      []map[string]interface{} `json:"credentialList,omitempty" yaml:"credentialList,omitempty"`
	IngressList               []Ingress                `json:"ingressList,omitempty" yaml:"ingressList,omitempty"`
	Labels                    map[string]string        `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                      string                   `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences           []OwnerReference         `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	PersistentVolumeClaimList []map[string]interface{} `json:"pvcList,omitempty" yaml:"pvcList,omitempty"`
	Removed                   string                   `json:"removed,omitempty" yaml:"removed,omitempty"`
	SecretList                []map[string]interface{} `json:"secretList,omitempty" yaml:"secretList,omitempty"`
	Target                    *CloneTarget             `json:"target,omitempty" yaml:"target,omitempty"`
	UUID                      string                   `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Workload                  *Workload                `json:"workload,omitempty" yaml:"workload,omitempty"`
}

type CloneAppCollection struct {
	types.Collection
	Data   []CloneApp `json:"data,omitempty"`
	client *CloneAppClient
}

type CloneAppClient struct {
	apiClient *Client
}

type CloneAppOperations interface {
	List(opts *types.ListOpts) (*CloneAppCollection, error)
	ListAll(opts *types.ListOpts) (*CloneAppCollection, error)
	Create(opts *CloneApp) (*CloneApp, error)
	Update(existing *CloneApp, updates interface{}) (*CloneApp, error)
	Replace(existing *CloneApp) (*CloneApp, error)
	ByID(id string) (*CloneApp, error)
	Delete(container *CloneApp) error
}

func newCloneAppClient(apiClient *Client) *CloneAppClient {
	return &CloneAppClient{
		apiClient: apiClient,
	}
}

func (c *CloneAppClient) Create(container *CloneApp) (*CloneApp, error) {
	resp := &CloneApp{}
	err := c.apiClient.Ops.DoCreate(CloneAppType, container, resp)
	return resp, err
}

func (c *CloneAppClient) Update(existing *CloneApp, updates interface{}) (*CloneApp, error) {
	resp := &CloneApp{}
	err := c.apiClient.Ops.DoUpdate(CloneAppType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *CloneAppClient) Replace(obj *CloneApp) (*CloneApp, error) {
	resp := &CloneApp{}
	err := c.apiClient.Ops.DoReplace(CloneAppType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *CloneAppClient) List(opts *types.ListOpts) (*CloneAppCollection, error) {
	resp := &CloneAppCollection{}
	err := c.apiClient.Ops.DoList(CloneAppType, opts, resp)
	resp.client = c
	return resp, err
}

func (c *CloneAppClient) ListAll(opts *types.ListOpts) (*CloneAppCollection, error) {
	resp := &CloneAppCollection{}
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

func (cc *CloneAppCollection) Next() (*CloneAppCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &CloneAppCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *CloneAppClient) ByID(id string) (*CloneApp, error) {
	resp := &CloneApp{}
	err := c.apiClient.Ops.DoByID(CloneAppType, id, resp)
	return resp, err
}

func (c *CloneAppClient) Delete(container *CloneApp) error {
	return c.apiClient.Ops.DoResourceDelete(CloneAppType, &container.Resource)
}
