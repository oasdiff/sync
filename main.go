package main

import (
	"github.com/oasdiff/go-common/ds"
	"github.com/oasdiff/go-common/env"
	"github.com/oasdiff/go-common/gcs"
	"github.com/oasdiff/go-common/slack"
	"github.com/oasdiff/sync/internal"
)

func main() {

	_ = internal.SetupRouter(ds.NewClientWrapper(env.GetGCloudProject()),
		gcs.NewStore(), slack.NewClientWrapper()).Run(":8080")
}
