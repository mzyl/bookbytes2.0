package bookbytes

import (
  "os"
  "log"
  "time"
  "bufio"
  "math/rand"
)

func GetFile(booklist string) (filename string) {
  rand.Seed(time.Now().UnixNano())

  file, err := os.Open(booklist)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  
  // TODO: needs to randomly generate from file line count
  randomfile := rand.Int63n(61240)
  println("Random Number:", randomfile)
  _, err = file.Seek(randomfile, 0)
  files := bufio.NewReader(file)
  filename, err = files.ReadString('\n')
  filename, err = files.ReadString('\n')
  filename = filename[:len(filename)-1]

  //println(len(files))
  println("File: ../library/htmlmirror/" + filename[2:])
  return "../library/htmlmirror/" + filename[2:]
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
    // print number of characters in paragraph
    println(len(text[randomparagraph]))
    if len(text[randomparagraph]) > 400 {
      index = randomparagraph
      break
    }
  }
  return index
}
