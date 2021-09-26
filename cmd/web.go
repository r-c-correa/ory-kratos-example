package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/r-c-correa/ory-kratos-example/web"
	"github.com/spf13/cobra"
)

var webCommand = &cobra.Command{
	Use:   "web",
	Short: "web application",
	RunE: func(cmd *cobra.Command, args []string) error {
		gin.SetMode(gin.DebugMode)

		r := gin.Default()
		r.LoadHTMLGlob("web/pages/*")
		r.GET("/", web.Home)
		r.GET("/registration", web.Registration)
		r.GET("/registration/", web.Registration)
		r.GET("/login", web.Login)
		r.GET("/settings", web.Settings)
		r.GET("/settings/", web.Settings)
		return r.Run(":4002")
	},
}

func init() {
	rootCmd.AddCommand(webCommand)
}
