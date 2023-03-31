package api

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"web/DBconnection"
	"web/configs"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		log.Printf("|handled \"hello\"| |request=GET| |status: %v|", http.StatusNotFound)
		http.NotFound(w, r)
		return
	}

	log.Printf("|handled \"hello\"| |request=GET| |status: %v|", http.StatusOK)

	page, parseError := template.ParseFiles(configs.PathForLogIvan)
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

	page, parseError := template.ParseFiles(configs.PathForHomeIvan)
	if parseError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(parseError.Error()))
		log.Fatal(parseError.Error())
	}
	page.Execute(w, nil)
}

func GetAll(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		// HANDLING GET REQUEST FOR ALL DATA
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

	case "POST":
		// HANDLING POST REQUEST FOR ALL DATA
		decoder := json.NewDecoder(r.Body)

		var item DBconnection.Product

		errDecoding := decoder.Decode(&item)
		if errDecoding != nil {
			log.Printf("[api/GetAll] cannot decode POST body (%v)", errDecoding)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errDecoding.Error()))
		}

		fmt.Println(item)
	}
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
