package main

import "net/http"

func healthCheck(w http.ResponseWriter, r *http.Request) {
	respond(w, http.StatusOK, "All good!")
}

func errCheck(w http.ResponseWriter, r *http.Request) {
	respondErr(w, http.StatusBadRequest, "All Bad!")
}
