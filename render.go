package main

import (
	"log"
	"net/http"
)

func render(w http.ResponseWriter, r *http.Request) {
	log.Println("Welcom in render")
	//vars := mux.Vars(r)
	clusterName := r.FormValue("clusterName")
	log.Println(clusterName)
}
func getRender(w http.ResponseWriter, r *http.Request) {
	log.Println("Get render")
}
