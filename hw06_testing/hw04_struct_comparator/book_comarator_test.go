package hw04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookComparator_Compare(t *testing.T) {
	tests := []struct {
		name       string
		comparator *BookComparator
		firstBook  Book
		secondBook Book
		want       bool
	}{
		{
			name:       "compare book by year: should return true when first book year is greater than second book year",
			comparator: NewBookComparator(CompareBookByYear),
			firstBook:  NewBook(1, "title", "author", 2002, 401, 8.0),
			secondBook: NewBook(1, "title", "author", 2001, 401, 8.0),
			want:       true,
		},
		{
			name:       "compare book by year: should return true when first book year is equal second book year",
			comparator: NewBookComparator(CompareBookByYear),
			firstBook:  NewBook(1, "title", "author", 2002, 401, 8.0),
			secondBook: NewBook(1, "title", "author", 2002, 401, 8.0),
			want:       true,
		},
		{
			name:       "compare book by year: should return false when first book year is less than second book year",
			comparator: NewBookComparator(CompareBookByYear),
			firstBook:  NewBook(1, "title", "author", 2001, 401, 8.0),
			secondBook: NewBook(1, "title", "author", 2002, 401, 8.0),
			want:       false,
		},
		{
			name:       "compare book by size: should return true when first book size is greater than second book size",
			comparator: NewBookComparator(CompareBookBySize),
			firstBook:  NewBook(1, "title", "author", 2002, 401, 8.0),
			secondBook: NewBook(1, "title", "author", 2002, 400, 8.0),
			want:       true,
		},
		{
			name:       "compare book by size: should return true when first book size is equal second book size",
			comparator: NewBookComparator(CompareBookBySize),
			firstBook:  NewBook(1, "title", "author", 2002, 401, 8.0),
			secondBook: NewBook(1, "title", "author", 2002, 401, 8.0),
			want:       true,
		},
		{
			name:       "compare book by size: should return false when first book size is less than second book size",
			comparator: NewBookComparator(CompareBookBySize),
			firstBook:  NewBook(1, "title", "author", 2002, 400, 8.0),
			secondBook: NewBook(1, "title", "author", 2002, 401, 8.0),
			want:       false,
		},
		{
			name:       "compare book by rate: should return true when first book rate is greater than second book rate",
			comparator: NewBookComparator(CompareBookByRate),
			firstBook:  NewBook(1, "title", "author", 2002, 401, 8.0),
			secondBook: NewBook(1, "title", "author", 2002, 401, 7.9),
			want:       true,
		},
		{
			name:       "compare book by rate: should return true when first book rate is equal second book rate",
			comparator: NewBookComparator(CompareBookByRate),
			firstBook:  NewBook(1, "title", "author", 2002, 401, 8.0),
			secondBook: NewBook(1, "title", "author", 2002, 401, 8.0),
			want:       true,
		},
		{
			name:       "compare book by rate: should return true when first book rate is less than second book rate",
			comparator: NewBookComparator(CompareBookByRate),
			firstBook:  NewBook(1, "title", "author", 2002, 401, 7.9),
			secondBook: NewBook(1, "title", "author", 2002, 401, 8.0),
			want:       false,
		},
		{
			name:       "compare book by default: should return true when first book id is greater than second book id",
			comparator: NewBookComparator(BookCompareType(42)),
			firstBook:  NewBook(2, "title", "author", 2002, 401, 8.0),
			secondBook: NewBook(1, "title", "author", 2002, 401, 8.0),
			want:       true,
		},
		{
			name:       "compare book by default: should return true when first book id is equal second book id",
			comparator: NewBookComparator(BookCompareType(42)),
			firstBook:  NewBook(1, "title", "author", 2002, 401, 8.0),
			secondBook: NewBook(1, "title", "author", 2002, 401, 8.0),
			want:       true,
		},
		{
			name:       "compare book by default: should return true when first book id is less than second book id",
			comparator: NewBookComparator(BookCompareType(42)),
			firstBook:  NewBook(1, "title", "author", 2002, 401, 8.0),
			secondBook: NewBook(2, "title", "author", 2002, 401, 8.0),
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.comparator.Compare(tt.firstBook, tt.secondBook))
		})
	}
}
