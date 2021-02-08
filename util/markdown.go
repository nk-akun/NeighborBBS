package util

import (
	"sync"

	"github.com/88250/lute"
)

var (
	engine *lute.Lute
	once   sync.Once
)

func getEngine() *lute.Lute {
	once.Do(func() {
		engine = lute.New(func(lute *lute.Lute) {
			lute.SetToC(true)
			lute.SetGFMTaskListItem(true)
		})
	})
	return engine
}

// ToHTML ...
func ToHTML(markdownStr string) string {
	if IsBlank(markdownStr) {
		return ""
	}
	return getEngine().MarkdownStr("", markdownStr)
}
