package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/* Json-ify the err response */
func respondErr(w http.ResponseWriter, status int, msg string) {
	if status > 499 {
		log.Println("Internal Server Error: ", msg)
	}

	respond(w, status, ErrResponse{
		Error: msg,
	})
}

/* Json-ify the response */
func respond(w http.ResponseWriter, status int, payload interface{}) {
	dat, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Error while responding: \n %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(status)
	w.Write(dat)
}
