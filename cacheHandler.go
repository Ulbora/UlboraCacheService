package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleCacheSet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch r.Method {
		case "POST":
			kv := new(keyVal)
			decoder := json.NewDecoder(r.Body)
			error := decoder.Decode(&kv)
			if error != nil {
				log.Println(error.Error())
				http.Error(w, error.Error(), http.StatusBadRequest)
			}
			err := c.Set(kv.Key, kv.Value, 86400)
			var cres cacheRes
			if err != nil {
				cres.Success = false
			} else {
				cres.Success = true
				resJSON, err := json.Marshal(cres)
				if err != nil {
					log.Println(error.Error())
					http.Error(w, "json output failed", http.StatusInternalServerError)
				}
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, string(resJSON))
			}
		}
	}
}

func handleCacheGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["key"]
	switch r.Method {
	case "GET":
		var kvr keyValRes
		val, err := c.Get(key)
		if err != nil {
			kvr.Success = false
		} else {
			kvr.Success = true
			kvr.Value = val
		}
		resJSON, err := json.Marshal(kvr)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "json output failed", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(resJSON))
	}
}

func handleCacheDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["key"]
	switch r.Method {
	case "DELETE":
		res := c.Delete(key)

		var cr cacheRes
		cr.Success = res
		resJSON, err := json.Marshal(cr)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "json output failed", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(resJSON))
	}
}
