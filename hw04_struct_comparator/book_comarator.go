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
		return a.GetYear() >= b.GetYear()
	case CompareBookBySize:
		return a.GetSize() >= b.GetSize()
	case CompareBookByRate:
		return a.GetRate() >= b.GetRate()
	default:
		return a.GetID() >= b.GetID()
	}
}
