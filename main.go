package main
import (
	"io/ioutil"
	"net/http"
	"html/template"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) Save() {
	filename := p.Title + "txt"
	ioutil.WriteFile(filename,p.Body, 0600)
}

func loadPage(title string) *Page {
	filename := title + "txt"
	body, _ := ioutil.ReadFile(filename)
	return &Page{title, body}
}

var templates = template.Must(template.ParseFiles("networking/server/edit.html", "networking/server/create.html", "networking/server/view.html"))
func create(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create.html", nil)
}

func save(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("body")
	title := r.FormValue("title")
	p := &Page{title, []byte(body)}
	p.Save()
	http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

func view(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	page := loadPage(title)
	p := &Page{title, page.Body}
	templates.ExecuteTemplate(w, "view.html", p)
}

func edit(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	page := loadPage(title)
	p := &Page{title, page.Body}
	templates.ExecuteTemplate(w, "edit.html", p)
}

func main() {
	http.HandleFunc("/create", create)
	http.HandleFunc("/save", save)
	http.HandleFunc("/view/", view)
	http.HandleFunc("/edit/", edit)
	err := http.ListenAndServe(":3050", nil)
	if err != nil {
		panic(err)
	}
}