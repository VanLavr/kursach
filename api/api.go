package api

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
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
		log.Printf("|handled \"root\"| |request=GET| |status: %v|", http.StatusNotFound)
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

func GetById(w http.ResponseWriter, r *http.Request) {
	AllIDs, err := DBconnection.GetAllIDs()
	if err != nil {
		log.Fatal(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")

	log.Printf("|handled \"all/:id\"| |request=GET| |status: %v|", http.StatusOK)

	id := strings.TrimPrefix(r.URL.Path, "/all/")
	intedID, e := strconv.Atoi(id)
	if e != nil {
		log.Fatal(e.Error())
	}

	flag := true
	for _, v := range AllIDs {
		if v == intedID {
			flag = false
		}
	}
	if flag {
		log.Printf("|handled \"all/:id\"| |request=GET| |status: %v|", http.StatusNotFound)
		http.NotFound(w, r)
		return
	}

	product, prErr := DBconnection.GetProductById(intedID)
	if prErr != nil {
		log.Printf("[api/GetById] cannot get %d product (%v)", id, prErr)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(prErr.Error()))
	}
	json.NewEncoder(w).Encode(product)
}
