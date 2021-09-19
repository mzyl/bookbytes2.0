package main

import (
	"github.com/mzyl/bookbytes/bookbytes"
)

func main() {
    filename := bookbytes.GetFile("compressedbooklist.txt")
	//book := bookbytes.GenerateBook(filename, 0)
    paragraph, index := bookbytes.Init(filename)
    book := bookbytes.GenerateBook(filename, index)
	bookbytes.BookPrinter(book)
    println(paragraph)
	//bookbytes.SplitTextToken(book.fullhtml)
}
