package api

import (
	"encoding/json"
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

	w.WriteHeader(http.StatusOK)
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

	w.WriteHeader(http.StatusOK)
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

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(catalog)
		break

	case "POST":
		// HANDLING POST REQUEST FOR ALL DATA
		if r.URL.Path != "/all" {
			log.Printf("|handled \"all\"| |request=POST| |status: %v|", http.StatusNotFound)
			http.NotFound(w, r)
			return
		}

		decoder := json.NewDecoder(r.Body)

		var item DBconnection.Product

		errDecoding := decoder.Decode(&item)
		if errDecoding != nil {
			log.Printf("[api/GetAll] cannot decode POST body (%v)", errDecoding)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errDecoding.Error()))
		}

		createError := DBconnection.CreateItem(item)
		if createError != nil {
			log.Printf("[api/GetAll] cannot POST item (%v)", createError)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(createError.Error()))
		}

		log.Printf("|handled \"all\"| |request=POST| |status: %v|", http.StatusOK)
		w.WriteHeader(http.StatusOK)
		break

	case "DELETE":		
		// HANDLING DELETE REQUEST FOR DATA BY ID
		if r.URL.Path != "/all" {
			log.Printf("|handled \"all\"| |request=DELETE| |status: %v|", http.StatusNotFound)
			http.NotFound(w, r)
			return
		}

		log.Printf("|handled \"all\"| |request=DELETE| |status: %v|", http.StatusOK)

		decoder := json.NewDecoder(r.Body)

		var itemID int
		
		errorDecoding := decoder.Decode(&itemID)
		if errorDecoding != nil {
			log.Printf("[api/GetAll] cannot decode DELETE body (%v)", errorDecoding)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorDecoding.Error()))
		}

		delErr := DBconnection.DeleteItem(itemID)
		if delErr != nil {
			log.Printf("[api/GetAll] cannot DELETE item (%v)", delErr)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(delErr.Error()))
		}

		break
	}
}

func GetById(w http.ResponseWriter, r *http.Request) {
	AllIDs, err := DBconnection.GetAllIDs()
	if err != nil {
		log.Fatal(err.Error())
	}

	switch r.Method {

	case "GET":
		// HANDLING GET REQUEST FOR DATA BY ID
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

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
		break
	}
}
