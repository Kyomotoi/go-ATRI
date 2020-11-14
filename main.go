package main

import (
	zero "github.com/wdvxdr1123/ZeroBot"
)

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
