package internal

import (
	"github.com/oasdiff/sync/internal/ds"
	"github.com/oasdiff/sync/internal/gcs"
	"github.com/oasdiff/sync/internal/slack"
)

type Handle struct {
	dsc   ds.Client
	store gcs.Client
	sc    slack.Client
}

func NewHandle(dsc ds.Client, store gcs.Client, sc slack.Client) *Handle {

	return &Handle{dsc: dsc, store: store, sc: sc}
}
