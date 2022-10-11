package influx

import (
	influxdb "github.com/influxdata/influxdb1-client/v2"
	"log"
	"math/rand"
	"time"
)

func GetPoints(cmd string) (res []influxdb.Result, err error) {
	q := influxdb.Query{Command: cmd, Database: "test"}
	if response, err := client.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	}

	return res, err
}
func WritePoint() {
	bp, err := influxdb.NewBatchPoints(influxdb.BatchPointsConfig{
		Database: "test",
		//	Precision: "s", //精度，默认ns
	})
	if err != nil {
		log.Fatal(err)
	}

	fields := map[string]interface{}{
		"temp":  100 + float64(rand.Intn(10)),
		"humi":  100 + float64(rand.Intn(10)),
		"pm":    100 + float64(rand.Intn(10)),
		"light": 100 + float64(rand.Intn(10)),
	}
	t := time.Now()
	pt, err := influxdb.NewPoint("pm", nil, fields, t)
	//pt, err := influxdb.NewPoint("pm_monitoring", nil, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)
	err = client.Write(bp)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println("insert success")
}
