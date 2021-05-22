package styles

import "sort"

var _ sort.Interface = make(styleSlice, 0)

type styleSlice []Style

// Len implements the sort.Interface interface.
func (ss styleSlice) Len() int {
	return len(ss)
}

// Less implements the sort.Interface interface.
func (ss styleSlice) Less(i, j int) bool {
	return ss[i].Weight() < ss[j].Weight()
}

// Swap implements the sort.Interface interface.
func (ss styleSlice) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}
