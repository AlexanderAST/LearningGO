package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name                  string
	age                   uint16
	money                 int16
	ave_grades, happiness float64
}

func (u User) getAllinfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money: %d рублей", u.name, u.age, u.money)
}
func (u *User) setNewname(newName string) {
	u.name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	john := User{name: "John", age: 18, money: 5, ave_grades: 5.5, happiness: 100}
	//fmt.Fprintf(w, "Eto shto moi sait ashalet`")
	john.setNewname("Sanya")
	fmt.Fprintf(w, john.getAllinfo())
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
	//var John User =... есть более простой
	//john := User{name:"John", age: 44, money: 14325,ave_grades: 5.5, happiness: 100}

}
