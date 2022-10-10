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
	//cmd := fmt.Sprintf("SELECT * FROM pm_monitoring ORDER BY desc limit 1")
	cmd := fmt.Sprintf("SELECT * FROM pm ORDER BY desc limit 1")
	res, err := influx.GetPoints(cmd)
	if err != nil {
		fmt.Println(err)
		return
	}
	//if len(res) == 0 {
	//	c.JSON(http.StatusGatewayTimeout, nil)
	//	return
	//}
	var data = map[string]interface{}{}
	r := res[0].Series[0]

	for i, v := range r.Columns {
		//fmt.Println(data[v], len(r.Values), i)
		data[v] = r.Values[0][i]
	}
	//	fmt.Println(r.Values)
	//	fmt.Println(data)
	c.JSON(http.StatusOK, gin.H(data))

}
