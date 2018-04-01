package main

import "log"

type ScoreEntry struct {
	MapID    int    `json:"mapId"`
	Position int    `json:"position"`
	Name     string `json:"name"`
	Moves    int    `json:"moves"`
}

func GetAllScoresForMap(mapId int) []ScoreEntry {
	rtn := make([]ScoreEntry, 0)

	result, err := db.Query("SELECT  mapId, name, moves FROM vwScoresForMap WHERE mapId = ?", mapId)

	if err != nil {
		log.Fatal(err)
	}

	for result.Next() {
		var row ScoreEntry

		result.Scan(&row.MapID, &row.Name, &row.Moves)
		row.Position = len(rtn) + 1

		rtn = append(rtn, row)
	}

	return rtn
}

func GetScoresForMap(mapId int, limit int, offset int) []ScoreEntry {
	rtn := make([]ScoreEntry, 0)

	result, err := db.Query("SELECT  mapId, name, moves FROM vwScoresForMap WHERE mapId = ? LIMIT ? OFFSET ?", mapId, limit, offset)

	if err != nil {
		log.Fatal(err)
	}

	for result.Next() {
		var row ScoreEntry

		result.Scan(&row.MapID, &row.Name, &row.Moves)
		row.Position = len(rtn) + 1 + offset

		rtn = append(rtn, row)
	}

	return rtn
}

func AddScore(score ScoreEntry) int {
	result, err := db.Exec("INSERT INTO score (mapId, name, moves) VALUES (?, ?, ?)", score.MapID, score.Name, score.Moves)

	if err != nil {
		log.Fatal(err)
	}

	insertedRecordId, _ := result.LastInsertId()

	var position int
	err = db.QueryRow("SELECT funGetScoreIdPosition(?)", insertedRecordId).Scan(&position)

	if err != nil {
		log.Fatal(err)
	}

	return position
}
