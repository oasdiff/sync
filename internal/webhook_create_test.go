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
		WebhookName: "OpenAI",
		Owner:       "openai",
		Repo:        "openai-openapi",
		Branch:      "master",
		Path:        "openapi.yaml",
		Spec:        "https://github.com/openai/openai-openapi/raw/e145786e70bf5fc1bc73c7cd19884f445d52c383/openapi.yaml",
	}))
	r, err := http.NewRequest(http.MethodPost, "/tenants/f1/webhooks", &buf)
	require.NoError(t, err)

	w := httptest.NewRecorder()

	internal.SetupRouter(ds.NewInMemoryClient(nil), gcs.NewInMemoryStore(nil),
		slack.NewInMemoryClient()).ServeHTTP(w, r)

	require.Equal(t, http.StatusCreated, w.Result().StatusCode)

	var res map[string]string
	err = json.NewDecoder(w.Result().Body).Decode(&res)
	require.NoError(t, err)
	require.NotEmpty(t, res["id"])
}
