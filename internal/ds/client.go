package ds

import (
	"context"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/oasdiff/sync/internal/env"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

type Kind string

const (
	KindTenant  Kind = "tenant"
	KindWebhook Kind = "webhook"

	namespace = "sync"
)

type Client interface {
	Get(kind Kind, id string, dst interface{}) error
	GetAll(kind Kind, dst interface{}) error
	Put(kind Kind, id string, src interface{}) error
}

type ClientWrapper struct {
	ds *datastore.Client
}

func NewClientWrapper(project string) Client {

	if key := env.GetDatastoreKey(); key != "" {
		conf, err := google.JWTConfigFromJSON([]byte(key), datastore.ScopeDatastore)
		if err != nil {
			logrus.Fatalf("failed to config datastore JWT from JSON key with '%v'", err)
		}

		ctx := context.Background()
		opt := []option.ClientOption{option.WithTokenSource(conf.TokenSource(ctx))}

		dataStoreEndPoint := os.Getenv("DATASTORE_ENDPOINT")
		if dataStoreEndPoint != "" {
			opt = append(opt, option.WithEndpoint(dataStoreEndPoint))
		}

		client, err := datastore.NewClient(ctx, project, opt...)
		if err != nil {
			logrus.Fatalf("failed to create datastore client with '%v'", err)
		}

		return &ClientWrapper{ds: client}
	}

	client, err := datastore.NewClient(context.Background(), project)
	if err != nil {
		logrus.Fatalf("failed to create datastore client without token with '%v'", err)
	}

	return &ClientWrapper{ds: client}
}

func (c *ClientWrapper) Get(kind Kind, id string, dst interface{}) error {

	err := c.ds.Get(context.Background(), getKey(kind, id), dst)
	if err != nil {
		logrus.Errorf("failed to get '%s' id '%s' from datastore namespace '%s' with '%v'", kind, id, namespace, err)
	}

	return err
}

func (c *ClientWrapper) GetAll(kind Kind, dst interface{}) error {

	q := datastore.NewQuery(string(kind)).Namespace(namespace)
	_, err := c.ds.GetAll(context.Background(), q, dst)
	if err != nil {
		logrus.Errorf("failed to get all '%s' from datastore namespace '%s' with '%v'", kind, namespace, err)
	}

	return err
}

func (c *ClientWrapper) Put(kind Kind, id string, src interface{}) error {

	_, err := c.ds.Put(context.Background(), getKey(kind, id), src)
	if err != nil {
		logrus.Errorf("failed to update '%s/%s' item '%+v' type: '%T' with '%v'", namespace, kind, src, src, err)
	} else {
		logrus.Infof("created '%s/%s: %+v' type: '%T'", namespace, kind, src, src)
	}

	return err
}

func getKey(kind Kind, id string) *datastore.Key {

	res := datastore.NameKey(string(kind), id, nil)
	res.Namespace = namespace

	return res
}
