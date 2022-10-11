package router

import (
	"github.com/gin-gonic/gin"
	"messageSystem/controller"
)

type BaseMessage struct {
	Temp    string `json:"temp"`
	Humi    string `json:"humi"`
	UsedMen string `json:"used_mem"`
	Light   string `json:"light"`
}

func NewEngine() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/api/base-message", controller.GetMessage)
	return r
}
