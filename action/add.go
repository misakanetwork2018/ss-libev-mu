package action

import (
	"github.com/gin-gonic/gin"
	"ss-libev-mu/model"
	"ss-libev-mu/ss_manager"
)

func AddUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		var user model.User

		if parseParams(c, &user) != nil {
			return
		}

		user.Mode = "tcp_and_udp"

		ok, err := ss_manager.Add(user)

		var msg string

		if err != nil {
			msg = err.Error()
		}

		c.JSON(200, gin.H{
			"success": ok,
			"msg":     msg,
		})
	}
}
