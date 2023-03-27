package main

// http.FileServer(http.Dir("D:\\desktop2\\GoProjects\\web\\frontend"))

import (
	"net/http"
	"log"
	"web/api"
	"web/DBconnection"
	"fmt"
)

func main() {
	DBconnection.Connect()
	a := DBconnection.Test()
	fmt.Println(a)

	log.Println("Listening and serving on 127.0.0.1:8080")

	router := http.NewServeMux()
	fileserver := http.FileServer(http.Dir("D:\\desktop2\\GoProjects\\web\\static"))
	router.Handle("/static/", http.StripPrefix("/static", fileserver))

	// root endoint
	router.HandleFunc("/", api.GetRoot)
	// hello endpoit
	router.HandleFunc("/hello", api.GetHello)

	serverError := http.ListenAndServe(":8080", router)
	if serverError != nil {
		log.Fatal(serverError.Error())
	}
}