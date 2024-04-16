package extract

import (
	"crypto/md5"
	"encoding/hex"
)

// Sum 返回 data 的 md5 值
func Sum(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
