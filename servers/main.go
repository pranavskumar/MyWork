package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type OriginPool struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

func handle_originPoolCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := OriginPool{Name: "origin_pool_1", Status: "created"}
	json.NewEncoder(w).Encode(data)
}

func handle_originPoolGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := OriginPool{Name: "origin_pool_1", Status: "active"}
	json.NewEncoder(w).Encode(data)
}

func handle_originPoolList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := OriginPool{Name: "origin_pool_1", Status: "active"}
	json.NewEncoder(w).Encode(data)
}

func handle_originPoolUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := OriginPool{Name: "origin_pool_1", Status: "updated"}
	json.NewEncoder(w).Encode(data)
}

func handle_originPoolDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := OriginPool{Name: "origin_pool_1", Status: "deleted"}
	json.NewEncoder(w).Encode(data)
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/config/namespaces/create/origin_pools", handle_originPoolCreate).Methods("POST")
	r.HandleFunc("/api/config/namespaces/get/origin_pools/{name}", handle_originPoolGet).Methods("GET")
	r.HandleFunc("/api/config/namespaces/list/origin_pools", handle_originPoolList).Methods("GET")
	r.HandleFunc("/api/config/namespaces/replace/origin_pools/{name}", handle_originPoolUpdate).Methods("PUT")
	r.HandleFunc("/api/config/namespaces/delete/origin_pools/{name}", handle_originPoolDelete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
