package main

import (
	zero "github.com/wdvxdr1123/ZeroBot"
)

func main() {

	zero.Run(zero.Option{
		Host:          "127.0.0.1",
		Port:          "8080",
		AccessToken:   "",
		NickName:      []string{"ATRI", "atri", "亚托莉", "アトリ"},
		CommandPrefix: "",
		SuperUsers:    []string{""},
	})
	select {}
}
