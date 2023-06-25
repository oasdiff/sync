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

func TestCreateTenant(t *testing.T) {

	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodPost, "/tenants",
		Encode(t, internal.CreateTenantRequest{
			Tenant:       "test",
			Email:        "james@my-company.com",
			Callback:     "https://api.my-company.com/webhooks",
			SlackChannel: "https://hooks.slack.com/services/TLDF14G/AG123/abcd",
		}))
	require.NoError(t, err)

	internal.SetupRouter(ds.NewInMemoryClient(nil), gcs.NewInMemoryStore(nil),
		slack.NewInMemoryClient()).ServeHTTP(w, r)

	require.Equal(t, http.StatusCreated, w.Result().StatusCode)

	var res map[string]string
	err = json.NewDecoder(w.Result().Body).Decode(&res)
	require.NoError(t, err)
	require.NotEmpty(t, res["id"])
}

func Encode(t *testing.T, v any) *bytes.Buffer {

	var buf bytes.Buffer
	require.NoError(t, json.NewEncoder(&buf).Encode(v))

	return &buf
}
