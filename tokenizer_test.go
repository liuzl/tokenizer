package tokenizer

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	cases := []string{
		`I'm ok. but I still don't know who's my boy.`,
		`This is Zhanliang's book. the boys' books.`,
	}
	for _, c := range cases {
		ret := TokenizePro(c)
		t.Log(ret)
	}
}
