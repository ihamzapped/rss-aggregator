package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
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

/* Parse json request into the given type */
func parseBody[T interface{}](w http.ResponseWriter, r *http.Request, body T) (T, error) {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&body)

	if err != nil {
		respondErr(w, http.StatusBadRequest, fmt.Sprintf("Error parsing request: %v", err))
		return body, err
	}

	return body, nil

}

/* parse xml file to rssfeed type */
func urlToFeed(url string) (RssFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)

	if err != nil {
		return RssFeed{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)

	if err != nil {
		return RssFeed{}, err
	}

	feed := RssFeed{}

	err = xml.Unmarshal(dat, &feed)

	if err != nil {
		return RssFeed{}, err
	}

	return feed, nil

}
