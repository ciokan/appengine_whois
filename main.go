package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type GeoRes struct {
	Region  string `json:"region"`
	Country string `json:"country"`
	City    string `json:"city"`
	LatLong string `json:"lat_long"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	
	region := r.Header.Get("X-AppEngine-Region")
	country := r.Header.Get("X-AppEngine-Country")
	city := r.Header.Get("X-AppEngine-City")
	latLong := r.Header.Get("X-AppEngine-CityLatLong")
	
	res := GeoRes{
		Region:  region,
		Country: country,
		City:    city,
		LatLong: latLong,
	}
	
	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(js); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
		log.Printf("Defaulting to port %s", port)
	}
	
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
