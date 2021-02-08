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

// SubString ...
func SubString(str string, start int, end int) string {
	if start >= len(str) || end >= len(str) || start > end || end < 0 {
		return ""
	}
	return string([]rune(str)[start:end])
}

// NowTimestamp return timestamp of now
func NowTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}

// Timestamp return timestamp of t
func Timestamp(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// MinInt64 return min(a,b)
func MinInt64(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// MaxInt64 return max(a,b)
func MaxInt64(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// MinInt return min(a,b)
func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxInt return max(a,b)
func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
