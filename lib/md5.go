package lib

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(c string) string {
	h := md5.New()
	h.Write([]byte(c))
	return hex.EncodeToString(h.Sum(nil))
}
