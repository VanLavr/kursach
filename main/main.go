package main

import (
	"fmt"
	"log"
	"net/http"
	"web/DBconnection"
	"web/api"
	"web/configs"
)

func main() {
	DBconnection.Connect()
	a := DBconnection.Test()
	fmt.Println(a)

	log.Println("Listening and serving on 127.0.0.1:8080")

	router := http.NewServeMux()
	fileserver := http.FileServer(http.Dir(configs.PathForFileServerIvan))
	router.Handle("/static/", http.StripPrefix("/static", fileserver))

	// [root] endoint
	router.HandleFunc("/", api.GetRoot)

	// [hello] endpoit
	router.HandleFunc("/hello", api.GetHello)

	// [get all and post] endpoint
	router.HandleFunc("/all", api.GetAll)

	// [get by id] endpoint
	router.HandleFunc("/all/", api.GetById)

	serverError := http.ListenAndServe(":8080", router)
	if serverError != nil {
		log.Fatal(serverError.Error())
	}
}
