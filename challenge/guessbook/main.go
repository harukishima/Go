package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var guests []string

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// indexHandler serves the main page
func indexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if req.Method == http.MethodGet {
		temp := template.New("index")
		temp, _ = temp.Parse(indexHTML)
		_ = temp.Execute(w, guests)
	} else {
		io.WriteString(w, "Not supported method")
	}
}

// addHandler add a name to the names list
func addHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		newGuest := req.FormValue("name")
		if newGuest != "" {
			guests = append(guests, newGuest)
		}
		http.Redirect(w, req, "/", http.StatusSeeOther)
	} else {
		io.WriteString(w, "Not supported method")
	}
}

var indexHTML = `
<!DOCTYPE html>
<html>
    <head>
		<title>Guest Book ::Web GUI</title>
    </head>
    <body>
		<h1>Guest Book :: Web GUI</h1>
		<form action="/add" method="post">
		Name: <input name="name" /><submit value="Sign Guest Book">
		</form>
		<hr />
		<h4>Previous Guests</h4>
		<ul>
			{{range .}}
			<li>{{.}}</li>
			{{end}}
		</ul>
	</body>
</html>
`
