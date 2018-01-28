package main

import (
	"fmt"
	"net/http"

	cacher "UlboraCacheService/cacher"

	"github.com/gorilla/mux"
)

var c cacher.Cacher

type keyVal struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type keyValRes struct {
	Success bool   `json:"success"`
	Value   string `json:"value"`
}
type cacheRes struct {
	Success bool `json:"success"`
}

func main() {
	c.NewCache(512 * 1024)
	fmt.Println("Cache Server running on port 3010!")
	router := mux.NewRouter()
	router.HandleFunc("/rs/cache/set", handleCacheSet).Methods("POST")
	router.HandleFunc("/rs/cache/get/{key}", handleCacheGet).Methods("GET")
	router.HandleFunc("/rs/cache/delete/{key}", handleCacheDelete).Methods("DELETE")
	http.ListenAndServe(":3010", router)
}
