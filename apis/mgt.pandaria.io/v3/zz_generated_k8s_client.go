package v3

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

	SensitiveFiltersGetter
}

type Client struct {
	sync.Mutex
	restClient rest.Interface
	starters   []controller.Starter

	sensitiveFilterControllers map[string]SensitiveFilterController
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

		sensitiveFilterControllers: map[string]SensitiveFilterController{},
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

type SensitiveFiltersGetter interface {
	SensitiveFilters(namespace string) SensitiveFilterInterface
}

func (c *Client) SensitiveFilters(namespace string) SensitiveFilterInterface {
	objectClient := objectclient.NewObjectClient(namespace, c.restClient, &SensitiveFilterResource, SensitiveFilterGroupVersionKind, sensitiveFilterFactory{})
	return &sensitiveFilterClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}
