package action

import (
	"github.com/gin-gonic/gin"
	"ss-libev-mu/utils"
)

var (
	OldTraffic = make(map[string]int)
)

func Traffic() func(c *gin.Context) {
	return func(c *gin.Context) {
		var data = make(map[string]int)

		stats, err := utils.GetStats()

		if err != nil {
			c.JSON(200, gin.H{
				"success": false,
				"msg":     err.Error(),
			})
		} else {
			for port, traffic := range stats {
				if oldTraffic, ok := OldTraffic[port]; ok {
					data[port] = traffic - oldTraffic
				} else { // for new
					data[port] = traffic
				}

				if data[port] < 0 {
					data[port] = traffic
				}
			}

			OldTraffic = stats

			c.JSON(200, gin.H{
				"success": true,
				"data":    data,
			})
		}
	}
}
