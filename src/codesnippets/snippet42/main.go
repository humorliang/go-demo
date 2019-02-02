package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

/*
go get github.com/gorilla/mux

Serve Static Assets using the Mux Router
创建静态文件用mux router
*/

const (
	STATIC_DIR = "/static/"
	PORT       = "8080"
)

func main() {
	router := NewRouter()
	err := http.ListenAndServe(":"+PORT, router)
	if err != nil {
		log.Fatal(err)
	}
}

func NewRouter() *mux.Router {
	//
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix(STATIC_DIR).
		Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("."+STATIC_DIR))))
	return router
}
