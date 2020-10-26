package util

import "time"

// NowTimestamp return timestamp of now
func NowTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}
