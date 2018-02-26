package tokenizer

import (
	"fmt"
	"github.com/liuzl/segment"
	"strings"
)

type Token struct {
	Text string
	Norm string
}

func (self *Token) String() string {
	return fmt.Sprintf("(text:%s/norm:%s)", self.Text, self.Norm)
}

func Tokenize(text string) []string {
	var ret []string
	for _, term := range TokenizePro(text) {
		ret = append(ret, term.Text)
	}
	return ret
}

func TokenizePro(text string) []*Token {
	var ret []*Token
	seg := segment.NewSegmenterDirect([]byte(text))
	for seg.Segment() {
		text := seg.Text()
		lowered := strings.ToLower(text)
		if items, has := Contractions["eng"][lowered]; has {
			pos := 0
			for i, term := range items.Terms {
				txt := text[pos : pos+len(term)]
				norm := ""
				if len(items.Norms) > 0 {
					norm = items.Norms[0][i]
				}
				ret = append(ret, &Token{Text: txt, Norm: norm})
				pos += len(term)
			}
		} else {
			// https://en.wikipedia.org/wiki/English_possessive
			if strings.HasSuffix(lowered, `'s`) {
				ret = append(ret, &Token{Text: text[:len(text)-2]})
				ret = append(ret, &Token{Text: text[len(text)-2:]})
			} else {
				ret = append(ret, &Token{Text: text})
			}
		}
	}
	return ret
}
