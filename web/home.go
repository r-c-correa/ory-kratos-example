package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	values := map[string]string{}

	for k, v := range c.Request.Header {
		finalValue := ""
		for _, v2 := range v {
			if finalValue != "" {
				finalValue += "<br>"
			}
			finalValue += v2
		}
		values[k] = finalValue
	}

	values["Sessions"] = ""
	for _, v := range c.Request.Cookies() {
		values["Sessions"] = v.String() + "<br>"
	}

	if _, s, err := SessionInfo(c.Writer, c.Request); err == nil {
		values["data"] = s
	} else {
		log.Println(err)
	}

	valuesSorted := make([]string, 0, len(values))
	for value := range values {
		valuesSorted = append(valuesSorted, value)
	}

	finalValue := ""
	sort.Strings(valuesSorted)
	for _, key := range valuesSorted {
		if key == "data" {
			continue
		}

		if finalValue != "" {
			finalValue += "<br>"
		}

		finalValue += fmt.Sprintf("<b>%v:</b> %v", key, values[key])
	}

	finalValue += "<br><b>User Data:</b> " + values["data"]

	c.HTML(http.StatusOK, "home.gohtml", gin.H{
		"info":       template.HTML(finalValue),
		"logout_url": LogoutURL(c, c.Request.Header.Get("Cookie")),
	})
}
