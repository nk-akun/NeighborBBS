package util

import "testing"

func Test_RandomAvatarURL(t *testing.T) {
	ans := RandomAvatarURL("danchunxiao2b@163.com")
	t.Errorf(ans)
}
