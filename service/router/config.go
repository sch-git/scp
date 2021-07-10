package router

import "github.com/gin-gonic/gin"

type routerConfig struct {
	Path   string
	Handle gin.HandlerFunc
}

func Init()  {
	r := gin.Default()
	handleSCP(r.Group("/scp"))
	r.Run("localhost:8080")
}
