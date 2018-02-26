# Multilingual Tokenizer
## Introduction
Package `tokenizer` is a golang library for multilingual tokenization. It is based on the [segment](https://github.com/liuzl/segment) package of [blevesearch](https://github.com/blevesearch), whose implementation is based on the description at [Unicode Standard Annex #29](http://www.unicode.org/reports/tr29/).
## Usage
```sh
go get github.com/liuzl/tokenizer
```
```go
package main

import (
    "fmt"
    "github.com/liuzl/tokenizer"
)

func main() {
    c := `Life is like a box of chocolates. You never know what you're gonna get.`
    var ret []string = tokenizer.Tokenize(c)
    for _, term := range ret {
        fmt.Println(term)
    }
}
```
## Implementation Details
1. Segment UTF-8 string as described at [Unicode Standard Annex #29](http://www.unicode.org/reports/tr29/).
2. Deal with [English contractions](https://en.wikipedia.org/wiki/Wikipedia:List_of_English_contractions).
3. Deal with English possessives.
## Licence
This package is licenced under the Apache License 2.0.
