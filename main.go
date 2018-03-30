package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Map struct {
	Width int32 `json:"width"`
	Height int32 `json:"height"`
	Walls []int32 `json:"walls"`
	Crates []int32 `json:"crates"`
	Targets []int32 `json:"targets"`
	Player int32 `json:"player"`
}

type MapEntry struct {
	ID int32 `json:"id"`
	Data Map `json:"data"`
}

var tempMaps []MapEntry

func GetMaps(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tempMaps)
}

func main() {
	tempMaps = append(tempMaps, MapEntry{1, Map{8,8, []int32{7, 8, 9}, []int32{19}, []int32{35}, 15}})
	tempMaps = append(tempMaps, MapEntry{2, Map{10,15, []int32{13, 15, 36}, []int32{51}, []int32{92}, 44}})
	tempMaps = append(tempMaps, MapEntry{3, Map{15,12, []int32{8, 9, 10, 11, 12, 64}, []int32{45, 48}, []int32{37, 38}, 25}})

	router := mux.NewRouter()
	router.HandleFunc("/map", GetMaps).Methods("GET")

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "OPTIONS"})

	log.Print("started")
	log.Fatal(http.ListenAndServe(":10500", handlers.CORS(allowedOrigins, allowedHeaders, allowedMethods)(router)))
}
