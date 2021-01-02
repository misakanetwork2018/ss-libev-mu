package action

import (
	"github.com/gin-gonic/gin"
	"v2ray-api/utils"
)

func Reboot() func(c *gin.Context) {
	return func(c *gin.Context) {
		err, okS, errS := utils.Shell("systemctl restart ss-manager")
		if err == nil {
			c.JSON(200, gin.H{
				"success": true,
				"msg":     okS,
			})
		} else {
			respondWithError(500, errS, c)
			return
		}

		_, _, _ = utils.Shell("systemctl restart ss-libev-mu")
	}
}
