package env

import (
	"os"

	"github.com/sirupsen/logrus"
)

func GetGCloudProject() string {

	return failIfEmpty("GCLOUD_PROJECT")
}

func failIfEmpty(key string) string {

	res := os.Getenv(key)
	if res == "" {
		logrus.Fatalf("Please, add environment variable '%s'", key)
	}
	logrus.Debugf("%s: %s", key, res)

	return res
}
