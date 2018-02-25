package tokenizer

import (
	"github.com/liuzl/segment"
)

func Tokenize(text string) []string {
	var ret []string
	seg := segment.NewSegmenterDirect([]byte(text))
	for seg.Segment() {
		ret = append(ret, seg.Text())
	}
	return ret
}
