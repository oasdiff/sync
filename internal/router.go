package internal

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/oasdiff/go-common/ds"
	"github.com/oasdiff/go-common/gcs"
	"github.com/oasdiff/go-common/slack"
)

const (
	PathParamTenantId    = "tenant-id"
	PathParamWebhookId   = "webhook-id"
	PathParamChangelogId = "changelog-id"
)

func SetupRouter(dsc ds.Client, store gcs.Client, sc slack.Client) *gin.Engine {

	h := NewHandle(dsc, store, sc)
	router := gin.Default()

	router.POST("/tenants", h.CreateTenant)
	router.POST(fmt.Sprintf("/tenants/:%s/webhooks", PathParamTenantId), h.CreateWebhook)
	router.GET(fmt.Sprintf("/tenants/:%s/webhooks/:%s/changelog/:%s",
		PathParamTenantId, PathParamWebhookId, PathParamChangelogId), h.GetChangelog)

	return router
}
