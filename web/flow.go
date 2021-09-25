package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsValidFlow(c *gin.Context) (bool, gin.H) {
	flowID := c.Query("flow")
	if flowID == "" {
		c.Redirect(http.StatusFound, KratosPublicURL+"/self-service/registration/browser")
		return false, gin.H{}
	}

	cookie := c.Request.Header.Get("Cookie")
	if cookie == "" {
		c.Redirect(http.StatusFound, KratosPublicURL+"/self-service/registration/browser")
		return false, gin.H{}
	}

	flow, _, err := KratosPublicClient.V0alpha1Api.GetSelfServiceRegistrationFlow(c.Request.Context()).Id(flowID).Cookie(cookie).Execute()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return false, gin.H{}
	}

	data := gin.H{
		"flow": flow,
	}

	logoutURL, _, err := KratosPublicClient.V0alpha1Api.CreateSelfServiceLogoutFlowUrlForBrowsers(c.Request.Context()).Cookie(cookie).Execute()
	if err == nil {
		data["logout_url"] = logoutURL.LogoutUrl
	}

	return true, data
}
