package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"

    "github.com/mzyl/bookbytes/bookbytes"
)

func main() {
    // for i in booklist.txt
    // gen book
    // (book.fulltext might actually be a minified version already)
    // write book.fulltext to ../minified/book.filename
    var line string
    var filepath string
//    var minified int
    var count int

    file, err := os.Open("doclist.txt")
    if err != nil {
        log.Fatal(err)
    }
    booklist := bufio.NewScanner(file)
    defer file.Close()

    for booklist.Scan() {
        line = booklist.Text()
        //filepath = "../library/htmlmirror/" + line[2:]
        filepath = "docs/" + line[2:]
        fmt.Println(filepath)
        book := bookbytes.GenerateBook(filepath, 0)

        newpath := "../library/minified/" + line[2:]
        pathsplit := strings.Split(newpath, "/")
        path := strings.Join(pathsplit[:len(pathsplit)-1], "/")
        fmt.Println()
        filename := pathsplit[len(pathsplit)-1]
        fmt.Println(filename)
        filenameslice := append([]string{"/"}, filename)
        fmt.Println("filename:", filenameslice)
        fmt.Println("path: ", path)
        fmt.Println(pathsplit[len(pathsplit)-1])

        os.MkdirAll(path, 0770)

        println("before")
        f, err := os.Create(newpath)
        if err != nil {
            fmt.Println(err) 
            log.Fatal(err)
        }
        defer f.Close()
        println("after")

        stripped := strings.TrimSpace(strings.Join(bookbytes.GetFullHtml(book), ""))
        
        w := bufio.NewWriter(f)
        //_, err = f.WriteString(bookbytes.GetFulltext(book))
        _, err = f.WriteString(stripped)
        if err != nil {
            log.Fatal(err)
        }
        w.Flush()

        println(filepath)
        bookbytes.BookPrinter(book)
        count++
        break
    }
    println(count)
}
