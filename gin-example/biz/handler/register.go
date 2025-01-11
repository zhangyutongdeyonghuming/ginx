
package handler

import (
	"github.com/gin-gonic/gin"
)

// 注册路由
func Register(r *gin.Engine) {
	r.GET("/ping", Ping)
}
		