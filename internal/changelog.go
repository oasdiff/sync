package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oasdiff/go-common/gcs"
)

func (h *Handle) GetChangelog(c *gin.Context) {

	tenant, webhook, changelog := c.Param(PathParamTenantId), c.Param(PathParamWebhookId), c.Param(PathParamChangelogId)
	ok := validateTenant(h.dsc, tenant)
	if !ok {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	html, err := h.store.Read(gcs.GetSpecPath(tenant, webhook, changelog))
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", html)
}
