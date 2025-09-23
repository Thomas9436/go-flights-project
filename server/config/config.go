package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string
	J1BaseURL  string // http://j-server1:4001
	J2BaseURL  string // http://j-server1:4001
}

func Load() Config {
	viper.AutomaticEnv()
	viper.SetDefault("SERVER_PORT", "3001")
	viper.SetDefault("JSERVER1_NAME", "j-server1")
	viper.SetDefault("JSERVER1_PORT", "4001")
	viper.SetDefault("JSERVER2_NAME", "j-server2")
	viper.SetDefault("JSERVER2_PORT", "4002")

	hostServ1 := viper.GetString("JSERVER1_NAME")
	portServ1 := viper.GetString("JSERVER1_PORT")

	hostServ2 := viper.GetString("JSERVER2_NAME")
	portServ2 := viper.GetString("JSERVER2_PORT")

	return Config{
		ServerPort: viper.GetString("SERVER_PORT"),
		J1BaseURL:  fmt.Sprintf("http://%s:%s", hostServ1, portServ1),
		J2BaseURL:  fmt.Sprintf("http://%s:%s", hostServ2, portServ2),
	}
}
