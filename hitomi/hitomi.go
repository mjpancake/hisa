package hitomi

import (
	"log"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/importcjj/sensitive"
)

var filter *sensitive.Filter

func Init(dictPath string) {
	filter = sensitive.New()
	err := filter.LoadWordDict(dictPath)
	if err != nil {
		log.Println("load sensitive dict", err)
	}
}

func CheckName(name string) bool {
	if strings.HasPrefix(name, "ⓝ") {
		return false
	}

	if found, _ := filter.FindIn(name); found {
		return false
	}

	p := bluemonday.UGCPolicy()
	ans := p.Sanitize(name)
	if (ans != name) {
		return false
	}

	return true
}

func Filter(s string) string {
	return filter.Replace(s, 42)
}
