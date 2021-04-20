package bookbytes

import (
  "strings"
)

type Book struct {
  fullhtml []string
  fulltext string
  title string
  author string
  booktext []string
  chaprefs []int
  currentchapref int
  chapter string
  paragraph int
}

func NewBook() Book {
  fullhtml := GetContents(GetFile())
  fulltext := StripLicense(fullhtml)
  booktext := SplitText(fullhtml)
  chaprefs := SetChapterReferences(booktext)
  return Book{
    fullhtml: fullhtml,
    fulltext: fulltext,
    title: SetTitle(fullhtml),
    author: SetAuthor(fullhtml),
    booktext: booktext,
    chaprefs: chaprefs,
    currentchapref: 0,
    chapter: "",
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
  return 
}

func SetChapterReferences(booktext []string) (chaprefs []int) {
  for i, line := range booktext {
    if strings.Contains(line, "name=\"chap") {
      chaprefs = append(chaprefs, i)
    }
  }
  chaprefs = append(chaprefs, len(booktext))
  return
}

func SetChapter() {
  var begin int
  var end int
  var beginindex int
  println("Current Paragraph Index: ", CurrentBook.paragraph)
  for i, ref := range CurrentBook.chaprefs {
    println(i, ":", ref)
    if ref > CurrentBook.paragraph {
      begin = CurrentBook.chaprefs[i-1]
      end = ref
      beginindex = i-1
      break;
    }
  }
  println(begin, ":", end)
  ret := strings.Join(CurrentBook.booktext[begin:end], " ")
  CurrentBook.currentchapref = beginindex
  CurrentBook.chapter = ret 
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
      begin = i
    } else if strings.Contains(line, "</p>") {
      end = i+1
      booktext = append(booktext, strings.Join(fullhtml[begin:end], " "))
    } else if strings.Contains(line, "name=\"chap") {
      booktext = append(booktext, line)
    }
  }
  return
}

/*** Getter Functions ***/

func BookPrinter(book Book) {
  println("Title: ", book.title)
  println("Author: ", book.author)
  println("Chapter References: ", book.chaprefs)
  //println(book.booktext[book.paragraph])
  //println(book.fulltext)
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
  CurrentBook.paragraph = CurrentBook.paragraph + 1
  return book.booktext[CurrentBook.paragraph]
}

func GetPreviousParagraph(book Book) string {
  CurrentBook.paragraph = CurrentBook.paragraph - 1
  return book.booktext[CurrentBook.paragraph]
}

func GetChapter() string {
  SetChapter()
  return CurrentBook.chapter
}
// need to check if at beginning or end of book
func GetNextChapter() string {
  begin := CurrentBook.chaprefs[CurrentBook.currentchapref+1]
  end := CurrentBook.chaprefs[CurrentBook.currentchapref+2] // BUG can't display last chapter because +2 is out of range
  ret := strings.Join(CurrentBook.booktext[begin:end], " ")
  CurrentBook.currentchapref = CurrentBook.currentchapref+1
  CurrentBook.chapter = ret 
  return CurrentBook.chapter
}

func GetPreviousChapter() string {
  begin := CurrentBook.chaprefs[CurrentBook.currentchapref-1]
  end := CurrentBook.chaprefs[CurrentBook.currentchapref]
  ret := strings.Join(CurrentBook.booktext[begin:end], " ")
  CurrentBook.currentchapref = CurrentBook.currentchapref-1
  CurrentBook.chapter = ret 
  return CurrentBook.chapter
}

func GetFirstChapter() string {
  begin := CurrentBook.chaprefs[0]
  end := CurrentBook.chaprefs[1]
  ret := strings.Join(CurrentBook.booktext[begin:end], " ")
  CurrentBook.currentchapref = 0
  CurrentBook.chapter = ret 
  return CurrentBook.chapter
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
