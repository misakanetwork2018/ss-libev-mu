package model

// for add
type User struct {
	ServerPort int    `json:"server_port" form:"port" binding:"required"`
	Password   string `json:"password" form:"pass" binding:"required"`
	Method     string `json:"method" form:"method" binding:"required"`
	FastOpen   bool   `json:"fast_open"`
	Mode       string `json:"mode"`
	Plugin     string `json:"plugin"`
	PluginOpts string `json:"plugin_opts"`
}
