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
	"github.com/stretchr/testify/require"
)

func TestCreateWebhook(t *testing.T) {

	var buf bytes.Buffer
	require.NoError(t, json.NewEncoder(&buf).Encode(internal.CreateWebhookRequest{
		Callback: "https://test/webhook",
		Spec:     "https://raw.githubusercontent.com/Tufin/oasdiff/main/data/simple.yaml",
	}))
	r, err := http.NewRequest(http.MethodPost, "/tenants/f1/webhooks", &buf)
	require.NoError(t, err)
	w := httptest.NewRecorder()

	internal.SetupRouter(ds.NewInMemoryClient(), gcs.NewInMemoryStore()).ServeHTTP(w, r)

	require.Equal(t, http.StatusCreated, w.Result().StatusCode)
}
