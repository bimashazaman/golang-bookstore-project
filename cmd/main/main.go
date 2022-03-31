package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/bimashazaman/golang-bookstore-project/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}