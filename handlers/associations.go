package handlers

import (
	"github.com/gin-gonic/gin"
)

type Search struct {
	Query string `json:"q"`
}

func GetAssociations(c *gin.Context) {
	var query Search
	search := c.ShouldBindQuery(&query)

	c.JSON(200, search)
}
