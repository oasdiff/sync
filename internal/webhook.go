package internal

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/oasdiff/sync/internal/ds"
	"github.com/sirupsen/logrus"
	"github.com/tufin/oasdiff/checker"
)

type CreateWebhookRequest struct {
	Url string `json:"url"`
}

func (h *Handle) CreateWebhook(c *gin.Context) {

	tenant := c.Param(PathParamTenantId)

	ok, request := getCreateWebhookRequest(tenant, c.Request.Body)
	if !ok {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	now := time.Now().Unix()
	name := strconv.FormatInt(now, 10)

	err := h.store.SaveSpec(tenant, name)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.dsc.Put(ds.KindWebhook, uuid.NewString(), &ds.Webhook{
		TenantId: tenant,
		Url:      request.Url,
		Spec:     name,
		Created:  now,
	})
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.WriteHeader(http.StatusCreated)
}

func getCreateWebhookRequest(tenant string, body io.ReadCloser) (bool, *CreateWebhookRequest) {

	if body == nil {
		logrus.Infof("invalid create webhook request with 'empty body' tenant '%s'", tenant)
		return false, nil
	}

	var payload CreateWebhookRequest
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		logrus.Infof("failed to decode create tenant request body with '%v' tenant '%s'", err, tenant)
		return false, nil
	}

	_, err = url.ParseRequestURI(payload.Url)
	if err != nil {
		logrus.Infof("invalid create webhook request url '%s' tenant '%s'", payload.Url, tenant)
		return false, nil
	}

	return validateSpec(tenant, payload.Url), &payload
}

func validateSpec(tenant string, url string) bool {

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true

	_, err := checker.LoadOpenAPISpecInfo(loader, url)
	if err != nil {
		logrus.Infof("failed to load OpenAPI spec from '%s' with '%v' tenant '%s'", url, err, tenant)
		return false
	}

	return true
}
