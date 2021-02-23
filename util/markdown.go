package util

import (
	"strings"
	"sync"

	"github.com/88250/lute"
	"github.com/PuerkitoBio/goquery"
	"github.com/nk-akun/NeighborBBS/logs"
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

// MarkdownToHTML ...
func MarkdownToHTML(markdownStr string) string {
	if IsBlank(markdownStr) {
		return ""
	}
	return getEngine().MarkdownStr("", markdownStr)
}

// GetHTMLText ...
func GetHTMLText(html string) string {
	txt, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		logs.Logger.Errorf("从html读取文本出错", txt)
	}
	return txt.Text()
}
