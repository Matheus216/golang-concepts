package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/view/", handleView)
	http.HandleFunc("/edit/", handleEdit)
	http.HandleFunc("/save/", handleSave)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func handleView(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	fmt.Printf("parameter: %s", title)

	p, _ := loadPage(title)

	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func handleEdit(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]

	p, _ := loadPage(title)

	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>",
		p.Title, p.Title, p.Body)
}

func handleSave(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		title := r.URL.Path[len("/save/"):]

		fmt.Fprintf(w, "<h1>Creating</h1>"+
			"<form action=\"/save/%s\" method=\"POST\">"+
			"<textarea name=\"body\"></textarea><br>"+
			"<input type=\"submit\" value=\"Save\">"+
			"</form>", title)
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
