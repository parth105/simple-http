package httpserver

import (
	"html/template"
	"log"
	"net/http"

	"github.com/parth105/simple-http/internal/wikipage"
)

func renderTemplate(w http.ResponseWriter, t string, p *wikipage.Page) {
	template, _ := template.ParseFiles(t)
	template.Execute(w, p)
}

/*
func handler(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Path[1:]
	returnString := ""
	if len(data) > 0 {
		returnString += fmt.Sprintf(" %s", data)
	}
	//fmt.Fprintf(w, "Welcome%s!", returnString)
	t, _ := template.ParseFiles("web/welcome.html")
	t.Execute(w, returnString)
	renderTemplate(w, "web/welcome.html", )
}
*/

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
	//http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	//http.HandleFunc("/list/", viewHandler)
	log.Fatal(http.ListenAndServe(":8089", nil))
}
