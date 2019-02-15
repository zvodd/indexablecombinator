package indexablecombinator

import (
	"testing"
)

func TestIndexs(t *testing.T) {
	p := &Combinator{[]int{4, 4, 4}}
	total := p.TotalLen()
	t.Log("Sizes:", p.Sizes)
	t.Log("Total:", total)
	for i := uint64(0); i < total; i++ {
		t.Log(p.IndexProduct(i))
	}
}

func TestStrings(t *testing.T) {
	sp := NewStringCombinator([][]string{
		{"1", "2", "3"},
		{"a", "b", "c"},
		{"x", "y", "z"},
	})

	total := sp.P.TotalLen()
	t.Log("Sizes:", sp.P.Sizes)
	t.Log("Total:", total)
	for i := uint64(0); i < total; i++ {
		t.Log(sp.IndexItem(i))
	}
}
