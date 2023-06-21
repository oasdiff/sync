package internal_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/oasdiff/sync/internal"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestETE(t *testing.T) {

	t.Skip()

	syncUrl := os.Getenv("SYNC_URL")
	require.NotEmpty(t, syncUrl)

	createWebhook(t, syncUrl, createTenant(t, syncUrl))
}

func createWebhook(t *testing.T, syncUrl string, tenant string) {

	var buf bytes.Buffer
	require.NoError(t, json.NewEncoder(&buf).Encode(internal.CreateWebhookRequest{
		Spec: "https://raw.githubusercontent.com/Tufin/oasdiff/main/data/simple.yaml",
	}))

	response, err := http.Post(fmt.Sprintf("%s/tenants/%s/webhooks", syncUrl, tenant), "application/json", &buf)
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, response.StatusCode)
}

func createTenant(t *testing.T, syncUrl string) string {

	var buf bytes.Buffer
	require.NoError(t, json.NewEncoder(&buf).Encode(internal.CreateTenantRequest{
		Tenant:   "itest",
		Callback: "https://test/me",
	}))

	response, err := http.Post(fmt.Sprintf("%s/tenants", syncUrl), "application/json", &buf)
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, response.StatusCode)

	var payload map[string]string
	require.NoError(t, json.NewDecoder(response.Body).Decode(&payload))
	logrus.Infof("created tenant '%s'", payload["id"])

	return payload["id"]
}
