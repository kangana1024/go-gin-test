package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/kangana1024/go-gin-test/components"
)

/*
* readResource
 */
func ReadResource(c *gin.Context) {

	c.JSON(200, components.RestResponse{Code: 1, Message: "read resource successfully", Data: "resource"})
}

func WriteResource(c *gin.Context) {

	c.JSON(200, components.RestResponse{Code: 1, Message: "write resource successfully", Data: "resource"})
}
