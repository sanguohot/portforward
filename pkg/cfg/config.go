package cfg

import (
	"github.com/sanguohot/log/v2"
	"github.com/spf13/viper"
	"time"
)

const (
	NetworkTcp  = "tcp"
	NetworkUdp  = "udp"
	ReadTimeout = time.Second * 5
	DialTimeout = time.Second * 2
	BufferSize  = 4096
)

var Config ConfigStruct

type ConfigForward struct {
	DstAddr string `json:"dstAddr"`
	SrcAddr string `json:"srcAddr"`
	Network string `json:"network"`
}

type ConfigStruct struct {
	Forwards []ConfigForward `json:"forwards"`
}

func LoadConfig() {
	filePath := "./etc/config.yaml"
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Logger.Fatal(err.Error())
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Logger.Fatal(err.Error())
	}
	log.Sugar.Info(Config)
}

func init() {
	viper.SetConfigType("yaml")
}
