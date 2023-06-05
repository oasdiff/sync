package internal

import (
	"github.com/oasdiff/go-common/ds"
	"github.com/oasdiff/go-common/gcs"
	"github.com/oasdiff/go-common/slack"
)

type Handle struct {
	dsc   ds.Client
	store gcs.Client
	sc    slack.Client
}

func NewHandle(dsc ds.Client, store gcs.Client, sc slack.Client) *Handle {

	return &Handle{dsc: dsc, store: store, sc: sc}
}
