package main

type BookCompareType int32

const (
	CompareBookByYear BookCompareType = iota
	CompareBookBySize
	CompareBookByRate
)

type BookComparator struct {
	compareType BookCompareType
}

func NewBookComparator(compareType BookCompareType) *BookComparator {
	return &BookComparator{compareType: compareType}
}

func (c BookComparator) Compare(a, b Book) bool {
	switch c.compareType {
	case CompareBookByYear:
		return a.Year() >= b.Year()
	case CompareBookBySize:
		return a.Size() >= b.Size()
	case CompareBookByRate:
		return a.Rate() >= b.Rate()
	default:
		return a.ID() >= b.ID()
	}
}
