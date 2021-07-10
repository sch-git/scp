package scp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddSCPReq struct {
	Name string `json:"name"`
}

func (m *AddSCPReq) handle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"data": "name",
	})
}

func AddSCP() gin.HandlerFunc {
	return func(c *gin.Context) {
		m := &AddSCPReq{}
		err := c.Bind(m)
		if nil != err {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})

			return
		}
		m.handle(c)
	}
}
