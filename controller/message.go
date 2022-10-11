package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"messageSystem/dao/influx"
	"net/http"
)

type BaseMessage struct {
	Temp    string `json:"temp"`
	Humi    string `json:"humi"`
	UsedMem string `json:"used_mem"`
	Light   string `json:"light"`
}

func GetMessage(c *gin.Context) {
	cmd := fmt.Sprintf("SELECT last(*) FROM pm ")
	//cmd := fmt.Sprintf("SELECT * FROM pm ORDER BY desc limit 1")
	res, err := influx.GetPoints(cmd)
	if err != nil {
		fmt.Println(err)
		return
	}
	var data = map[string]interface{}{}
	r := res[0].Series[0]
	for i, v := range r.Columns {
		data[v] = r.Values[0][i]
	}
	c.JSON(http.StatusOK, gin.H(data))

}
