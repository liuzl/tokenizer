package tokenizer

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/liuzl/segment"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/width"
)

var (
	NumberWithUnitPattern = `^(\d*\.?\d+|\d{1,3}(?:,\d{3})+)([a-zA-Z]{1,3})$`
	NumberWithUnitRegex   = regexp.MustCompile(NumberWithUnitPattern)

	trans = transform.Chain(
		norm.NFD,
		transform.RemoveFunc(func(r rune) bool {
			return unicode.Is(unicode.Mn, r)
		}),
		norm.NFC)

	replacer = strings.NewReplacer(
		`｡`, `.`, // half period in Chinese
		`。`, `.`, // full period in Chinese
		`【`, `[`,
		`】`, `]`,
		`“`, `"`,
		`”`, `"`,
		`‘`, `'`,
		`’`, `'`,
		`—`, `-`,
		`〔`, `{`,
		`〕`, `}`,
		`《`, `<`,
		`》`, `>`,
	)
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
	res, _, err := transform.String(trans, text)
	if err != nil {
		return nil
	}
	// full to half
	ntext := replacer.Replace(width.Narrow.String(res))
	fmt.Println(text, len([]rune(text)))
	fmt.Println(ntext, len([]rune(ntext)))

	runes := []rune(text)

	var ret []*Token
	seg := segment.NewSegmenterDirect([]byte(ntext))
	p := 0
	for seg.Segment() {
		token := seg.Text()
		rtoken := []rune(token)
		length := len(rtoken)
		raw := runes[p : p+length]
		p += length

		lowered := strings.ToLower(token)
		items, has := Contractions["eng"][lowered]
		switch {
		case has: // deal with contractions
			pos := 0
			for i, term := range items.Terms {
				rterm := []rune(term)
				//txt := token[pos : pos+len(term)]
				txt := raw[pos : pos+len(rterm)]
				norm := string(rtoken[pos : pos+len(rterm)])
				if len(items.Norms) > 0 {
					norm = items.Norms[0][i]
				}
				ret = append(ret, &Token{Text: string(txt), Norm: norm})
				pos += len(rterm)
			}
			// https://en.wikipedia.org/wiki/English_possessive
		case strings.HasSuffix(lowered, `'s`):
			ret = append(ret, &Token{
				Text: string(raw[:len(raw)-2]), Norm: token[:len(token)-2]})
			ret = append(ret, &Token{
				Text: string(raw[len(raw)-2:]), Norm: token[len(token)-2:]})
		case NumberWithUnitRegex.MatchString(token):
			ss := NumberWithUnitRegex.FindStringSubmatch(token)
			j := len([]rune(ss[1]))
			ret = append(ret, &Token{Text: string(raw[:j]), Norm: ss[1]})
			ret = append(ret, &Token{Text: string(raw[j:]), Norm: ss[2]})
		default:
			ret = append(ret, &Token{Text: string(raw), Norm: token})
		}
	}
	return ret
}
