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
	PlayerID      int    `json:"playerid"`
	GameID        int    `json:"gameid"`
	Placing       int    `json:"placing"`
	CommanderName string `json:"commandername"` //scryfall, maybe foreignkey??
}

type FullGame struct {
	PlayerGames []PlayerGame
	Game        Game
}
