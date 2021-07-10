package router

import (
	"github.com/gin-gonic/gin"
	"service/httpcontroller/scp"
)

func handleSCP(router *gin.RouterGroup) {
	routerConfig := []routerConfig{
		{"/add", scp.AddSCP()},
	}

	for _, conf := range routerConfig {
		router.POST(conf.Path, conf.Handle)
	}
}
