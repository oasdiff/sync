package internal

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/oasdiff/sync/internal/ds"
	"github.com/oasdiff/sync/internal/gcs"
	"github.com/oasdiff/sync/internal/slack"
)

const (
	PathParamTenantId = "tenant-id"
)

func SetupRouter(dsc ds.Client, store gcs.Client, sc slack.Client) *gin.Engine {

	h := NewHandle(dsc, store, sc)
	router := gin.Default()

	router.POST("/tenants", h.CreateTenant)
	router.POST(fmt.Sprintf("/tenants/:%s/webhooks", PathParamTenantId), h.CreateWebhook)

	return router
}
