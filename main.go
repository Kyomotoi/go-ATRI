package main

import (
	"github.com/Kyomotoi/go-ATRI/models"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	zero "github.com/wdvxdr1123/ZeroBot"
)

func init() {
	logrus.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "%time% [%lvl%] ATRI | %msg% \n",
	})

	models.LoadConfig("config.json")
}

func main() {

	zero.Run(zero.Option{
		Host:          "127.0.0.1",
		Port:          "51817",
		AccessToken:   "",
		NickName:      []string{"ATRI", "atri", "亚托莉", "アトリ"},
		CommandPrefix: "",
		SuperUsers:    []string{"1172294279"},
	})
	select {}
}
