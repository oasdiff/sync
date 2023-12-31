package main

import (
	"os"

	"github.com/oasdiff/go-common/ds"
	"github.com/oasdiff/go-common/env"
	"github.com/oasdiff/go-common/gcs"
	"github.com/oasdiff/go-common/slack"
	"github.com/oasdiff/sync/internal"
)

func main() {

	dsc := ds.NewClient(env.GetGCPProject(), "sync")
	defer dsc.Close()

	store := gcs.NewStore()
	defer store.Close()

	if err := internal.SetupRouter(dsc, store, slack.NewClientWrapper()).
		Run(":8080"); err != nil {
		os.Exit(1)
	}
}
