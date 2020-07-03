package main

import (
	"asso_api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/asso", handlers.GetAssociations())
	r.Run()
}
