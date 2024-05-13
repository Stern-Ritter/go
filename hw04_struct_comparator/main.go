package main

import "fmt"

func main() {
	firstBook := NewBook(1, "First book title", "First book author", 1990, 220, 8.1)
	secondBook := NewBook(2, "Second book title", "Second book author", 1988, 432, 8.1)

	compareByYearComparator := NewBookComparator(CompareBookByYear)
	compareBySizeComparator := NewBookComparator(CompareBookBySize)
	compareByRateComparator := NewBookComparator(CompareBookByRate)

	fmt.Printf("By year: %v\n", compareByYearComparator.Compare(firstBook, secondBook))
	fmt.Printf("By size: %v\n", compareBySizeComparator.Compare(firstBook, secondBook))
	fmt.Printf("By rate: %v\n", compareByRateComparator.Compare(firstBook, secondBook))
}
