package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

var (
	avatarURLTemplate = "https://robohash.org/%s?set=set%d"
	avatarGroupNum    = []int{1, 2, 3, 4}
)

// RandomAvatarURL return random profile photo url
// str to md5
// thanks to https://robohash.org
func RandomAvatarURL(str string) string {
	data := []byte(str)
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	return fmt.Sprintf(avatarURLTemplate, hex.EncodeToString(cipherStr), str[0]%4+1)
}

// NowTimestamp return timestamp of now
func NowTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}

// Timestamp return timestamp of t
func Timestamp(t time.Time) int64 {
	return t.UnixNano() / 1e6
}
