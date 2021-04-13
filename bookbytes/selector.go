package bookbytes

import (
  "os"
  "fmt"
  "log"
  "time"
  "bufio"
  "math/rand"
  "path/filepath"
)

func GetFile() (filename string) {
  var files []string
  var randomfile int
  root := "./books"
  rand.Seed(time.Now().UnixNano())
  err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
    files = append(files, path)
    return nil
  })
  if err != nil {
    panic(err)
  }
  for range files {
    randomfile = rand.Intn(len(files))
    fmt.Println(randomfile)
    if randomfile != 0 {
      filename = files[randomfile]
      break
    }
  }
  fmt.Println()
  fmt.Println("./" + filename)
  return "./" + filename
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
    fmt.Println(len(text[randomparagraph]))
    if len(text[randomparagraph]) > 400 {
      index = randomparagraph
      break
    }
  }
  return index
}

func NextParagraph(book Book) int {
  return book.paragraph + 1
}

func PreviousParagraph(book Book) int {
  return book.paragraph - 1
}
