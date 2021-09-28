package main

import (
	"fmt"

	// "log"
	"net/http"
)

type User struct {
}

// func main() {
// 	http.HandleFunc("/", handler)                // URL, method
// 	log.Fatal(http.ListenAndServe(":8080", nil)) // обращение к методу - ListenAndServe. Создаем локальный сервер для отслеживания порта
// }

func home(w http.ResponseWriter, r *http.Request) { // обращаемся через w и выводим страницу. ResponseWriter - тип данных.
	fmt.Fprintf(w, "Hi there!")
}
func contacts(w http.ResponseWriter, r *http.Request) { // обращаемся через w и выводим страницу. ResponseWriter - тип данных.
	fmt.Fprintf(w, "contacts!")
}

func handleRequest() { // обращаемся через w и выводим страницу. ResponseWriter - тип данных.
	http.HandleFunc("/", home)
	http.HandleFunc("/contacts/", contacts)
	http.ListenAndServe(":8080", nil)

}

func main() {
	handleRequest()
}
