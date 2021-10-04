package main

import (
	"fmt"
	"html/template" //загружает разметку с new записями
	"log"
	"net/http"
)

var html *template.Template
var err error

type Data struct {
	input     string
	Output    string
	Font      string
	ErrorNum  int
	ErrorText string
}

func init() {
	//download a form HTML of template
	html, err = template.ParseFiles("templates/index.html") //templates/index.html - used for create new value of Template
	if err != nil {
		log.Fatal("ssc", err)
	}
}

// ResponseWriter передается для записи вывода
// Execute(метод) - чтобы получить вывод от значения Template из переменной html
func home(wr http.ResponseWriter, req *http.Request) {
	html.Execute(wr, nil)
}

func ascii(wr http.ResponseWriter, req *http.Request) {
	D := Data{}
	text := req.FormValue("input") // get value
	// font := req.FormValue("font")
	// out, _ := dirs.GetArt(text, font)
	fmt.Println("HERE")
	// fmt.Println(font)
	D.Output = text
	// if req.FormValue("process") == "show" {
	html.Execute(wr, D)
	// }
}

func main() {
	//http := http.NewServeMux()
	http.HandleFunc("/", home)
	http.HandleFunc("/ascii-art", ascii)
	fmt.Println("8080")
	er := http.ListenAndServe(":8080", nil) //
	if er != nil {
		fmt.Println(er)
	}
}
