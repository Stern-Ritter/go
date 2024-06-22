package main

import (
	"encoding/json"
	"testing"

	p "github.com/Stern-Ritter/go/hw09_serialize/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func compareBooks(b1, b2 *p.Book) bool {
	return proto.Equal(b1, b2)
}

func compareBookLists(b1, b2 *p.BookList) bool {
	return proto.Equal(b1, b2)
}

func TestBookMarshalJSON(t *testing.T) {
	type want struct {
		json string
		err  bool
	}

	tests := []struct {
		name string
		book Book
		want want
	}{
		{
			name: "should return correct json when all fields are filled",
			book: NewBook(1, "First title", "First author", 1932, 752, 7.8),
			want: want{
				json: `{"title":"First title","author":"First author","id":"1","year":"1932","size":"752","rate":"7.8"}`,
				err:  false,
			},
		},
		{
			name: "should return correct json with a rounded down rate value",
			book: NewBook(2, "Second title", "Second author", 1977, 346, 8.74),
			want: want{
				json: `{"title":"Second title","author":"Second author","id":"2","year":"1977","size":"346","rate":"8.7"}`,
				err:  false,
			},
		},
		{
			name: "should return correct json with a rounded up rate value",
			book: NewBook(3, "Third title", "Third author", 1988, 532, 9.75),
			want: want{
				json: `{"title":"Third title","author":"Third author","id":"3","year":"1988","size":"532","rate":"9.8"}`,
				err:  false,
			},
		},
		{
			name: "should return correct json without fields with tag omitempty and zero value",
			book: Book{},
			want: want{
				json: `{"title":"","id":"0","year":"0","size":"0","rate":"0.0"}`,
				err:  false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.book.MarshalJSON()
			if tt.want.err {
				require.Error(t, err, "json serialization error was expected")
			} else {
				require.NoError(t, err, "unexpected json serialization error")
			}

			assert.Equal(t, tt.want.json, string(got), "expected serialization result: %s, but got: %s",
				string(got), tt.want.json)
		})
	}
}

func TestBookListMarshalJSON(t *testing.T) {
	type want struct {
		json string
		err  bool
	}

	tests := []struct {
		name  string
		books []Book
		want  want
	}{
		{
			name: "should return correct json for a list of books #1",
			books: []Book{
				NewBook(1, "First title", "First author", 1932, 752, 7.8),
			},
			want: want{
				json: `[{"title":"First title","author":"First author","id":"1","year":"1932","size":"752","rate":"7.8"}]`,
				err:  false,
			},
		},
		{
			name: "should return correct json for a list of books #2",
			books: []Book{
				NewBook(1, "First title", "First author", 1932, 752, 7.8),
				NewBook(2, "Second title", "Second author", 1977, 346, 8.74),
			},
			want: want{
				json: `[{"title":"First title","author":"First author","id":"1","year":"1932","size":"752","rate":"7.8"},` +
					`{"title":"Second title","author":"Second author","id":"2","year":"1977","size":"346","rate":"8.7"}]`,
				err: false,
			},
		},
		{
			name:  "should return correct json for an empty list of books #3",
			books: []Book{},
			want: want{
				json: `[]`,
				err:  false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.books)
			if tt.want.err {
				require.Error(t, err, "json serialization error was expected")
			} else {
				require.NoError(t, err, "unexpected json serialization error")
			}

			assert.JSONEq(t, tt.want.json, string(got), "expected serialization result: %s, but got: %s",
				tt.want.json, string(got))
		})
	}
}

func TestBookUnmarshalJSON(t *testing.T) {
	type want struct {
		book Book
		err  bool
	}

	tests := []struct {
		name string
		data []byte
		want want
	}{
		{
			name: "should return correct book object when all fields are filled",
			data: []byte(`{"id":"1","title":"First title","author":"First author","year":"1932","size":"752","rate":"7.8"}`),
			want: want{
				book: NewBook(1, "First title", "First author", 1932, 752, 7.8),
				err:  false,
			},
		},
		{
			name: "should return correct book object when some of fields are missing",
			data: []byte(`{"id":"2","title":"Second title","author":"Second author","year":"1956","size":"856"}`),
			want: want{
				book: NewBook(2, "Second title", "Second author", 1956, 856, 0),
				err:  false,
			},
		},
		{
			name: "should return error when year has invalid value",
			data: []byte(`{"id":"3","title":"Third title","author":"Third author","year":"$3!","size":"856","rate":"7.8"}`),
			want: want{
				err: true,
			},
		},
		{
			name: "should return error when size has invalid value",
			data: []byte(`{"id":"3","title":"Third title","author":"Third author","year":"1956","size":"#122","rate":"7.8"}`),
			want: want{
				err: true,
			},
		},
		{
			name: "should return error when rate has invalid value",
			data: []byte(`{"id":"3","title":"Third title","author":"Third author","year":"1956","size":"856","rate":"0,76^"}`),
			want: want{
				err: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			book := Book{}
			err := json.Unmarshal(tt.data, &book)
			if tt.want.err {
				require.Error(t, err, "json deserialization error was expected")
			} else {
				require.NoError(t, err, "unexpected json deserialization error")
				assert.Equal(t, tt.want.book, book, "expected deserialization result: %v, but got: %v", tt.want.book, book)
			}
		})
	}
}

func TestBookListUnmarshalJSON(t *testing.T) {
	type want struct {
		books []Book
		err   bool
	}

	tests := []struct {
		name string
		data []byte
		want want
	}{
		{
			name: "should return correct book list when all fields are filled",
			data: []byte(`[{"id":"1","title":"First title","author":"First author","year":"1932","size":"752","rate":"7.8"},` +
				`{"id":"2","title":"Second title","author":"Second author","year":"1977","size":"346","rate":"8.7"}]`),
			want: want{
				books: []Book{
					NewBook(1, "First title", "First author", 1932, 752, 7.8),
					NewBook(2, "Second title", "Second author", 1977, 346, 8.7),
				},
				err: false,
			},
		},
		{
			name: "should return correct book list with missing optional fields",
			data: []byte(`[{"id":"3","title":"Third title","author":"Third author","year":"1988","size":"532"}]`),
			want: want{
				books: []Book{
					NewBook(3, "Third title", "Third author", 1988, 532, 0),
				},
				err: false,
			},
		},
		{
			name: "should return error when year has invalid value",
			data: []byte(`[{"id":"4","title":"Fourth title","author":"Fourth author","year":"$3!","size":"856","rate":"7.8"}]`),
			want: want{
				err: true,
			},
		},
		{
			name: "should return error when size has invalid value",
			data: []byte(`[{"id":"5","title":"Fifth title","author":"Fifth author","year":"1956","size":"#122","rate":"7.8"}]`),
			want: want{
				err: true,
			},
		},
		{
			name: "should return error when rate has invalid value",
			data: []byte(`[{"id":"6","title":"Sixth title","author":"Sixth author","year":"1956","size":"856","rate":"0,76^"}]`),
			want: want{
				err: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var books []Book
			err := json.Unmarshal(tt.data, &books)
			if tt.want.err {
				require.Error(t, err, "json deserialization error was expected")
			} else {
				require.NoError(t, err, "unexpected json deserialization error")
				assert.Equal(t, tt.want.books, books, "expected deserialization result: %v, but got: %v", tt.want.books, books)
			}
		})
	}
}

func TestBookProtobufSerialization(t *testing.T) {
	tests := []struct {
		name string
		book *p.Book
	}{
		{
			name: "should protobuf serialize a book #1",
			book: &p.Book{
				Id: 1, Title: "First title", Author: "First author", Year: 1932, Size: 752, Rate: 7.8,
			},
		},
		{
			name: "should protobuf serialize a book #2",
			book: &p.Book{
				Id: 2, Title: "Second title", Author: "Second author", Year: 1956, Size: 856, Rate: 8.7,
			},
		},
		{
			name: "should protobuf serialize a book #3" +
				" a book #3",
			book: &p.Book{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := proto.Marshal(tt.book)
			require.NoError(t, err, "unexpected protobuf serialization error")

			book := &p.Book{}
			err = proto.Unmarshal(data, book)
			require.NoError(t, err, "unexpected protobuf deserialization error")

			assert.True(t, compareBooks(tt.book, book), "expected deserialization result: %v, but got: %v", tt.book, book)
		})
	}
}

func TestBookListProtobufSerialization(t *testing.T) {
	tests := []struct {
		name     string
		bookList *p.BookList
	}{
		{
			name: "should protobuf serialize a book list #1",
			bookList: &p.BookList{
				Books: []*p.Book{
					{Id: 1, Title: "First title", Author: "First author", Year: 1932, Size: 752, Rate: 7.8},
				},
			},
		},
		{
			name: "should protobuf serialize a book list #2",
			bookList: &p.BookList{
				Books: []*p.Book{
					{Id: 1, Title: "First title", Author: "First author", Year: 1932, Size: 752, Rate: 7.8},
					{Id: 2, Title: "Second title", Author: "Second author", Year: 1945, Size: 500, Rate: 8.5},
				},
			},
		},
		{
			name: "should protobuf serialize a book list #3",
			bookList: &p.BookList{
				Books: []*p.Book{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := proto.Marshal(tt.bookList)
			require.NoError(t, err, "unexpected protobuf serialization error")

			bookList := &p.BookList{}
			err = proto.Unmarshal(data, bookList)
			require.NoError(t, err, "unexpected protobuf deserialization error")

			assert.True(t, compareBookLists(tt.bookList, bookList), "expected deserialization result: %v, but got: %v",
				&tt.bookList, bookList)
		})
	}
}
