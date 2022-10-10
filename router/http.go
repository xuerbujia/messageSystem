package router

import (
	"github.com/gin-gonic/gin"
	"messageSystem/controller"
	"net/http"
)

type BaseMessage struct {
	Temp    string `json:"temp"`
	Humi    string `json:"humi"`
	UsedMen string `json:"used_mem"`
	Light   string `json:"light"`
}

func NewEngine() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLFiles("./static/index.html")
	r.Static("./static", "static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/api/base-message", controller.GetMessage)
	return r
}
