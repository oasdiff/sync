package internal

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/oasdiff/sync/internal/ds"
	"github.com/oasdiff/sync/internal/gcs"
)

const (
	PathParamTenantId = "tenant-id"
)

func SetupRouter(dsc ds.Client, store gcs.Client) *gin.Engine {

	h := NewHandle(dsc, store)
	router := gin.Default()

	router.POST("/tenants", h.CreateTenant)
	router.POST(fmt.Sprintf("/tenants/:%s/webhooks", PathParamTenantId), h.CreateWebhook)

	return router
}
