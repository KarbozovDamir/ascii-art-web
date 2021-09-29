package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	// "github.com/test/dirs"
)

var Tmp *template.Template
var err error

type Data struct {
	input     string
	Output    string
	Font      string
	ErrorNum  int
	ErrorText string
}

func init() {
	Tmp, err = template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("ssc", err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	Tmp.Execute(w, nil)
}

func ascii(w http.ResponseWriter, r *http.Request) {
	D := Data{}
	text := r.FormValue("input")
	// font := r.FormValue("font")
	// out, _ := dirs.GetArt(text, font)
	fmt.Println("HERE")
	// fmt.Println(font)
	D.Output = text
	// if r.FormValue("process") == "show" {
	Tmp.Execute(w, D)
	// }
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/ascii-art", ascii)
	fmt.Println("8080")
	er := http.ListenAndServe(":8080", mux)
	if er != nil {
		fmt.Println(er)
	}
}
