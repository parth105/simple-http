package httpserver

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/parth105/simple-http/internal/wikipage"
)

func renderTemplate(w http.ResponseWriter, t string, p *wikipage.Page) {
	template, _ := template.ParseFiles(t)
	template.Execute(w, p)
}

func handler(w http.ResponseWriter, r *http.Request) {
	files, _ := ioutil.ReadDir(".")
	var pages []string

	for _, file := range files {
		if strings.Contains(file.Name(), ".page") {
			pages = append(pages, strings.Replace(file.Name(), ".page", "", -1))
		}
	}
	t := template.Must(template.ParseFiles("../web/welcome.html"))
	t.Execute(w, pages)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := wikipage.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "../web/view.html", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := wikipage.LoadPage(title)
	if err != nil {
		p = &wikipage.Page{Title: title}
	}
	renderTemplate(w, "../web/edit.html", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	content := r.FormValue("body")
	p := &wikipage.Page{Title: title, Body: []byte(content)}
	p.Save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func StartServer() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	//http.HandleFunc("/list/", viewHandler)
	log.Fatal(http.ListenAndServe(":8089", nil))
}
