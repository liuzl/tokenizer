# Multilingual Tokenizer
## Introduction
Package `tokenizer` is a golang library for multilingual tokenization. It is based on the [segment](https://github.com/liuzl/segment) package of [blevesearch](https://github.com/blevesearch), whose implementation is based on the description at [Unicode Standard Annex #29](http://www.unicode.org/reports/tr29/).
## Install
```sh
go get github.com/liuzl/tokenizer
```
## Usage
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
