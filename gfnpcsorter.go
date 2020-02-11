package main

import (
	"sort"
)

type lessFunc func(p1, p2 *GFNPC) bool

type GFNPCSorter struct {
	gfnpcs []GFNPC
	less   []lessFunc
}

func (gs *GFNPCSorter) Sort(gfnpcs []GFNPC) {
	gs.gfnpcs = gfnpcs
	sort.Sort(gs)
}

func OrderedBy(less ...lessFunc) *GFNPCSorter {
	return &GFNPCSorter{
		less: less,
	}
}

func (gs *GFNPCSorter) Len() int {
	return len(gs.gfnpcs)
}

func (gs *GFNPCSorter) Swap(i, j int) {
	gs.gfnpcs[i], gs.gfnpcs[j] = gs.gfnpcs[j], gs.gfnpcs[i]
}

func (gs *GFNPCSorter) Less(i, j int) bool {
	p, q := &gs.gfnpcs[i], &gs.gfnpcs[j]

	var k int
	for k = 0; k < len(gs.less)-1; k++ {
		less := gs.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return gs.less[k](p, q)
}
