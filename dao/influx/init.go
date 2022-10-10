package influx

import (
	influxdb "github.com/influxdata/influxdb1-client/v2"
	"messageSystem/conf"
)

var client influxdb.Client

func Init(config *conf.Config) (err error) {
	client, err = influxdb.NewHTTPClient(influxdb.HTTPConfig{
		Addr:     config.Influx.Addr,
		Username: config.Influx.UserName,
		Password: config.Influx.Password,
	})
	return err
}
