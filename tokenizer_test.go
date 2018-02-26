package tokenizer

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	cases := []string{
		"Emerging after two hours of talks with Chinese President Xi Jinping, Trump said he doesn't fault China for taking advantage of differences between the way the two countries do business.",
		"I'm ok. but I still don't know who's my boy.",
	}
	for _, c := range cases {
		ret := TokenizePro(c)
		t.Log(ret)
	}
}
