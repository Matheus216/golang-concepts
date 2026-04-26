package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func handleView(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	fmt.Printf("parameter: %s", title)

	p, err := loadPage(title)

	if err != nil {
		log.Fatal(err)
	}

	t, _ := template.ParseFiles("view.html")

	t.Execute(w, p)
}

func handleEdit(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]

	p, _ := loadPage(title)

	t, _ := template.ParseFiles("edit.html")

	t.Execute(w, p)
}

func handleSave(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		title := r.URL.Path[len("/save/"):]

		t, _ := template.ParseFiles("edit.html")

		t.Execute(w, Page{Title: title})
	}
	if r.Method == http.MethodPost {
		handleSavePost(w, r)
	}

}

func handleSavePost(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")

	p := Page{Title: title, Body: []byte(body)}

	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

// this is a way of create a method to struct
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)

	// 0600 is a read-write permission
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/view/", handleView)
	http.HandleFunc("/edit/", handleEdit)
	http.HandleFunc("/save/", handleSave)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
