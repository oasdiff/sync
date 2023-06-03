package internal

import (
	"github.com/oasdiff/sync/internal/ds"
	"github.com/oasdiff/sync/internal/gcs"
)

type Handle struct {
	dsc   ds.Client
	store gcs.Client
}

func NewHandle(dsc ds.Client, store gcs.Client) *Handle {

	return &Handle{dsc: dsc, store: store}
}
