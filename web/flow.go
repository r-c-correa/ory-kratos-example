package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FlowType int

const (
	FlowLogin FlowType = iota + 1
	FlowSignup
	FlowSettings
)

var ErrInvalidFlowType = errors.New("flow type is invalid")

func IsValidFlow(c *gin.Context, flowType FlowType) (bool, gin.H, error) {
	flowTypeS := ""
	switch flowType {
	case FlowLogin:
		flowTypeS = "login"
	case FlowSignup:
		flowTypeS = "registration"
	case FlowSettings:
		flowTypeS = "settings"
	default:
		return false, gin.H{}, ErrInvalidFlowType
	}

	flowID := c.Query("flow")
	if flowID == "" {
		c.Redirect(http.StatusFound, KratosPublicURL+"/self-service/"+flowTypeS+"/browser")
		return false, gin.H{}, nil
	}

	cookie := c.Request.Header.Get("Cookie")
	if cookie == "" {
		c.Redirect(http.StatusFound, KratosPublicURL+"/self-service/"+flowTypeS+"/browser")
		return false, gin.H{}, nil
	}

	data := gin.H{}
	var flow interface{}
	var err error

	switch flowType {
	case FlowLogin:
		flow, _, err = KratosPublicClient.V0alpha1Api.GetSelfServiceLoginFlow(c.Request.Context()).Id(flowID).Cookie(cookie).Execute()
	case FlowSignup:
		flow, _, err = KratosPublicClient.V0alpha1Api.GetSelfServiceRegistrationFlow(c.Request.Context()).Id(flowID).Cookie(cookie).Execute()
	case FlowSettings:
		flow, _, err = KratosPublicClient.V0alpha1Api.GetSelfServiceSettingsFlow(c.Request.Context()).Id(flowID).Cookie(cookie).Execute()
	}

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return false, data, err
	}
	data["flow"] = flow

	if logoutURL := LogoutURL(c, cookie); logoutURL != "" {
		data["logout_url"] = logoutURL
	}

	return true, data, nil
}

func LogoutURL(c *gin.Context, cookie string) string {
	logoutURL, _, err := KratosPublicClient.V0alpha1Api.CreateSelfServiceLogoutFlowUrlForBrowsers(c.Request.Context()).Cookie(cookie).Execute()
	if err != nil {
		return ""
	}

	return logoutURL.LogoutUrl
}
