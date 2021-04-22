package bookbytes

import (
  "os"
  "log"
  "time"
  "bufio"
  "math/rand"
  "path/filepath"
)

var Files []string

func GetFile() (filename string) {
  var randomfile int
  rand.Seed(time.Now().UnixNano())
  
  if Files == nil {
    //root := "./books" // for testing
    root := "../library/htmlmirror"
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
      if filepath.Ext(path) == ".htm" || filepath.Ext(path) == ".html" {
        Files = append(Files, path)
        println(path)
      }
      return nil
    })
    if err != nil {
      panic(err)
    }
  }
  
  for range Files {
    randomfile = rand.Intn(len(Files))
    println(randomfile)
    if randomfile != 0 {
      filename = Files[randomfile]
      break
    }
  }

  println(len(Files))
  println()
  println("./" + filename)
  return "./" + filename
  //return "../library/htmlmirror/3/1/2/0/31200/31200-h/31200-h.htm"
}

func GetContents(filename string) (text []string) {
  file, err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
  } 
  defer file.Close()
  
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    text = append(text, scanner.Text())
  }
  return
}

func NewParagraph(book Book) (index int) {
  text := book.booktext
  rand.Seed(time.Now().UnixNano())
  var randomparagraph int
  for range text {
    randomparagraph = rand.Intn(len(text))
    println(len(text[randomparagraph]))
    if len(text[randomparagraph]) > 400 {
      index = randomparagraph
      break
    }
  }
  return index
}
