package main

import (
	"github.com/mzyl/bookbytes/bookbytes"
)

func main() {
    filename := bookbytes.GetFile("booklist.txt")
	book := bookbytes.GenerateBook(filename, 0)
	bookbytes.BookPrinter(book)
	//bookbytes.SplitTextToken(book.fullhtml)
}
