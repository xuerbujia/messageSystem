package conf

import "github.com/spf13/viper"

var Conf = new(Config)

type Config struct {
	Influx InfluxConfig `mapstructure:"influx"`
}
type InfluxConfig struct {
	Addr     string `mapstructure:"addr"`
	UserName string `mapstructure:"user_name"`
	Password string `mapstructure:"password"`
}

func Init() (err error) {
	viper.SetConfigFile("./config.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	return viper.Unmarshal(Conf)
}
