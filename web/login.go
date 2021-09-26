package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	isValidFlow, dataFlow, _ := IsValidFlow(c, FlowLogin)
	if !isValidFlow {
		return
	}

	c.HTML(http.StatusOK, "login.gohtml", dataFlow)
}
