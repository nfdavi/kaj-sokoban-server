package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var db *sql.DB

func GetMaps(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(GetAllMaps())
	log.Print("Served GetMaps")
}

func GetScores(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	mapId, err := strconv.Atoi(params["mapId"])
	if err != nil {
		log.Fatal(err)
	}

	var limit, offset int

	if params["limit"] != "" {
		limit, err = strconv.Atoi(params["limit"])
		if err != nil {
			log.Fatal(err)
		}
	}

	if params["offset"] != "" {
		offset, err = strconv.Atoi(params["offset"])
		if err != nil {
			log.Fatal(err)
		}
	}

	var scores []ScoreEntry

	if limit == 0 {
		scores = GetAllScoresForMap(mapId)
	} else {
		scores = GetScoresForMap(mapId, limit, offset)
	}

	json.NewEncoder(w).Encode(scores)
	log.Print("Served GetScores")
}

func PostScore(w http.ResponseWriter, r *http.Request) {
	var score ScoreEntry

	err := json.NewDecoder(r.Body).Decode(&score)

	if err != nil {
		log.Fatal(err)
	}

	position := strconv.Itoa(AddScore(score))
	w.Write([]byte(position))

	log.Print("Served PostScore")
}

func main() {
	localDb, err := sql.Open("mysql", CreateDsnFromConfig("settings.ini"))
	defer localDb.Close()

	if err != nil {
		log.Fatal(err)
	}

	db = localDb

	router := mux.NewRouter()

	router.HandleFunc("/map", GetMaps).Methods("GET")
	router.HandleFunc("/score/{mapId}", GetScores).Methods("GET")
	router.HandleFunc("/score/{mapId}/{limit}", GetScores).Methods("GET")
	router.HandleFunc("/score/{mapId}/{limit}/{offset}", GetScores).Methods("GET")
	router.HandleFunc("/score", PostScore).Methods("POST")

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	log.Print("started")
	log.Fatal(http.ListenAndServe(":10500", handlers.CORS(allowedOrigins, allowedHeaders, allowedMethods)(router)))
}
