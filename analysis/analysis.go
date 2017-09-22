package analysis

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/ikawaha/kagome/tokenizer"

	"github.com/RyomaK/twitter_ex/regexp"
)

type Word struct {
	Word      string
	WordClass string
}

//形態素解析
func analisys(str string) *[]Word {
	t := tokenizer.New()
	tokens := t.Tokenize(str)
	words := &[]Word{}

	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			// BOS: Begin Of Sentence, EOS: End Of Sentence.
			continue
		}
		status := token.Features()
		if status[1] == "数" {
			*words = append(*words, Word{token.Surface, status[1]})
		} else if status[0] == "名詞" {
			*words = append(*words, Word{status[7], status[0]})
		} else {
			*words = append(*words, Word{token.Surface, status[0]})
		}

	}
	return words
}

//韻を生成
func rhyme(str string) string {
	doc, err := goquery.NewDocument("https://kujirahand.com/web-tools/Words.php?m=boin-search&opt=comp&key=" + str)
	if err != nil {
		fmt.Errorf("err rhyme %v", err)
		return str
	}
	var rhymeWord string
	doc.Find("rb").EachWithBreak(func(_ int, s *goquery.Selection) bool {
		rhymeWord = s.Text()
		return false
	})

	if rhymeWord != "" {
		return rhymeWord
	}
	return str
}

//韻を返す
func GetRhyme(str string) string {
	if !regexp.IsOnlyJapanese(str) {
		return str
	}
	words := analisys(str)
	rhymeSentence := ""
	for _, v := range *words {
		if v.WordClass == "名詞" {
			rhymeSentence += rhyme(v.Word)
		} else {
			rhymeSentence += v.Word
		}
	}
	return rhymeSentence
}
