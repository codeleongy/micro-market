package handler

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	once   sync.Once
	router *gin.Engine
)

func GetRouter() *gin.Engine {
	once.Do(func() {
		InitRouter()
	})
	return router
}

func InitRouter() {

	router.GET("/*path", Proxy)
}

func Proxy(c *gin.Context) {

}
