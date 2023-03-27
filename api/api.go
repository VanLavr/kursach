package api

import (
	"net/http"
	"html/template"
	"log"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		log.Printf("|handled \"hello\"| |request=GET| |status: %v|", http.StatusNotFound)
		http.NotFound(w, r)
		return
	}

	log.Printf("|handled \"hello\"| |request=GET| |status: %v|", http.StatusOK)

	page, parseError := template.ParseFiles("D:\\desktop2\\GoProjects\\web\\static\\html\\hello.html")
	if parseError != nil {
		log.Fatal(parseError.Error())
	}
	page.Execute(w, nil)
}

func GetRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Printf("|handled \"home\"| |request=GET| |status: %v|", http.StatusNotFound)
		http.NotFound(w, r)
		return
	}

	log.Printf("|handled \"root\"| |request=GET| |status: %v|", http.StatusOK)

	page, parseError := template.ParseFiles("D:\\desktop2\\GoProjects\\web\\static\\html\\home.html")
	if parseError != nil {
		log.Fatal(parseError.Error())
	}
	page.Execute(w, nil)
}