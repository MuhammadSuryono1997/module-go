package response

import "github.com/gin-gonic/gin"

func Json(code int, message string, obj interface{}) {
	var c *gin.Context
	c.JSON(code, gin.H{"code": code, "message": message, "data": obj})
}

