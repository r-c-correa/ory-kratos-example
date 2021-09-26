package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Registration(c *gin.Context) {
	isValidFlow, dataFlow, _ := IsValidFlow(c, FlowRegistration)
	if !isValidFlow {
		return
	}

	c.HTML(http.StatusOK, "registration.gohtml", dataFlow)
}
