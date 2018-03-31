package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"strconv"
)


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

	scores := GetScoresForMap(mapId)
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

func setupTempValues() {
	tempMaps = append(tempMaps, MapEntry{1, Map{8,8, []int{7, 8, 9}, []int{19}, []int{35}, 15}})
	tempMaps = append(tempMaps, MapEntry{2, Map{10,15, []int{13, 15, 36}, []int{51}, []int{92}, 44}})
	tempMaps = append(tempMaps, MapEntry{3, Map{15,12, []int{8, 9, 10, 11, 12, 64}, []int{45, 48}, []int{37, 38}, 25}})

	tempScores = append(tempScores, &ScoreEntry{1, 1, 1, "aaa bbb", 25})
	tempScores = append(tempScores, &ScoreEntry{2, 1, 2, "aaa bbb", 26})
	tempScores = append(tempScores, &ScoreEntry{3, 1, 3, "abc bbb", 27})
	tempScores = append(tempScores, &ScoreEntry{4, 1, 4, "abc bbb", 28})

	tempScores = append(tempScores, &ScoreEntry{5, 2, 1, "xaa bbb", 25})
	tempScores = append(tempScores, &ScoreEntry{6, 2, 2, "xaa bbb", 26})
	tempScores = append(tempScores, &ScoreEntry{7, 2, 3, "xbc bbb", 27})

	tempScores = append(tempScores, &ScoreEntry{8, 3, 1, "yaa bbb", 25})
	tempScores = append(tempScores, &ScoreEntry{9, 3, 2, "yaa bbb", 26})
}

func main() {
	setupTempValues()

	router := mux.NewRouter()

	router.HandleFunc("/map", GetMaps).Methods("GET")
	router.HandleFunc("/score/{mapId}", GetScores).Methods("GET")
	//router.HandleFunc("/score/{mapId}/{count}", GetScores).Methods("GET")
	router.HandleFunc("/score", PostScore).Methods("POST")

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	log.Print("started")
	log.Fatal(http.ListenAndServe(":10500", handlers.CORS(allowedOrigins, allowedHeaders, allowedMethods)(router)))
}
