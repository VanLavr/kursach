package api

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"web/DBconnection"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		log.Printf("|handled \"hello\"| |request=GET| |status: %v|", http.StatusNotFound)
		http.NotFound(w, r)
		return
	}

	log.Printf("|handled \"hello\"| |request=GET| |status: %v|", http.StatusOK)

	page, parseError := template.ParseFiles("D:\\Projects\\IVANLAVR\\kursach\\static\\html\\hello.html")
	if parseError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(parseError.Error()))
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

	page, parseError := template.ParseFiles("D:\\Projects\\IVANLAVR\\kursach\\static\\html\\home.html")
	if parseError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(parseError.Error()))
		log.Fatal(parseError.Error())
	}
	page.Execute(w, nil)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.URL.Path != "/all" {
		log.Printf("|handled \"all\"| |request=GET| |status: %v|", http.StatusNotFound)
		http.NotFound(w, r)
		return
	}

	log.Printf("|handled \"all\"| |request=GET| |status: %v|", http.StatusOK)

	catalog, catErr := DBconnection.GetAllProducts()
	if catErr != nil {
		log.Printf("[api/GetAll] cannot get all products (%v)", catErr)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(catErr.Error()))
	}
	json.NewEncoder(w).Encode(catalog)
}