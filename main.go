package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) { // обращаемся через w и выводим страницу. ResponseWriter - тип данных.
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)                // URL, method
	log.Fatal(http.ListenAndServe(":8080", nil)) // обращение к методу - ListenAndServe. Создаем локальный сервер для отслеживания порта
}