package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

var templates = template.Must(template.ParseFiles("edit.html", "view.html")) /* to avoid calling everytime the ParseFiles function. The ParseFiles
load and analyze the templates, then compiles them and it prepares them in a usable format. Then, the compiled template is always the same!!!
la funzione Must, se la funzione che wrappa restituisce un error, fa un panic stoppando l'esecuzione del programma. Annulla quindi il controllo degli
errori stoppando forzatamente l'esecuzione*/

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$") // we use this regular expression to validate the input path to prevent
// the user from providing any path he wants

type Page struct { // the fields of this structure, will be the place_holder in the HTNL template
	Title string
	Body  []byte
}

func (p *Page) save() error { // to save pages on the disk
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600) // 0600 means that the file should be created with read-write permissions for the current user only
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):] /* extrat the title from the url beacuse r.UR.Path is the path of the URL that the user has requested.
	If the user has requested for /save/my_page, r.UR.Path is a string containing that. Using [len("/save/"):] we take only the part [6:], that, for the
	previous example is the title my_page*/
	body := r.FormValue("body") // in edit.html we have inside the form, a tag textarea with the name "body". We are loading exactly it!!!
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound) // after the save, we redirect to the view page
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):] // extrat the title from the url
	p, err := loadPage(title)           //load the page from the disk
	if err != nil {                     // if the page doesn't exist it creates it
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("edit.html") // to parse (analyze) the template view.html to prepare this to be executed
	t.Execute(w, p)                          // executes the template writing the generating HTML to http.ResponseWriter
	// see the view.html to see the comment regardin the attributes of `p`
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p) // uses the templates that have been compiled at the start of the execution
	// and substitutes all the placeholders with the real data
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}

func main() {
	// function `HandleFunc` routes the incoming requests to the various handler functions
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil)) //`log.Fatal` logs possible errors returned by `ListentAndServe`
}
