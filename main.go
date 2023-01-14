package main
// https://freshman.tech/web-development-with-go/
// https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
import (
	"net/http"
	"os"
	"html/template"

	
)

var tpl = template.Must(template.ParseFiles("tmpl/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("<h1>Hello World!</h1>"))
	tpl.Execute(w, nil)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	// fs := http.FileServer(http.Dir("ui/static/css"))
	// mux.Handle("/ui/css", http.StripPrefix("/ui/static/css", fs))

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))


	http.ListenAndServe(":"+port, mux)
}
