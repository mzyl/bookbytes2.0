package bookbytes

import (
  "regexp"
  "strings"
)

type Book struct {
  filename string
  fullhtml []string
  fulltext string
  title string
  author string
  language string
  booktext []string
  chaprefs []int
  currentchapref int
  chapter string
  paragraph int
}

func NewBook() Book {
  filename := GetFile("booklist.txt")
  fullhtml := GetContents(filename)
  fulltext := StripLicense(fullhtml)
  booktext := SplitText(fullhtml)
  chaprefs := SetChapterReferences(booktext)
  return Book{
    filename: filename,
    fullhtml: fullhtml,
    fulltext: fulltext,
    title: SetTitle(fullhtml),
    author: SetAuthor(fullhtml),
    language: SetLanguage(fullhtml),
    booktext: booktext,
    chaprefs: chaprefs,
    currentchapref: 0,
    chapter: "",
    paragraph: 0,
  }
}

func NewBookFromFilename(filename string, paragraph int) Book {
  fullhtml := GetContents(filename)
  fulltext := StripLicense(fullhtml)
  booktext := SplitText(fullhtml)
  chaprefs := SetChapterReferences(booktext)
  return Book{
    filename: filename,
    fullhtml: fullhtml,
    fulltext: fulltext,
    title: SetTitle(fullhtml),
    author: SetAuthor(fullhtml),
    language: SetLanguage(fullhtml),
    booktext: booktext,
    chaprefs: chaprefs,
    currentchapref: 0,
    chapter: "",
    paragraph: paragraph,
  }
}


/*** Setter Functions ***/


func SetTitle(booktext []string) string {
  return Get(booktext, "Title")
}

func SetAuthor(booktext []string) string {
  return Get(booktext, "Author")
}

func SetLanguage(booktext []string) string {
  return Get(booktext, "Language")
}

func Get(booktext []string, attr string) (ret string) {
  for _, line := range booktext {
    if strings.Contains(line, attr+":") {
      text := strings.SplitAfter(line, attr+":")
      ret = strings.TrimSpace(strings.Join(text[1:], " "))
      break;
    }
  }
  return 
}

func SetChapterReferences(booktext []string) (chaprefs []int) {
  match, _ := regexp.Compile("<h[1-6]")
  for i, line := range booktext {
    //if strings.Contains(line, "name=\"chap") {
    if match.MatchString(strings.ToLower(line)) {
      chaprefs = append(chaprefs, i)
    }
  }
  chaprefs = append(chaprefs, len(booktext))
  println("Chapters found: ", len(chaprefs)-1)
  return
}

func SetChapter(filename string, paragraph int) (string, int) {
  book := NewBookFromFilename(filename, paragraph)
  var begin int
  var end int
  var beginindex int
  for i, ref := range book.chaprefs {
    if ref > book.paragraph {
      begin = book.chaprefs[i-1]
      end = ref
      beginindex = i-1
      break;
    }
  }
 ret := strings.Join(book.booktext[begin:end], " ")
  book.currentchapref = beginindex
  return ret, book.currentchapref
}


func StripLicense(fullhtml []string) (bookstring string) {
  // May need to have title and author come out of <pre> in the future
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
  match, _ := regexp.Compile("<h[1-6]")
  for i, line := range fullhtml {
    if strings.Contains(strings.ToLower(line), "<p") {
      begin = i
    } else if strings.Contains(strings.ToLower(line), "</p>") {
      end = i+1
      booktext = append(booktext, strings.Join(fullhtml[begin:end], " "))
    //} else if strings.Contains(line, "=\"chap") {
    //} else if strings.Contains(line, "<h2>") {
    } else if match.MatchString(strings.ToLower(line)) {
      booktext = append(booktext, line)
      booktext = append(booktext, fullhtml[i+1])
    }
  }
  booktext = append(booktext, "<h5><i>Fin.</i></h5>")
  return
}

/*** Getter Functions ***/

func BookPrinter(book Book) {
  println("Title: ", book.title)
  println("Author: ", book.author)
  println("Language: ", book.language)
  //println(book.booktext[book.paragraph])
  //println(book.fulltext)
}

func GetFilename(book Book) string{
  return book.filename
}

func Init() (paragraph string, filename string, index int) {
  book := NewBook()
  filename = book.filename
  index = NewParagraph(book)
  paragraph = book.booktext[index]
  return
}

func GetNewParagraph(filename string) (paragraph string, index int) {
  book := NewBookFromFilename(filename, 0)
  index = NewParagraph(book)
  paragraph = book.booktext[index]
  return
}

func GetTitle(book Book) string {
  return book.title
}

func GetAuthor(book Book) string {
  return book.author
}

func GetLanguage(book Book) string {
  return book.language
}

func GetInfo(filename string) string {
  book := NewBookFromFilename(filename, 0)
  return "This passage is from " + "<i>" + GetTitle(book) + "</i>" + 
    " written by " + GetAuthor(book) + " in " + GetLanguage(book) + "."
}

func GetParagraphIndex(book Book) int {
  return book.paragraph
}

func GetNextParagraph(filename string, paragraph int) string {
  book := NewBookFromFilename(filename, paragraph)
  book.paragraph = paragraph + 1
  return book.booktext[book.paragraph]
}

func GetPreviousParagraph(filename string, paragraph int) string {
  book := NewBookFromFilename(filename, paragraph)
  book.paragraph = paragraph - 1
  return book.booktext[book.paragraph]
}

func GetChapter(filename string, paragraph int) (chapter string, chapterref int) {
  chapter, chapterref = SetChapter(filename, paragraph)
  return 
}

func GetNextChapter(filename string, index int) (chapter string, chapterref int) {
  book := NewBookFromFilename(filename, 0)
  begin := book.chaprefs[index + 1]
  end := book.chaprefs[index + 2]
  chapter = strings.Join(book.booktext[begin:end], " ")
  chapterref = index + 1
  return
}

func GetPreviousChapter(filename string, index int) (chapter string, chapterref int) {
  book := NewBookFromFilename(filename, 0)
  begin := book.chaprefs[index - 1]
  end := book.chaprefs[index]
  chapter = strings.Join(book.booktext[begin:end], " ")
  chapterref = index - 1
  return
}

func GetFirstChapter(filename string) (chapter string, chapterref int) {
  book := NewBookFromFilename(filename, 0)
  begin := book.chaprefs[0]
  end := book.chaprefs[1]
  chapter = strings.Join(book.booktext[begin:end], " ")
  chapterref = 0
  return
}

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
