package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Ave_grades, Happiness float64
}

func (u User) getAllinfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money: %d рублей", u.Name, u.Age, u.Money)
}
func (u *User) setNewname(newName string) {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	john := User{Name: "Sanya", Age: 18, Money: 5, Ave_grades: 4.5, Happiness: 100}
	//fmt.Fprintf(w, "Eto shto moi sait ashalet`")
	//john.setNewname("Sanya")
	//fmt.Fprintf(w, john.getAllinfo())

	templ, _ := template.ParseFiles("templates/home_page.html")
	templ.Execute(w, john)

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
