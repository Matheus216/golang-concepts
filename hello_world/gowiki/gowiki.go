package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type Page struct {
	Title string
	Body  []byte
}

type IndexData struct {
	Items []string
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	entries, err := os.ReadDir("./pages")

	t, _ := template.ParseFiles("./tmlp/index.html")

	if err != nil || len(entries) == 0 {
		t.Execute(w, []string{})
		return
	}

	pages := []string{}

	for _, entry := range entries {
		if !entry.IsDir() {
			pages = append(pages, strings.Replace(entry.Name(), ".txt", "", 1))
		}
	}

	t.Execute(w, IndexData{Items: pages})
}

func handleView(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	fmt.Printf("parameter: %s", title)

	p, err := loadPage(title)

	if err != nil {
		http.Redirect(w, r, "/save/"+title, http.StatusFound)
		return
	}

	t, _ := template.ParseFiles("./tmlp/view.html")

	t.Execute(w, p)
}

func handleEdit(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]

	p, errLoad := loadPage(title)

	if errLoad != nil {
		http.Redirect(w, r, "/save/"+title, http.StatusFound)
		return
	}

	t, _ := template.ParseFiles("./tmlp/edit.html")

	t.Execute(w, p)
}

func handleSave(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		title := r.URL.Path[len("/save/"):]

		t, _ := template.ParseFiles("./tmlp/edit.html")

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

// this is a way of create a method to struct
func (p *Page) save() error {
	filename := "./pages/" + p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)

	// 0600 is a read-write permission
}

func loadPage(title string) (*Page, error) {
	filename := "./pages/" + title + ".txt"
	body, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/view/", handleView)
	http.HandleFunc("/edit/", handleEdit)
	http.HandleFunc("/save/", handleSave)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
