package tokenizer

import (
	"strings"
)

type Items struct {
	Terms []string
	Norms [][]string
}
type Dict map[string]*Items

var Contractions map[string]Dict = make(map[string]Dict)

func init() {
	initEngContractions()
}

func initEngContractions() {
	m := make(Dict)
	for _, line := range strings.Split(EngContractions, "\n") {
		// all lowered
		items := strings.Split(strings.ToLower(line), "\t")
		if len(items) < 3 {
			continue
		}
		value := &Items{Terms: strings.Split(items[1], " ")}
		l := len(value.Terms)
		invalid := false
		for _, n := range items[2:] {
			terms := strings.Split(n, " ")
			if len(terms) != l {
				invalid = true
				break
			}
			value.Norms = append(value.Norms, terms)
		}
		if invalid {
			continue
		}
		m[items[0]] = value
	}
	Contractions["eng"] = m
}
