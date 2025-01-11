
package main

import (
	"gin-example/biz/handler"
	"gin-example/biz/mw"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	mw.Register(r)
	handler.Register(r)

	r.Run()
}
	