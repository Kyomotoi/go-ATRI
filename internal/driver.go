package internal

import "github.com/Kyomotoi/go-ATRI/internal/gocqhttp"

func InitDriver(version string, account int64, password string, host string, port string, proc int) error {
	return gocqhttp.InitGoCQHTTP(version, account, password, host, port, proc)
}
