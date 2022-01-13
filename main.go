package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "postd")
		name := c.DefaultQuery("name", "fage")
		fmt.Println(types)
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	r.Run()
}
