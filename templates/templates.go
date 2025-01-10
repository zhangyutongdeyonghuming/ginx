package templates

// Template 结构体定义了模板的名称和内容
type Template struct {
	Name     string // 模板名称/文件名
	Template string // 模板内容
	Path     string // 模板路径
}

// Templates 定义了所有的模板
var Templates = []Template{
	{
		Name: "main.go",
		Template: `
package main

import (
	"{{.moduleName}}/biz/handler"
	"{{.moduleName}}/biz/mw"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	mw.RegisterMiddleware(r)
	register(r)

	r.Run()
}

// 注册路由
func register(r *gin.Engine) {
	r.GET("/ping", handler.Ping)
}

	`,
		Path: "./",
	},
	{
		Name: "go.mod",
		Template: `
module {{.moduleName}}

go {{.goVersion}}

		`,
		Path: "./",
	},
	{
		Name: "ping.go",
		Template: `
package handler

import "github.com/gin-gonic/gin"

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
		`,
		Path: "./biz/handler/",
	},
}
