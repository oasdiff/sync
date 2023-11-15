package internal_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oasdiff/go-common/ds"
	"github.com/oasdiff/go-common/gcs"
	"github.com/oasdiff/go-common/slack"
	"github.com/oasdiff/sync/internal"
	"github.com/stretchr/testify/require"
)

func TestHandle_GetChangelog(t *testing.T) {

	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/tenants/1/webhooks/2/changelog/3.html", nil)
	require.NoError(t, err)

	internal.SetupRouter(ds.NewInMemoryClient(nil), gcs.NewInMemoryStore(nil),
		slack.NewInMemoryClient()).ServeHTTP(w, r)
}
