package main

// http.FileServer(http.Dir("D:\\desktop2\\GoProjects\\web\\frontend"))

import (
	"fmt"
	"log"
	"net/http"
	"web/DBconnection"
	"web/api"
)

func main() {
	DBconnection.Connect()
	a := DBconnection.Test()
	fmt.Println(a)

	log.Println("Listening and serving on 127.0.0.1:8080")

	router := http.NewServeMux()
	fileserver := http.FileServer(http.Dir("D:\\Projects\\IVANLAVR\\kursach\\static"))
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
