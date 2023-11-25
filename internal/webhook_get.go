package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oasdiff/go-common/ds"
)

func (h *Handle) GetWebhooks(c *gin.Context) {

	tenant := c.Param(PathParamTenantId)
	ok := validateTenant(h.dsc, tenant)
	if !ok {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var webhooks []ds.Webhook
	err := h.dsc.GetFilter(ds.KindWebhook,
		[]ds.FilterField{{Name: "tenant_id", Operator: ds.Equal1, Value: tenant}},
		&webhooks)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, webhooks)
}
