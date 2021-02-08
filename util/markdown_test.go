package util

import (
	"testing"
)

func Test_ToHTML(t *testing.T) {
	str := "## 你是谁"

	t.Errorf(MarkdownToHTML(str))
}
