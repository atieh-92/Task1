package library

import "strings"

type Library struct {
	capacity int
	books    map[string]string // bookName(lowercased) -> borrower name, empty string if available
	count    int               // total number of books added (not returned)
}

func normalize(name string) string {
	return strings.ToLower(name)
}

func NewLibrary(capacity int) *Library {
	return &Library{
		capacity: capacity,
		books:    make(map[string]string),
		count:    0,
	}
}

func (library *Library) AddBook(name string) string {
	key := normalize(name)

	if _, exists := library.books[key]; exists {
		return "the book is already in the library"
	}

	if library.count >= library.capacity {
		return "Not enough capacity"
	}

	library.books[key] = ""
	library.count++
	return "OK"
}

func (library *Library) BorrowBook(bookName, personName string) string {
	key := normalize(bookName)

	borrower, exists := library.books[key]
	if !exists {
		return "the book has not been borrowed"
	}

	if borrower != "" {
		return "the book is already been borrowed by " + borrower
	}

	library.books[key] = personName
	return "OK"
}

func (library *Library) ReturnBook(bookName string) string {
	key := normalize(bookName)

	borrower, exists := library.books[key]
	if !exists {
		return "the book has not been borrowed"
	}

	if borrower == "" {
		return "the book has not been borrowed"
	}

	library.books[key] = ""
	return "OK"
}
