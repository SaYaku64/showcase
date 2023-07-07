package gin

import (
	"github.com/SaYaku64/showcase/internal/broker"
	"github.com/gin-gonic/gin"
)

type ginModule struct {
	router *gin.Engine
}

func New() broker.IBroker {
	module := ginModule{
		router: gin.New(),
	}

	go module.Run()

	return &module
}

func (gm *ginModule) Run() {
	v1 := gm.router.Group("/api/v1")
	v1.GET("/info/get", gm.GetInfo)
	v1.POST("/info/set", gm.SetInfo)
	v1.DELETE("/info/delete")

	v2 := gm.router.Group("/api/v2")
	v2.GET("/user/:userID", gm.GetUserByID)

	gm.router.Run()
}

func (gm *ginModule) Send(address string, data any) {
	//
}
