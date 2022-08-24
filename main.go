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
	Hobbies               []string
}

func (u User) getAllinfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money: %d рублей", u.Name, u.Age, u.Money)
}
func (u *User) setNewname(newName string) {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	john := User{Name: "Sanya", Age: 18, Money: 5, Ave_grades: 4.5, Happiness: 100, Hobbies: []string{"Дота,Кс,Пабг,курить,пить"}}
	//fmt.Fprintf(w, "Eto shto moi sait ashalet`")
	//john.setNewname("Sanya")
	//fmt.Fprintf(w, john.getAllinfo())

	templ, _ := template.ParseFiles("templates/home_page.html")
	templ.Execute(w, john)

}
func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Vtoraa stranica OOOOOOOOOOO")
}
func about_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Саня 18 лет типо кодит")
}
func learn_more(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "что вам больше надо то, ну ладно я сижу")
}
func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.HandleFunc("/about/", about_page)
	http.HandleFunc("/learn_more/", learn_more)
	http.ListenAndServe(":7777", nil)
}
func main() {
	handleRequest()
	//var John User =... есть более простой
	//john := User{name:"John", age: 44, money: 14325,ave_grades: 5.5, happiness: 100}

}
