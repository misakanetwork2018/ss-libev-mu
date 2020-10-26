package action

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"ss-libev-mu/model"
	"ss-libev-mu/ss_manager"
)

func List() func(c *gin.Context) {
	return func(c *gin.Context) {
		var users []model.Port

		str, _, err := ss_manager.List()

		if err != nil {
			fmt.Println("get stat error: ", err.Error())
			respondWithError(500, err.Error(), c)
			return
		}

		if err = json.Unmarshal([]byte(str), &users); err != nil {
			fmt.Println("parse json error: ", err.Error())
			respondWithError(500, err.Error(), c)
			return
		}

		c.JSON(200, gin.H{
			"success": true,
			"data":    users,
		})
	}
}
