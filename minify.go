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
    var line string
    var filename string
//    var minified int
    var count int

    file, err := os.Open("booklist.txt")
    if err != nil {
        log.Fatal(err)
    }
    booklist := bufio.NewScanner(file)
    defer file.Close()

    for booklist.Scan() {
        line = booklist.Text()
        filename = "../library/htmlmirror/" + line[2:]
        book := bookbytes.GenerateBook(filename, 0)
        
        println("before")
        f, err := os.Create("../library/minified/" + line[2:])
        if err != nil {
            log.Fatal(err)
        }
        defer f.Close()
        println("after")
        
        w := bufio.NewWriter(f)
        _, err = f.WriteString(bookbytes.GetFulltext(book))
        if err != nil {
            log.Fatal(err)
        }
        w.Flush()

        println(filename)
        bookbytes.BookPrinter(book)
        count++
        break
    }
    println(count)
}
