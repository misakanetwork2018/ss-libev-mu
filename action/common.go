package action

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func respondWithError(code int, message string, c *gin.Context) {
	c.JSON(code, gin.H{
		"success": false,
		"msg":     message,
	})
	c.Abort()
}

func parseParams(c *gin.Context, model interface{}) error {
	err := c.ShouldBind(model)

	if err != nil {
		fmt.Println(err.Error())
		respondWithError(500, "param error", c)
	}

	return err
}
