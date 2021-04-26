package bookbytes

import (
  "os"
  "log"
  "time"
  "bufio"
  "math/rand"
  //"path/filepath"
)

//var Files []string

func GetFile(booklist string) (filename string) {
  //var randomfile int
  rand.Seed(time.Now().UnixNano())

  file, err := os.Open(booklist)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  
  // TODO: needs to randomly generate from file line count
  randomfile := rand.Int63n(61240)
  _, err = file.Seek(randomfile, 0)
  files := bufio.NewReader(file)
  filename, err = files.ReadString('\n')
  filename, err = files.ReadString('\n')
  filename = filename[:len(filename)-1]
  println("File: ")
  println(filename)

/**  
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
  for range files {
    randomfile = rand.Intn(len(files))
    println(randomfile)
    if randomfile != 0 {
      filename = files[randomfile]
      break
    }
  }
  **/

  //println(len(files))
  return "../library/htmlmirror/" + filename
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
