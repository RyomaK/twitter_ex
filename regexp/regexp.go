package regexp

import (
	"regexp"
	"strings"
)

//aタグをつけて返す
func ChangeURL(text string) string {
	repUrl := regexp.MustCompile(`https?://[\w/:%#\$&\?\(\)~\.=\+\-]+`)
	urls := repUrl.FindAllString(text, -1)
	for _, url := range urls {
		link := "<a href=\"" + url + "\">" + url + "</a>"
		text = strings.Replace(text, url, link, -1)
	}
	return text
}

func IsOnlyJapanese(str string) bool {
	rep := regexp.MustCompile(`[!-/:-~]`)
	if rep.MatchString(str) {
		return false
	}
	return true
}
