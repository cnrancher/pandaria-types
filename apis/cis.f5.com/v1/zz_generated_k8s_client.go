package v1

import (
	"context"
	"sync"

	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/objectclient"
	"github.com/rancher/norman/objectclient/dynamic"
	"github.com/rancher/norman/restwatch"
	"k8s.io/client-go/rest"
)

type (
	contextKeyType        struct{}
	contextClientsKeyType struct{}
)

type Interface interface {
	RESTClient() rest.Interface
	controller.Starter

	VirtualServersGetter
	TLSProfilesGetter
	TransportServersGetter
	ExternalDNSsGetter
}

type Client struct {
	sync.Mutex
	restClient rest.Interface
	starters   []controller.Starter

	virtualServerControllers   map[string]VirtualServerController
	tlsProfileControllers      map[string]TLSProfileController
	transportServerControllers map[string]TransportServerController
	externalDNSControllers     map[string]ExternalDNSController
}

func NewForConfig(config rest.Config) (Interface, error) {
	if config.NegotiatedSerializer == nil {
		config.NegotiatedSerializer = dynamic.NegotiatedSerializer
	}

	restClient, err := restwatch.UnversionedRESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &Client{
		restClient: restClient,

		virtualServerControllers:   map[string]VirtualServerController{},
		tlsProfileControllers:      map[string]TLSProfileController{},
		transportServerControllers: map[string]TransportServerController{},
		externalDNSControllers:     map[string]ExternalDNSController{},
	}, nil
}

func (c *Client) RESTClient() rest.Interface {
	return c.restClient
}

func (c *Client) Sync(ctx context.Context) error {
	return controller.Sync(ctx, c.starters...)
}

func (c *Client) Start(ctx context.Context, threadiness int) error {
	return controller.Start(ctx, threadiness, c.starters...)
}

type VirtualServersGetter interface {
	VirtualServers(namespace string) VirtualServerInterface
}

func (c *Client) VirtualServers(namespace string) VirtualServerInterface {
	objectClient := objectclient.NewObjectClient(namespace, c.restClient, &VirtualServerResource, VirtualServerGroupVersionKind, virtualServerFactory{})
	return &virtualServerClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}

type TLSProfilesGetter interface {
	TLSProfiles(namespace string) TLSProfileInterface
}

func (c *Client) TLSProfiles(namespace string) TLSProfileInterface {
	objectClient := objectclient.NewObjectClient(namespace, c.restClient, &TLSProfileResource, TLSProfileGroupVersionKind, tlsProfileFactory{})
	return &tlsProfileClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}

type TransportServersGetter interface {
	TransportServers(namespace string) TransportServerInterface
}

func (c *Client) TransportServers(namespace string) TransportServerInterface {
	objectClient := objectclient.NewObjectClient(namespace, c.restClient, &TransportServerResource, TransportServerGroupVersionKind, transportServerFactory{})
	return &transportServerClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}

type ExternalDNSsGetter interface {
	ExternalDNSs(namespace string) ExternalDNSInterface
}

func (c *Client) ExternalDNSs(namespace string) ExternalDNSInterface {
	objectClient := objectclient.NewObjectClient(namespace, c.restClient, &ExternalDNSResource, ExternalDNSGroupVersionKind, externalDNSFactory{})
	return &externalDNSClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}
