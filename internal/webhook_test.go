package internal_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oasdiff/go-common/ds"
	"github.com/oasdiff/go-common/gcs"
	"github.com/oasdiff/go-common/slack"
	"github.com/oasdiff/sync/internal"
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

	internal.SetupRouter(ds.NewInMemoryClient(), gcs.NewInMemoryStore(),
		slack.NewInMemoryClient()).ServeHTTP(w, r)

	require.Equal(t, http.StatusCreated, w.Result().StatusCode)
}
