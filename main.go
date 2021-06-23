package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/go-chi/chi"
	"github.com/mzyl/bookbytes/bookbytes"
)

type Config struct {
	HttpAddr string `arg:"env:HTTP_ADDR"`
}

type Response struct {
	Headline  string `json:"headline"`
	Info      string `json:"info"`
	Filename  string `json:"filename"`
	Paragraph int    `json:"paragraph"`
	Chapter   int    `json:"chapter"`
}

func main() {
	rand.Seed(time.Now().Unix())

	c := Config{
		HttpAddr: ":8080",
	}

	arg.MustParse(&c)

	r := chi.NewRouter()
	r.Get("/", HandleFunc(file("./web/index.html")))
	r.Handle("/web/*", http.StripPrefix("/web",
		http.FileServer(http.Dir("./web"))))
	r.Post("/generate", HandleFunc(generate()))
	r.Post("/info", HandleFunc(info()))
	r.Post("/nextpg", HandleFunc(nextpg()))
	r.Post("/prevpg", HandleFunc(prevpg()))
	r.Post("/chapter", HandleFunc(chapter()))
	r.Post("/nextchapter", HandleFunc(nextchapter()))
	r.Post("/prevchapter", HandleFunc(prevchapter()))
	r.Post("/beginning", HandleFunc(beginning()))
	r.Post("/newbook", HandleFunc(newbook()))

	if err := http.ListenAndServe(c.HttpAddr, r); err != nil {
		log.Fatal(err)
	}
}

type HandlerFunc func(http.ResponseWriter, *http.Request) (int, error)

func HandleFunc(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statusCode, err := h(w, r)
		if err != nil {
			http.Error(w, err.Error(), statusCode)
			fmt.Println(err)
			return
		}
		if statusCode != http.StatusOK {
			w.WriteHeader(statusCode)
		}
	}
}

func file(filename string) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		http.ServeFile(w, r, filename)
		return http.StatusOK, nil
	}
}

func generate() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		body, err := ioutil.ReadAll(r.Body)
		filename := string(body)
		println("Filename:", filename)
		paragraph, index := bookbytes.GetNewParagraph(filename)
		var resp = Response{
			Headline:  paragraph,
			Paragraph: index,
		}
		b, err := json.Marshal(resp)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(b))

		return http.StatusOK, nil
	}
}

func info() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		body, err := ioutil.ReadAll(r.Body)
		filename := string(body)
		println("Filename:", filename)
		var resp = Response{
			Info: bookbytes.GetInfo(filename),
		}
		b, err := json.Marshal(resp)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(b))

		return http.StatusOK, nil
	}
}

func nextpg() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		body, err := ioutil.ReadAll(r.Body)
		data := strings.Split(string(body), ",")
		filename := data[0]
		index, _ := strconv.Atoi(data[1])
		println("Filename:", filename)
		var resp = Response{
			Headline:  bookbytes.GetNextParagraph(filename, index),
			Paragraph: index + 1,
		}
		b, err := json.Marshal(resp)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(b))

		return http.StatusOK, nil
	}
}

func prevpg() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		body, err := ioutil.ReadAll(r.Body)
		data := strings.Split(string(body), ",")
		filename := data[0]
		index, _ := strconv.Atoi(data[1])
		println("Filename:", filename)
		var resp = Response{
			Headline:  bookbytes.GetPreviousParagraph(filename, index),
			Paragraph: index - 1,
		}
		b, err := json.Marshal(resp)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(b))

		return http.StatusOK, nil
	}
}

func chapter() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		body, err := ioutil.ReadAll(r.Body)
		data := strings.Split(string(body), ",")
		filename := data[0]
		index, _ := strconv.Atoi(data[1])
		println("Filename:", filename)
		chapter, chapterref := bookbytes.GetChapter(filename, index)
		var resp = Response{
			Headline: chapter,
			Chapter:  chapterref,
		}
		b, err := json.Marshal(resp)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(b))

		return http.StatusOK, nil
	}
}

func nextchapter() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		body, err := ioutil.ReadAll(r.Body)
		data := strings.Split(string(body), ",")
		filename := data[0]
		index, _ := strconv.Atoi(data[1])
		println("Filename:", filename)
		chapter, chapterref := bookbytes.GetNextChapter(filename, index)
		var resp = Response{
			Headline: chapter,
			Chapter:  chapterref,
		}
		b, err := json.Marshal(resp)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(b))

		return http.StatusOK, nil
	}
}

func prevchapter() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		body, err := ioutil.ReadAll(r.Body)
		data := strings.Split(string(body), ",")
		filename := data[0]
		index, _ := strconv.Atoi(data[1])
		println("Filename:", filename)
		chapter, chapterref := bookbytes.GetPreviousChapter(filename, index)
		var resp = Response{
			Headline: chapter,
			Chapter:  chapterref,
		}
		b, err := json.Marshal(resp)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(b))

		return http.StatusOK, nil
	}
}

func beginning() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		body, err := ioutil.ReadAll(r.Body)
		filename := string(body)
		println("Filename:", filename)
		chapter, chapterref := bookbytes.GetFirstChapter(filename)
		var resp = Response{
			Headline: chapter,
			Chapter:  chapterref,
		}
		b, err := json.Marshal(resp)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(b))

		return http.StatusOK, nil
	}
}

func newbook() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		paragraph, filename, index := bookbytes.Init()
		var resp = Response{
			Headline:  paragraph,
			Filename:  filename,
			Paragraph: index,
		}
		b, err := json.Marshal(resp)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(b))

		return http.StatusOK, nil
	}
}
