package models

import (
	"time"
)

type Game struct {
	GameID      int       `json:"gameid"` //primary key
	PlayerCount int       `json:"playercount"`
	Date        time.Time `json:"date"`
}

type Player struct {
	PlayerID int    `json:"playerid"` //primary key
	Name     string `json:"name"`
}

type PlayerGame struct {
	GameID    int    `json:"gameid"`   //foreign key   <
	PlayerID  int    `json:"playerid"` //foreign key < composite primary key
	Placing   int    `json:"placing"`
	Commander string `json:"commander"` //scryfall, maybe foreignkey??
}
