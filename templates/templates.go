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
	mw.Register(r)
	handler.Register(r)

	r.Run()
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
	{
		Name: "register.go",
		Template: `
package handler

import (
	"github.com/gin-gonic/gin"
)

// 注册路由
func Register(r *gin.Engine) {
	r.GET("/ping", Ping)
}
		`,
		Path: "./biz/handler/",
	},
	{
		Name: "middleware.go",
		Template: `
package mw

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Register 注册中间件
func Register(r *gin.Engine) {
	r.Use(corsMw())
}

// cors 跨域中间件
func corsMw() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
		`,
		Path: "./biz/mw/",
	},
}
