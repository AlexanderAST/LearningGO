package main

import (
	"fmt"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Eto shto moi sait ashalet`")
}
func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Vtoraa stranica OOOOOOOOOOO")
}
func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":7777", nil)
}
func main() {
	handleRequest()
}
