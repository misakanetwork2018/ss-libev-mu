package main

import (
	"flag"
	"fmt"
	"github.com/akkuman/parseConfig"
	"github.com/gin-gonic/gin"
	"os"
	"ss-libev-mu/action"
	"ss-libev-mu/ss_manager"
	"v2ray-api/utils"
)

var (
	accessKey string
)

func main() {
	var configFile string

	flag.StringVar(&configFile, "c", "/etc/ss_mu.json", "Config file path, default: /etc/ss_mu.json")
	flag.Parse()

	var config = parseConfig.New(configFile)
	configI := config.Get("manager_address")
	if configI == nil {
		fmt.Println("You should provide a manager_address config, like /var/run/shadowsocks-manager.sock")
		os.Exit(1)
	}
	managerAddress := configI.(string)
	configI = config.Get("key")
	if configI == nil {
		fmt.Println("You should provide a key config")
		os.Exit(1)
	}
	accessKey = configI.(string)
	var bindAddress string
	configI = config.Get("bind_address")
	if configI == nil {
		bindAddress = "/tmp/ss-mu-api.sock"
	} else {
		bindAddress = configI.(string)
	}

	var address string
	addressI := config.Get("address")
	if addressI == nil {
		address = "127.0.0.1:8080"
	} else {
		address = addressI.(string)
	}

	ss_manager.Init(managerAddress, bindAddress)

	r := gin.Default()
	r.Use(webMiddleware)

	r.POST("/add", action.AddUser())
	r.POST("/del", action.DelUser())
	r.GET("/list", action.List())
	r.GET("/traffic", action.Traffic())
	r.GET("/status", action.Status())
	r.POST("/reboot", action.Reboot()) // only one-click version can use it

	_ = r.Run(address)
}

func webMiddleware(c *gin.Context) {
	token := c.GetHeader("X-Auth-Token")
	if token == "" {
		utils.RespondWithError(401, "API token required", c)
		return
	}
	if token != accessKey {
		utils.RespondWithError(401, "API token incorrect", c)
		return
	}
	c.Next()
}
