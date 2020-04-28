package datastore

import (
	"github.com/matt-FFFFFF/bookdata-api/loader"
	"strings"
)

// Books is the memory-backed datastore used by the API
// It contains a single field 'Store', which is (a pointer to) a slice of loader.BookData struct pointers
type Books struct {
	Store *[]*loader.BookData `json:"store"`
}

// Initialize is the method used to populate the in-memory datastore.
// At the beginning, this simply returns a pointer to the struct literal.
// You need to change this to load data from the CSV file
func (b *Books) Initialize() {
	loader.LoadCSV()
	b.Store = &loader.BooksLiteral
}

// GetAllBooks returns the entire dataset, subjet to the rudimentary limit & skip parameters
func (b *Books) GetAllBooks(limit, skip int) *[]*loader.BookData {
	if limit == 0 || limit > len(*b.Store) {
		limit = len(*b.Store)
	}
	ret := (*b.Store)[skip:limit]
	return &ret
}

// GetBooksByAuthor will return books filtered by the given author
func (b *Books) GetBooksByAuthor(author string) *[]*loader.BookData {
	tmp := (*b.Store)[:0]
	for _, book := range (*b.Store) {
		if strings.Contains(book.Authors, author) {
			tmp = append(tmp,book)
		}
	}
	return &tmp
}


