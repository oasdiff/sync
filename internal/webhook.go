package internal

import (
	"context"
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
	"gopkg.in/yaml.v3"
)

type CreateWebhookRequest struct {
	Callback string `json:"callback"`
	Spec     string `json:"spec"`
}

func (h *Handle) CreateWebhook(c *gin.Context) {

	tenant := c.Param(PathParamTenantId)

	ok, request, oas := getCreateWebhookRequest(tenant, c.Request.Body)
	if !ok {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	payload, err := yaml.Marshal(oas)
	if err != nil {
		logrus.Errorf("failed to marshal OAS with '%v' tenant '%s' callback '%s'", err, tenant, request.Callback)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	now := time.Now().Unix()
	name := strconv.FormatInt(now, 10)

	err = h.store.UploadSpecFile(tenant, name, payload)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.dsc.Put(ds.KindWebhook, uuid.NewString(), &ds.Webhook{
		TenantId: tenant,
		Callback: request.Callback,
		Spec:     request.Spec,
		Copy:     name,
		Created:  now,
	})
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.WriteHeader(http.StatusCreated)
}

func getCreateWebhookRequest(tenant string, body io.ReadCloser) (bool, *CreateWebhookRequest, *openapi3.T) {

	if body == nil {
		logrus.Infof("invalid create webhook request with 'empty body' tenant '%s'", tenant)
		return false, nil, nil
	}

	var payload CreateWebhookRequest
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		logrus.Infof("failed to decode create tenant request body with '%v' tenant '%s'", err, tenant)
		return false, nil, nil
	}

	_, err = url.ParseRequestURI(payload.Callback)
	if err != nil {
		logrus.Infof("invalid callback '%s' tenant '%s'", payload.Callback, tenant)
		return false, nil, nil
	}

	oas, ok := validateSpec(tenant, payload.Spec)
	if !ok {
		return false, nil, nil
	}

	return ok, &payload, oas
}

func validateSpec(tenant string, spec string) (*openapi3.T, bool) {

	u, err := url.ParseRequestURI(spec)
	if err != nil {
		logrus.Infof("invalid spec url '%s' tenant '%s'", spec, tenant)
		return nil, false
	}

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true

	t, err := loader.LoadFromURI(u)
	if err != nil {
		logrus.Infof("failed to load OpenAPI spec from '%s' with '%v' tenant '%s'", spec, err, tenant)
		return nil, false
	}
	err = t.Validate(context.Background())
	if err != nil {
		logrus.Infof("failed to validate OpenAPI spec from '%s' with '%v' tenant '%s'", spec, err, tenant)
		return nil, false
	}

	return t, true
}
