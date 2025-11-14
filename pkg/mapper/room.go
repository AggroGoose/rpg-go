package mapper

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	m "rpg-go/pkg/monster"
)
type Directions struct {
	north int
	east int
	south int
	west int
}
type MapRoom struct {
	id int
	name string
	description string
	symbol byte
	directions Directions
	monsters []m.Monster
}

type ImportRoom struct {
	Id int `json:"id"`
	North int `json:"north"`
	East int `json:"east"`
	South int `json:"south"`
	West int `json:"west"`
}

func CreateRoom(id int, north int, east int, south int, west int) MapRoom {
	return MapRoom{
		id: id,
		name: "field",
		description: "An open field",
		symbol: '+',
		directions: Directions{
			north: north,
			east: east,
			south: south,
			west: west,
		},
		monsters: []m.Monster{},
	}
}

func GenerateRooms() map[int]MapRoom {
	rooms := make(map[int]MapRoom)

	file, err := os.Open("rooms.json")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	importRooms := []ImportRoom{}

	json.Unmarshal(byteValue, &importRooms)

	for i := range importRooms {
		room := CreateRoom(importRooms[i].Id, importRooms[i].North, importRooms[i].East, importRooms[i].South, importRooms[i].West)
		rooms[room.id] = room
	}
 
	return rooms
}