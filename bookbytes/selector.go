package bookbytes

import (
	"bufio"
    "compress/bzip2"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"golang.org/x/net/html/charset"
)

func GetFile(booklist string) (filename string) {
	rand.Seed(time.Now().UnixNano())
	// number of bytes in booklist.txt
    bytes, err := os.Open("")
    switch booklist {
    case "compressedbooklist.txt" :   
        bytes, err = os.Open("compressedbytecount.txt")
        if err != nil {
            log.Fatal(err)
        }
    case "doclist.txt" :
        bytes, err = os.Open("docbytecount.txt")
        if err != nil {
            log.Fatal(err)
        }
    }
	defer bytes.Close()

	// TODO: Consider a better solution for getting a number out of a file.
	var line []string
	buffersize := bufio.NewScanner(bytes)
	for buffersize.Scan() {
		line = append(line, buffersize.Text())
	}

	num, _ := strconv.ParseInt(line[0], 10, 64)
	println("Number in file:", num)
	randombyte := rand.Int63n(num)
	println("Random Number:", randombyte)

	file, err := os.Open(booklist)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Seek(randombyte, 0)
	files := bufio.NewReader(file)
	filename, err = files.ReadString('\n')
	filename, err = files.ReadString('\n')
	filename = filename[:len(filename)-1]

	//println("File: ../library/htmlmirror/" + filename[2:])
	//return "../library/htmlmirror/" + filename[2:]
    switch booklist{
    case "compressedbooklist.txt" :
        println("File: ../library/compressed/" + filename[2:])
        return "../library/compressed/" + filename[2:]
    case "doclist.txt" :
        println("File: docs/" + filename[2:])
        return "docs/" + filename[2:]
    }

	//return "docs/11-h.htm"
    //return "../library/minified/120-h.htm"
    //return "docs/" + filename[2:]
	//return "../library/htmlmirror/4/8/8/2/48827/48827-h/48827-h.htm"
}

func GetContents(filename string) (text []string) {
    println("filename in contents: ", filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


    // gzip to come here
    zip := bzip2.NewReader(file)

	// TODO:
	// converts charmap based on what it finds in the file,
	// but defaults to windows1252 if it can't find a utf option.
	// may need to parse this more strictly if errors start coming up.
	// i.e. ISO-8859-1 instead of Windows1252
	reader, err := charset.NewReader(zip, "")

	scanner := bufio.NewScanner(reader)
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
		if len(text[randomparagraph]) > 400 {
			index = randomparagraph
			break
		}
	}
	return index
}
