package simplenotes

import (
    "net/http"
    "text/template"
)

type Note struct {
    Title string
    Content string
}

func Home(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    data := Note{
        Title: "World",
        Content: "",
    }
    tmpl.Execute(w, data)
}

func GetNote(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    tmpl := template.Must(template.ParseFiles("templates/note.html"))
    data := Note {
        Title: r.PostFormValue("Title"),
        Content: r.PostFormValue("Content"),
    }
    tmpl.Execute(w, data)
}


