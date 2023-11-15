package internal

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/oasdiff/go-common/ds"
)

type CreateWebhookRequest struct {
	WebhookName string `json:"webhook_name"`
	Owner       string `json:"owner"`
	Repo        string `json:"repo"`
	Branch      string `json:"branch"`
	Path        string `json:"path"`
	Spec        string `json:"spec"`
}

func (h *Handle) CreateWebhook(c *gin.Context) {

	tenant := c.Param(PathParamTenantId)
	ok := validateTenant(h.dsc, tenant)
	if !ok {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	ok, request, oas := getCreateWebhookRequest(tenant, c.Request.Body)
	if !ok {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	id := uuid.NewString()
	now := time.Now().Unix()

	// upload OpenAPI base spec as a copy to GCS
	file := strconv.FormatInt(now, 10)
	err := h.store.UploadSpec(tenant, id, file, oas)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	webhook := ds.Webhook{
		Id:       id,
		Name:     request.WebhookName,
		TenantId: tenant,
		Owner:    request.Owner,
		Repo:     request.Repo,
		Path:     request.Path,
		Branch:   request.Branch,
		Spec:     request.Spec,
		Copy:     file,
		Created:  now,
		Updated:  now,
	}
	err = h.dsc.Put(ds.KindWebhook, id, &webhook)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.sc.Info(fmt.Sprintf("webhook created '%+v'", webhook))
	c.Writer.WriteHeader(http.StatusCreated)
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func validateTenant(dsc ds.Client, tenantId string) bool {

	var tenant ds.Tenant
	err := dsc.Get(ds.KindTenant, tenantId, &tenant)
	if err != nil || tenant.Id == "" || tenant.Id != tenantId {
		return false
	}

	return true
}

func getCreateWebhookRequest(tenant string, body io.ReadCloser) (bool, *CreateWebhookRequest, *openapi3.T) {

	if body == nil {
		slog.Info("invalid create webhook request", "error", "empty body", "tenant", tenant)
		return false, nil, nil
	}

	var payload CreateWebhookRequest
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		slog.Info("failed to decode create webhook request", "error", err, "tenant", tenant)
		return false, nil, nil
	}

	if ok := validateRequest(tenant, payload); !ok {
		return false, nil, nil
	}

	oas, ok := validateSpec(tenant, payload.Spec)
	if !ok {
		return false, nil, nil
	}

	return ok, &payload, oas
}

func validateRequest(tenant string, r CreateWebhookRequest) bool {

	if r.WebhookName == "" {
		slog.Info("invalid create webhook request", "error", "empty webhook name", "tenant", tenant)
		return false
	}
	if r.Owner == "" {
		slog.Info("invalid create webhook request", "error", "empty webhook repo owner", "tenant", tenant)
		return false
	}
	if r.Repo == "" {
		slog.Info("invalid create webhook request", "error", "empty webhook repo", "tenant", tenant)
		return false
	}
	if r.Branch == "" {
		slog.Info("invalid create webhook request", "error", "empty webhook branch", "tenant", tenant)
		return false
	}
	if r.Path == "" {
		slog.Info("invalid create webhook request", "error", "empty webhook OpenAPI revision file path", "tenant", tenant)
		return false
	}

	revision := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/%s", r.Owner, r.Repo, r.Branch, r.Path)
	if _, ok := validateSpec(tenant, revision); !ok {
		slog.Info("invalid create webhook request", "error", "invalid revision spec", "revision spec URL", revision, "tenant", tenant)
		return false
	}

	return true
}

func validateSpec(tenant string, specUrl string) (*openapi3.T, bool) {

	u, err := url.ParseRequestURI(specUrl)
	if err != nil {
		slog.Info("invalid spec url", "error", err, "spec URL", specUrl, "tenant", tenant)
		return nil, false
	}

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true

	t, err := loader.LoadFromURI(u)
	if err != nil {
		slog.Info("failed to load OpenAPI spec", err, "spec URL", specUrl, "tenant", tenant)
		return nil, false
	}
	// err = t.Validate(context.Background())
	// if err != nil {
	// 	logrus.Infof("failed to validate OpenAPI spec from '%s' with '%v' tenant '%s'", spec, err, tenant)
	// 	return nil, false
	// }

	return t, true
}
