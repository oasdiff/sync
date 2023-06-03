package internal_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oasdiff/sync/internal"
	"github.com/oasdiff/sync/internal/ds"
	"github.com/oasdiff/sync/internal/gcs"
	"github.com/stretchr/testify/require"
)

func TestCreateWebhook(t *testing.T) {

	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodPost, "/tenants/1/webhooks", nil)
	require.NoError(t, err)

	internal.SetupRouter(ds.NewInMemoryClient(), gcs.NewInMemoryStore()).ServeHTTP(w, r)

	require.Equal(t, http.StatusCreated, w.Result().StatusCode)
}
