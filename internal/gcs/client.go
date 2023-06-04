package gcs

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
	"github.com/sirupsen/logrus"
)

const (
	bucketName       = "syncc"
	pathTemplateSpec = "%s/spec/%s"
)

type Client interface {
	UploadSpecFile(tenantId string, name string, file []byte) error
}

type Store struct {
	client *storage.Client
}

func NewStore() Client {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()

	client, err := storage.NewClient(ctx)
	if err != nil {
		logrus.Fatalf("failed to create storage client with '%v'", err)
	}

	return &Store{client: client}
}

// Buckets/syncc/{tenant-id}/spec/[]spec
func (store *Store) UploadSpecFile(tenantId string, name string, file []byte) error {

	w := store.client.Bucket(bucketName).
		Object(fmt.Sprintf(pathTemplateSpec, tenantId, name)).
		NewWriter(context.Background())
	defer func() {
		if err := w.Close(); err != nil {
			logrus.Errorf("failed to close gcs bucket '%s' writer file '%s' with '%v'",
				bucketName, name, err)
		}
	}()

	if _, err := w.Write(file); err != nil {
		logrus.Errorf("failed to create file in GCS bucket '%s' file '%s' with '%v'", bucketName, name, err)
		return err
	}

	return nil
}
