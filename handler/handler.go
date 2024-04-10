package simplenotes

import (
    "math/rand/v2"
    "net/http"
    "strconv"
    "text/template"
)

type Note struct {
    Id string
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
    id := rand.Int()
    str_id := "#" + strconv.Itoa(id)
    tmpl := template.Must(template.ParseFiles("templates/note.html"))
    data := Note {
        Id: str_id,
        Title: r.PostFormValue("Title"),
        Content: r.PostFormValue("Content"),
    }
    tmpl.Execute(w, data)
}


