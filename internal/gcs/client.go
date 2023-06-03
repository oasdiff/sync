package gcs

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
	log "github.com/sirupsen/logrus"
)

const (
	bucketName       = "sync"
	pathTemplateSpec = "%s/spec/%s"
)

type Client interface {
	SaveSpec(tenantId string, name string) error
}

type Store struct {
	client *storage.Client
}

func NewStore() Client {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to create storage client with '%v'", err)
	}

	return &Store{client: client}
}

// Buckets/sync/{tenant-id}/spec/[]spec
func (store *Store) SaveSpec(tenantId string, name string) error {

	w := store.client.Bucket(bucketName).
		Object(fmt.Sprintf(pathTemplateSpec, tenantId, name)).
		NewWriter(context.Background())
	defer func() {
		if err := w.Close(); err != nil {
			log.Errorf("failed to close gcs writer with '%v'", err)
		}
	}()

	return nil
}
