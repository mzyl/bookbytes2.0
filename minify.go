package main

import (
    "bufio"
    "log"
    "os"

    "github.com/mzyl/bookbytes/bookbytes"
)

func main() {
    // for i in booklist.txt
    // gen book
    // (book.fulltext might actually be a minified version already)
    // write book.fulltext to ../minified/book.filename
    var filename string
    var count int

    file, err := os.Open("booklist.txt")
    if err != nil {
        log.Fatal(err)
    }
    booklist := bufio.NewScanner(file)
    defer file.Close()

    for booklist.Scan() {
        filename = booklist.Text()
        book := bookbytes.GenerateBook(filename, 0)
        println(filename)
        bookbytes.BookPrinter(book)
        count++
    }
    println(count)
}