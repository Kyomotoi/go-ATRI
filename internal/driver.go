package internal

import "github.com/Kyomotoi/go-ATRI/internal/gocqhttp"

func InitDriver(v string) error {
	return gocqhttp.InitGoCQHTTP(v)
}
