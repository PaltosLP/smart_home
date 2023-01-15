package main
// https://freshman.tech/web-development-with-go/
// https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
import (
	"net/http"
	"os"
	"html/template"
	"fmt"
	"gopkg.in/yaml.v3"
)

type Config struct {
    Server struct {
        Port string `yaml:"port"`
        Host string `yaml:"host"`
    } `yaml:"server"`
    Weather_api struct {
        Url string `yaml:"url"`
        Auth string `yaml:"auth"`
    } `yaml:"weather_api"`
}


var tpl = template.Must(template.ParseFiles("tmpl/index.html"))


func indexHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("<h1>Hello World!</h1>"))
	tpl.Execute(w, nil)
}


func main() {
	//Load config----------------------
	f, err := os.Open("config.yml")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(cfg)
	//Load config----------------------


	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))


	http.ListenAndServe(":"+cfg.Server.Port, mux)
}
