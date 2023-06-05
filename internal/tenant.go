package internal

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/oasdiff/go-common/ds"
	"github.com/sirupsen/logrus"
)

type CreateTenantRequest struct {
	Tenant string `json:"tenant"`
}

func (h *Handle) CreateTenant(c *gin.Context) {

	if c.Request.Body == nil {
		logrus.Info("invalid create tenant request with 'empty body'")
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var payload CreateTenantRequest
	err := json.NewDecoder(c.Request.Body).Decode(&payload)
	if err != nil {
		logrus.Infof("failed to decode create tenant request body with '%v'", err)
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if payload.Tenant == "" {
		logrus.Info("invalid create tenant request with 'empty tenant name'")
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	id := uuid.NewString()
	err = h.dsc.Put(ds.KindTenant, id, &ds.Tenant{
		Id:      id,
		Name:    payload.Tenant,
		Created: time.Now().Unix(),
	})
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.sc.Info(fmt.Sprintf("tenant '%s' created", id))
	c.JSON(http.StatusCreated, gin.H{"id": id})
}
