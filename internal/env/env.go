package env

import (
	"os"

	"github.com/sirupsen/logrus"
)

func GetGCloudProject() string {

	return failIfEmpty("GCLOUD_PROJECT")
}

func GetDatastoreKey() string {

	return failIfEmpty("DATASTORE_KEY")
}

func GetSlackInfoUrl() string {

	return failIfEmpty("SLACK_INFO_URL")
}

func failIfEmpty(key string) string {

	res := os.Getenv(key)
	if res == "" {
		logrus.Fatalf("Please, add environment variable '%s'", key)
	}

	return res
}
