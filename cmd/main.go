package main

import (
	"messageSystem/conf"
	"messageSystem/dao/influx"
	"messageSystem/router"
	"time"
)

func init() {
	err := conf.Init()
	if err != nil {
		panic(err)
	}
	err = influx.Init(conf.Conf)
	if err != nil {
		panic(err)
	}
}
func main() {
	go func() {
		for {
			influx.WritePoint()
			time.Sleep(time.Second)
		}
	}()
	r := router.NewEngine()
	r.Run(":8081")

}
