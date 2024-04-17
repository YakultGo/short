package base62

import "strings"

/*
A-Z 26个
a-z 26个
0-9 10个
不带上base64的两个符号+和/，刚好是62个字符
*/
const base62Str string = `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789`

// Encode 10进制转62进制
func Encode(seq uint64) string {
	if seq == 0 {
		return string(base62Str[0])
	}
	ret := make([]byte, 0)
	for seq > 0 {
		mod := seq % 62
		seq = seq / 62
		ret = append(ret, base62Str[mod])
	}
	reverse(ret)
	return string(ret)
}

func reverse(ret []byte) {
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
}

func Decode(seq string) uint64 {
	var ret uint64
	for _, i := range seq {
		ret = ret*62 + uint64(strings.Index(base62Str, string(i)))
	}
	return ret
}
