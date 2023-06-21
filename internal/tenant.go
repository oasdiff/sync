package internal

import (
	"errors"
	"fmt"
	"net/http"
	"net/mail"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/oasdiff/go-common/ds"
	"github.com/sirupsen/logrus"
)

type CreateTenantRequest struct {
	Tenant       string `json:"tenant"`
	Email        string `json:"email"`
	Callback     string `json:"callback"`
	SlackChannel string `json:"slack_channel"`
}

func (h *Handle) CreateTenant(c *gin.Context) {

	payload, err := getCreateTenantRequest(c.Request)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	id := uuid.NewString()
	t := ds.Tenant{
		Id:   id,
		Name: payload.Tenant,

		Created: time.Now().Unix(),
	}
	err = h.dsc.Put(ds.KindTenant, id, &t)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.sc.Info(fmt.Sprintf("tenant created '%+v'", t))
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func getCreateTenantRequest(request *http.Request) (*CreateTenantRequest, error) {

	if request.Body == nil {
		err := errors.New("invalid create tenant request with 'empty body'")
		logrus.Info(err.Error())
		return nil, err
	}

	var res CreateTenantRequest
	if err := json.NewDecoder(request.Body).Decode(&res); err != nil {
		err = fmt.Errorf("failed to decode create tenant request body with '%v'", err)
		logrus.Info(err.Error())
		return nil, err
	}

	if res.Tenant == "" {
		err := errors.New("invalid create tenant request with 'empty tenant name'")
		logrus.Info(err.Error())
		return nil, err
	}

	if _, err := mail.ParseAddress(res.Email); err != nil {
		err = fmt.Errorf("[create tenant request] failed to parse email '%s' with '%v'", res.Email, err)
		logrus.Info(err.Error())
		return nil, err
	}

	if res.Callback != "" {
		if _, err := url.Parse(res.Callback); err != nil {
			err = fmt.Errorf("[create tenant request] failed to parse callback '%s' with '%v'", res.Callback, err)
			logrus.Info(err.Error())
			return nil, err
		}
	}

	if res.SlackChannel != "" {
		if !strings.HasPrefix(res.SlackChannel, "https://hooks.slack.com/") {
			err := fmt.Errorf("[create tenant request] invalid slack channel url '%s' should start with 'https://hooks.slack.com/'", res.SlackChannel)
			logrus.Info(err.Error())
			return nil, err
		} else if _, err := url.Parse(res.SlackChannel); err != nil {
			err = fmt.Errorf("[create tenant request] failed to parse slack channel url '%s' with '%v'", res.SlackChannel, err)
			logrus.Info(err.Error())
			return nil, err
		}
	}

	if res.Callback == "" && res.SlackChannel == "" {
		err := errors.New("create tenant request failed with 'both callback and slack channel are empty'")
		logrus.Info(err.Error())
		return nil, err
	}

	return &res, nil
}
