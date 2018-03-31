package main

type ScoreEntry struct {
	ID int  `json:"-"`
	MapID int `json:"mapId"`
	Position int `json:"position"`
	Name string `json:"name"`
	Moves int `json:"moves"`
}

var tempScores []*ScoreEntry

func GetScoresForMap(mapId int) []*ScoreEntry {
	var rtn []*ScoreEntry

	for _, item := range tempScores {
		if item.MapID == mapId {
			rtn = append(rtn, item)
		}
	}

	return rtn
}

func AddScore(score ScoreEntry) int {
	itemIx := len(tempScores)
	score.ID = itemIx

	tempScores = append(tempScores, &score)

	mapScores := GetScoresForMap(score.MapID)

	score.Position = len(mapScores)

	return score.Position
}