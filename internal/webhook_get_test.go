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

func TestGetWebhooks(t *testing.T) {

	r, err := http.NewRequest(http.MethodGet, "/tenants/f1/webhooks", nil)
	require.NoError(t, err)

	w := httptest.NewRecorder()

	internal.SetupRouter(ds.NewInMemoryClient(nil), gcs.NewInMemoryStore(nil),
		slack.NewInMemoryClient()).ServeHTTP(w, r)

	require.Equal(t, http.StatusOK, w.Result().StatusCode)
}
