package main

import (
	"github.com/oasdiff/sync/internal"
	"github.com/oasdiff/sync/internal/ds"
	"github.com/oasdiff/sync/internal/env"
	"github.com/oasdiff/sync/internal/gcs"
	"github.com/oasdiff/sync/internal/slack"
)

func main() {

	_ = internal.SetupRouter(ds.NewClientWrapper(env.GetGCloudProject()),
		gcs.NewStore(), slack.NewClientWrapper()).Run(":8080")
}
