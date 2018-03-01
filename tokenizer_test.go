package tokenizer

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	cases := []string{
		`I'm ok. but I still don't know who's my boy.`,
		`This is Zhanliang's book. the boys' books.`,
		`自建房2樓3室2廳1衛1廚92.00㎡戶型圖，2,123km到北京`,
		`自建房2樓3室2廳1衛1廚92.00㎡戶型圖，92km到北京`,
		`自建房2樓3室2廳1衛1廚92.00㎡戶型圖，92.54km到北京`,
	}
	for _, c := range cases {
		ret := Tokenize(c)
		t.Log(ret)
	}
}
