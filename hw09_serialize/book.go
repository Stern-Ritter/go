package main

import (
	"encoding/json"
	"strconv"
)

type Book struct {
	ID     uint64  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author,omitempty"`
	Year   uint32  `json:"year"`
	Size   uint32  `json:"size"`
	Rate   float32 `json:"rate"`
}

func NewBook(id uint64, title string, author string, year uint32, size uint32, rate float32) Book {
	return Book{
		ID:     id,
		Title:  title,
		Author: author,
		Year:   year,
		Size:   size,
		Rate:   rate,
	}
}

func (b *Book) MarshalJSON() ([]byte, error) {
	id := strconv.FormatUint(b.ID, 10)
	year := strconv.FormatUint(uint64(b.Year), 10)
	size := strconv.FormatUint(uint64(b.Size), 10)
	rate := strconv.FormatFloat(float64(b.Rate), 'f', 1, 32)

	type alias Book
	book := &struct {
		alias
		ID   string `json:"id"`
		Year string `json:"year"`
		Size string `json:"size"`
		Rate string `json:"rate"`
	}{
		alias: alias(*b),
		ID:    id,
		Year:  year,
		Size:  size,
		Rate:  rate,
	}

	return json.Marshal(book)
}

func (b *Book) UnmarshalJSON(data []byte) error {
	type alias Book

	book := &struct {
		*alias
		ID   string `json:"id"`
		Year string `json:"year"`
		Size string `json:"size"`
		Rate string `json:"rate"`
	}{
		alias: (*alias)(b),
	}

	if err := json.Unmarshal(data, book); err != nil {
		return err
	}

	if book.ID != "" {
		id, err := strconv.ParseUint(book.ID, 10, 64)
		if err != nil {
			return err
		}
		b.ID = id
	}

	if book.Year != "" {
		year, err := strconv.ParseUint(book.Year, 10, 32)
		if err != nil {
			return err
		}
		b.Year = uint32(year)
	}

	if book.Size != "" {
		size, err := strconv.ParseUint(book.Size, 10, 32)
		if err != nil {
			return err
		}
		b.Size = uint32(size)
	}

	if book.Rate != "" {
		rate, err := strconv.ParseFloat(book.Rate, 32)
		if err != nil {
			return err
		}
		b.Rate = float32(rate)
	}

	return nil
}
