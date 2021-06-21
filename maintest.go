package main

import (
	"github.com/mzyl/bookbytes/bookbytes"
)

func main() {
	book := bookbytes.NewBook()
	bookbytes.BookPrinter(book)
    //bookbytes.SplitTextToken(book.fullhtml)
}
