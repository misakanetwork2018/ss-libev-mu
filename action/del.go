package action

import (
	"github.com/gin-gonic/gin"
	"ss-libev-mu/ss_manager"
	"ss-libev-mu/utils"
)

func DelUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			msg  string
			ok   bool
			stat int
		)

		port := c.PostForm("port")

		// get last traffic info before delete
		stats, err := utils.GetStats()

		if err == nil {
			stat = stats[port]

			ok, err = ss_manager.Remove(port)

			if err != nil {
				msg = err.Error()
			}
		} else {
			msg = err.Error()
		}

		c.JSON(200, gin.H{
			"success":  ok,
			"msg":      msg,
			"transfer": stat,
		})
	}
}
