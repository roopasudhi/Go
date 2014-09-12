package md5

import (
	"crypto/md5"
	"encoding/hex"

)


func Md5sum(content []byte) string {
	h := md5.New()
	h.Write(content)
	return hex.EncodeToString(h.Sum(nil))
}
