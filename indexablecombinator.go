package indexablecombinator

type Combinator struct {
	Sizes []int
}

func (p *Combinator) TotalLen() uint64 {
	t := uint64(1)
	for _, x := range p.Sizes {
		t *= uint64(x)
	}
	return t
}

func (p *Combinator) IndexProduct(index uint64) (indexes []int) {
	rindexes := make([]int, 0)
	x := index
	for i := len(p.Sizes) - 1; i >= 0; i-- {
		base := uint64(p.Sizes[i])
		amputatedIndex := int(x % base)
		rindexes = append(rindexes, amputatedIndex)
		x = x / base
	}
	indexes = make([]int, len(rindexes))
	for i := 0; i < len(p.Sizes); i++ {
		j := len(p.Sizes) - 1 - i
		indexes[i] = rindexes[j]
	}
	return
}

type StringCombinator struct {
	P     *Combinator
	Lists [][]string
}

func NewStringCombinator(strings [][]string) (sp *StringCombinator) {
	s := make([]int, len(strings))
	for i := range s {
		s[i] = len(strings[i])
	}
	p := &Combinator{s}
	sp = &StringCombinator{p, strings}
	return sp
}

func (sp *StringCombinator) IndexIndexes(index []int) []string {
	rlist := make([]string, len(sp.Lists))
	for i := 0; i < len(index); i++ {
		rlist[i] = sp.Lists[i][index[i]]
	}
	return rlist
}

func (sp *StringCombinator) IndexItem(index uint64) []string {
	return sp.IndexIndexes(sp.P.IndexProduct(index))
}
