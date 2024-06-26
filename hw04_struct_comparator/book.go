package main

type Book struct {
	id     int64
	title  string
	author string
	year   int32
	size   int32
	rate   float32
}

func NewBook(id int64, title string, author string, year int32, size int32, rate float32) Book {
	return Book{id, title, author, year, size, rate}
}

func (b *Book) ID() int64 {
	return b.id
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) Year() int32 {
	return b.year
}

func (b *Book) Size() int32 {
	return b.size
}

func (b *Book) Rate() float32 {
	return b.rate
}

func (b *Book) SetID(id int64) {
	b.id = id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetYear(year int32) {
	b.year = year
}

func (b *Book) SetSize(size int32) {
	b.size = size
}

func (b *Book) SetRate(rate float32) {
	b.rate = rate
}
