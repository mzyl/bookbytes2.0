package bookbytes

import (
  "fmt"
  "strings"
)

type Book struct {
  fullhtml []string
  fulltext string
  title string
  author string
  booktext []string
  paragraph int
}

func NewBook() Book {
  fullhtml := GetContents(GetFile())
  fulltext := StripLicense(fullhtml)
  booktext := SplitText(fullhtml)
  return Book{
    fullhtml: fullhtml,
    fulltext: fulltext,
    title: SetTitle(fullhtml),
    author: SetAuthor(fullhtml),
    booktext: booktext,
    paragraph: 0,
  }
}

var CurrentBook = NewBook()

/*** Setter Functions ***/

func GetNewBook() {
  CurrentBook = NewBook()
}

func SetTitle(booktext []string) string {
  return Get(booktext, "Title")
}

func SetAuthor(booktext []string) string {
  return Get(booktext, "Author")
}

func Get(booktext []string, attr string) (ret string) {
  for _, line := range booktext {
    if strings.Contains(line, attr) {
      text := strings.SplitAfter(line, attr+":")
      ret = strings.TrimSpace(strings.Join(text[1:], " "))
      break;
    }
  }
  return ret
}

func SetParagraph() int {
  CurrentBook.paragraph = NewParagraph(CurrentBook)
  return CurrentBook.paragraph 
}

func StripLicense(fullhtml []string) (bookstring string) {
  var booktext []string
  begin := 0
  mid := 0
  end := 0
  for i, line := range fullhtml {
    switch line {
    case "<pre>":
      begin = i
      if end < begin {
        booktext = append(booktext, strings.Join(fullhtml[mid:begin], " "))
      }
    case "</pre>":
      end = i
      if mid == 0 {
        mid = end + 1
      }
    }
  }
  booktext = append(booktext, strings.Join(fullhtml[end+1:], " "))
  bookstring = strings.Join(booktext, " ")
  return 
}

func SplitText(fullhtml []string) (booktext []string) {
  begin := 0
  end := 0
  for i, line := range fullhtml {
    if strings.Contains(line, "<p") { // be on lookout for weird behavior because of this change
      begin = i+1
    } else if strings.Contains(line, "</p>") {
      end = i
      booktext = append(booktext, strings.Join(fullhtml[begin:end], " "))
    } else if strings.Contains(line, "name=\"chap") {
      booktext = append(booktext, line)
    }
  }
  return
}

/*** Getter Functions ***/

func BookPrinter(book Book) {
  fmt.Println("Title: ", book.title)
  fmt.Println("Author: ", book.author)
  //fmt.Println(book.booktext[book.paragraph])
  //fmt.Println(book.fulltext)
}

func GetTitle(book Book) string {
  return book.title
}

func GetAuthor(book Book) string {
  return book.author
}

func GetInfo(book Book) string {
  return "This passage is from " + "<i>" + GetTitle(book) + "</i>" + " written by " + GetAuthor(book) + "."
}

func GetParagraph(book Book) string {
  return book.booktext[SetParagraph()]
}

func GetNextParagraph(book Book) string {
  CurrentBook.paragraph = NextParagraph(book)
  return book.booktext[CurrentBook.paragraph]
}

func GetPreviousParagraph(book Book) string {
  CurrentBook.paragraph = PreviousParagraph(book)
  return book.booktext[CurrentBook.paragraph]
}

// Need new function to call new paragraph and "print"

/*** Helper Functions ***/

func Between(line string, a string, b string) (ret string) {
  first := strings.Index(line, a)
  if first == -1 {
    return ""
  }

  last := strings.Index(line, b)
  if last == -1 {
    return ""
  }

  firstAdjusted := first + len(a)
  if firstAdjusted >= last {
    return ""
  }
  ret = line[firstAdjusted:last]
  return
}
