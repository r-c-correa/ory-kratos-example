package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Settings(c *gin.Context) {
	isValidFlow, dataFlow, _ := IsValidFlow(c, FlowSettings)
	if !isValidFlow {
		return
	}

	c.HTML(http.StatusOK, "settings.gohtml", dataFlow)
}
