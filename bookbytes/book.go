package bookbytes

import (
    "fmt" // Remove later
    "io"
    "log"
    //"errors"
    "bytes"

	"regexp"
	"strings"

    "golang.org/x/net/html"
)

type Book struct {
	filename       string
	fullhtml       []string
	fulltext       []string
	title          string
	author         string
	language       string
	booktext       []string
	chaprefs       []int
	currentchapref int
	chapter        string
	paragraph      int
}

// TODO: Find solution to eliminate one of these functions.
func NewBook() Book {
	filename := GetFile("booklist.txt")
	fullhtml := GetContents(filename)
	fulltext := StripLicense(fullhtml)
	booktext := SplitTextNode(fulltext)
	chaprefs := SetChapterReferences(booktext)
	return Book{
		filename:       filename,
		fullhtml:       fullhtml,
		fulltext:       fulltext,
		title:          SetTitle(fullhtml),
		author:         SetAuthor(fullhtml),
		language:       SetLanguage(fullhtml),
		booktext:       booktext,
		chaprefs:       chaprefs,
		currentchapref: 0,
		chapter:        "",
		paragraph:      0,
	}
}

func NewBookFromFilename(filename string, paragraph int) Book {
	fullhtml := GetContents(filename)
	fulltext := StripLicense(fullhtml)
	booktext := SplitTextNode(fulltext)
	chaprefs := SetChapterReferences(booktext)
	return Book{
		filename:       filename,
		fullhtml:       fullhtml,
		fulltext:       fulltext,
		title:          SetTitle(fullhtml),
		author:         SetAuthor(fullhtml),
		language:       SetLanguage(fullhtml),
		booktext:       booktext,
		chaprefs:       chaprefs,
		currentchapref: 0,
		chapter:        "",
		paragraph:      paragraph,
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
			break
		}
	}
	return
}

// Maybe replace the regexp here?
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
			beginindex = i - 1
			break
		}
	}
	ret := strings.Join(book.booktext[begin:end], " ")
	book.currentchapref = beginindex
	return ret, book.currentchapref
}

// I think regexp is fine to use here.
func StripLicense(fullhtml []string) []string {
	var booktext []string
	begin := -1
	end := -1
    match, _ := regexp.Compile(`\*{3} *((?:START|END)[\w\W]+)\*{3}`)

	for i, line := range fullhtml {
        if match.MatchString(line) {
            if begin == -1 {
                println("Found Start")
                begin = i + 1
            } else {
                println("Found End")
                end = i - 1
                break
            }
        }
        end = i
	}

    if begin == -1 {
        begin = 0
    }

    for _, line := range fullhtml[begin:end] {
        booktext = append(booktext, line)
    }

	return booktext
}

// This is no longer in use. Using html.Node instead.
func SplitText(fullhtml []string) (booktext []string) {
	begin := 0
	end := 0
    // try regexp for different paragraph tags as well? p, span, etc.
    paragraphbegin, _ := regexp.Compile(`<((?:p|div)[\w\W]+)`)
    paragraphend, _ := regexp.Compile(`</((?:p|div))>`)
	chapter, _ := regexp.Compile(`<h[1-6]`)

	for i, line := range fullhtml {
		if paragraphbegin.MatchString(strings.ToLower(line)) {
			begin = i
		} 
        if paragraphend.MatchString(strings.ToLower(line)) {
			end = i + 1
			booktext = append(booktext, strings.
				Join(fullhtml[begin:end], " "))
		} 
        if chapter.MatchString(strings.ToLower(line)) {
			booktext = append(booktext, line)
			booktext = append(booktext, fullhtml[i+1])
		}
	}
	booktext = append(booktext, "<h5><i>Fin.</i></h5>")
	return
}

func SplitTextNode(fullhtml []string) (booktext []string) {                                 
    text := strings.Join(fullhtml, " ")
    doc, _ := html.Parse(strings.NewReader(text))

    var body *html.Node
    var crawler func(*html.Node)

    textTags := []string {
        "h1", "h2", "h3", "h4", "h5", "h6",
        "p",
    }

    tag := ""
    enter := false

    crawler = func(node *html.Node) {
        tag = node.Data
        for _, t := range textTags {
            if tag == t {
                enter = true
                break
            }
        }
        if node.Type == html.ElementNode && enter {
           body = node
           booktext = append(booktext, renderNode(body))
           //fmt.Println(renderNode(body))
           enter = false
           return
        }
        for child := node.FirstChild; child != nil; child = child.NextSibling {
          crawler(child)
        }
    }
    crawler(doc)
    return
}

func renderNode(n *html.Node) string {
    var buf bytes.Buffer
    w := io.Writer(&buf)
    html.Render(w, n)
    return buf.String()
}

// No longer using Token because it would split lines with new tags, e.g. <i>
func SplitTextToken(fullhtml []string) {
    text := strings.Join(fullhtml, " ")
    reader := strings.NewReader(text)
    tokenizer := html.NewTokenizer(reader)

    textTags := []string {
        "p", "h1", "h2", "h3", "h4", "h5", "h6",
    }

    tag := ""
    enter := false
    
    for {
        tt := tokenizer.Next()
        t := tokenizer.Token()

        err := tokenizer.Err()
        if err == io.EOF {
            break
        }

        switch tt {
        case html.ErrorToken:
            log.Fatal(err)
        case html.StartTagToken, html.SelfClosingTagToken:
            if enter {
                break
            }
            tag = t.Data

            for _, ttt := range textTags {
                if tag == ttt {
                    enter = true
                    break
                }
            }
        case html.EndTagToken:
            // if EndTagToken != tag: break
            if t.Data != tag {
                break
            }
            enter = false
        case html.TextToken:
            if enter {
                data := strings.TrimSpace(t.Data)
                if len(data) > 0 {
                    fmt.Println(data)
                }
            }
        }
    }
}

/*** Getter Functions ***/

func BookPrinter(book Book) {
	println("Title: ", book.title)
	println("Author: ", book.author)
	println("Language: ", book.language)
	//println(book.booktext[book.paragraph])
	//fmt.Println(book.fulltext)
    //SplitTextToken(book.fulltext)
    SplitTextNode(book.fulltext)
}

func GetFilename(book Book) string {
	return book.filename
}

func Init() (string, string, int) {
	book := NewBook()
	filename := book.filename
	index := NewParagraph(book)
	paragraph := book.booktext[index]
	return paragraph, filename, index
}

func GetNewParagraph(filename string) (string, int) {
	book := NewBookFromFilename(filename, 0)
	index := NewParagraph(book)
	paragraph := book.booktext[index]
	return paragraph, index
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

func GetChapter(filename string, paragraph int) (string, int) {
	chapter, chapterref := SetChapter(filename, paragraph)
	return chapter, chapterref
}

func GetNextChapter(filename string, index int) (string, int) {
	book := NewBookFromFilename(filename, 0)
	begin := book.chaprefs[index+1]
	end := book.chaprefs[index+2]
	chapter := strings.Join(book.booktext[begin:end], " ")
	chapterref := index + 1
	return chapter, chapterref
}

func GetPreviousChapter(filename string, index int) (string, int) {
	book := NewBookFromFilename(filename, 0)
	begin := book.chaprefs[index-1]
	end := book.chaprefs[index]
	chapter := strings.Join(book.booktext[begin:end], " ")
	chapterref := index - 1
	return chapter, chapterref
}

func GetFirstChapter(filename string) (string, int) {
	book := NewBookFromFilename(filename, 0)
	begin := book.chaprefs[0]
	end := book.chaprefs[1]
	chapter := strings.Join(book.booktext[begin:end], " ")
	chapterref := 0
	return chapter, chapterref
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
