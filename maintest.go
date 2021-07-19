package main

import (
	"github.com/mzyl/bookbytes/bookbytes"
)

func main() {
    //filename := bookbytes.GetFile("booklist.txt")
	//book := bookbytes.GenerateBook(filename, 0)
    book := bookbytes.Init()
	bookbytes.BookPrinter(book)
	//bookbytes.SplitTextToken(book.fullhtml)
}
