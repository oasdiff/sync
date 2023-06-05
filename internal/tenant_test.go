package internal_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oasdiff/sync/internal"
	"github.com/oasdiff/sync/internal/ds"
	"github.com/oasdiff/sync/internal/gcs"
	"github.com/oasdiff/sync/internal/slack"
	"github.com/stretchr/testify/require"
)

func TestCreateTenant(t *testing.T) {

	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodPost, "/tenants", Encode(t, internal.CreateTenantRequest{Tenant: "test"}))
	require.NoError(t, err)

	internal.SetupRouter(ds.NewInMemoryClient(), gcs.NewInMemoryStore(),
		slack.NewInMemoryClient()).ServeHTTP(w, r)

	require.Equal(t, http.StatusCreated, w.Result().StatusCode)
}

func Encode(t *testing.T, v any) *bytes.Buffer {

	var buf bytes.Buffer
	require.NoError(t, json.NewEncoder(&buf).Encode(v))

	return &buf
}
