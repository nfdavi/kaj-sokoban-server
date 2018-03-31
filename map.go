package main

import (
	"encoding/json"
	"log"
)

type Map struct {
	Width   int   `json:"width"`
	Height  int   `json:"height"`
	Walls   []int `json:"walls"`
	Crates  []int `json:"crates"`
	Targets []int `json:"targets"`
	Player  int   `json:"player"`
}

type MapEntry struct {
	ID   int `json:"id"`
	Data Map `json:"data"`
}

type MapDatabaseEntry struct {
	id        int
	data      string
	published int
}

func GetAllMaps() []MapEntry {
	var rtn []MapEntry

	result, err := db.Query("SELECT * FROM map WHERE published=1")

	if err != nil {
		log.Fatal(err)
	}

	for result.Next() {
		var row MapDatabaseEntry
		var mapData Map

		result.Scan(&row.id, &row.data, &row.published)

		err = json.Unmarshal([]byte(row.data), &mapData)

		if err != nil {
			log.Fatal(err)
		}

		rtn = append(rtn, MapEntry{row.id, mapData})
	}

	return rtn
}
