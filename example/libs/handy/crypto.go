package handy

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(s []byte) string {
	sum := md5.Sum(s)
	return hex.EncodeToString(sum[:])
}
