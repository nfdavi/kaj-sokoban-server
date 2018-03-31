package main

type Map struct {
	Width int `json:"width"`
	Height int `json:"height"`
	Walls []int `json:"walls"`
	Crates []int `json:"crates"`
	Targets []int `json:"targets"`
	Player int `json:"player"`
}

type MapEntry struct {
	ID int `json:"id"`
	Data Map `json:"data"`
}

var tempMaps []MapEntry

func GetAllMaps() []MapEntry {
	return tempMaps
}