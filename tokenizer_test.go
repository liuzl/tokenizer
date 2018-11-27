package tokenizer

import (
	"strings"
	"testing"
)

func TestTokenize(t *testing.T) {
	cases := []string{
		`I'm ok. but I still don't know who's my boy.`,
		`This is Zhanliang's book. the boys' books.`,
		`自建房2樓3室2廳1衛1廚92.00㎡戶型圖，2,123km到北京`,
		`自建房2樓3室2廳1衛1廚92.00㎡戶型圖，92km到北京`,
		`自建房2樓3室2廳1衛1廚92.00㎡戶型圖，+92.54km到北京`,
		`山西省煤炭地质１４８勘查院煤层气工程处`,
		`北京ทันทุกเหตุการ有限公司`,
		`23:45:31.965.805`,
		`23:45:31.965.805cm`,
	}
	for _, c := range cases {
		ret := TokenizePro(c)
		var words []string
		var norms []string
		for _, term := range ret {
			words = append(words, term.Text)
			norms = append(norms, term.Norm)
		}
		t.Log(strings.Join(words, "/"))
		t.Log(strings.Join(norms, "/"))
	}
}
