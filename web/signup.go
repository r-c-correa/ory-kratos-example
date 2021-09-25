package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	isValidFlow, dataFlow := IsValidFlow(c)
	if !isValidFlow {
		return
	}

	c.HTML(http.StatusOK, "signup.gohtml", dataFlow)
}
