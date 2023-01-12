package main
// https://freshman.tech/web-development-with-go/

import (
	"net/http"
	"os"
	"log"
	"html/template"

	
	"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("web/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("<h1>Hello World!</h1>"))
	tpl.Execute(w, nil)
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
