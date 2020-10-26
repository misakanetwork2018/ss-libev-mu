package action

import (
	"github.com/gin-gonic/gin"
	"ss-libev-mu/utils"
)

func Traffic() func(c *gin.Context) {
	return func(c *gin.Context) {
		data, err := utils.GetStats()

		if err != nil {
			c.JSON(200, gin.H{
				"success": false,
				"msg":     err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"success": true,
				"data":    data,
			})
		}
	}
}
