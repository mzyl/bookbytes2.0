// credit to Christian Bargmann (cbrgm)
package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/go-chi/chi"
  "github.com/mzyl/bookbytes/bookbytes"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Config struct {
	HttpAddr string `arg:"env:HTTP_ADDR"`
}

type Response struct {
	Headline string `json:"headline"`
}


func main() {
	rand.Seed(time.Now().Unix())

	c := Config{
		HttpAddr: ":8080",
	}

	arg.MustParse(&c)

  //cbg := bookbytes.CurrentBook

	r := chi.NewRouter()
	r.Get("/", HandleFunc(file("./web/index.html")))
	r.Handle("/web/*", http.StripPrefix("/web", http.FileServer(http.Dir("./web"))))
	r.Post("/generate", HandleFunc(generate(bookbytes.CurrentBook)))
  r.Post("/info", HandleFunc(info(bookbytes.CurrentBook)))
	r.Post("/nextpg", HandleFunc(nextpg(bookbytes.CurrentBook)))
	r.Post("/prevpg", HandleFunc(prevpg(bookbytes.CurrentBook)))
  r.Post("/newbook", HandleFunc(newbook(bookbytes.CurrentBook)))

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

func generate(bookbytes.Book) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		var resp = Response{
			Headline: bookbytes.GetParagraph(bookbytes.CurrentBook), // I don't think I should have to use a func?
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

func info(bookbytes.Book) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		var resp = Response{
			Headline: bookbytes.GetInfo(bookbytes.CurrentBook), // I don't think I should have to use a func?
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

func nextpg(bookbytes.Book) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		var resp = Response{
			Headline: bookbytes.GetNextParagraph(bookbytes.CurrentBook), // I don't think I should have to use a func?
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

func prevpg(bookbytes.Book) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		var resp = Response{
			Headline: bookbytes.GetPreviousParagraph(bookbytes.CurrentBook), // I don't think I should have to use a func?
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

func newbook(bookbytes.Book) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
    bookbytes.GetNewBook()
    fmt.Println(bookbytes.GetTitle(bookbytes.CurrentBook))
		var resp = Response{
			Headline: bookbytes.GetParagraph(bookbytes.CurrentBook), // I don't think I should have to use a func?
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
