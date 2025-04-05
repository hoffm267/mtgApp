package bl

import (
	"example/mtgApp/dl"
	"example/mtgApp/models"
)

func GetAllGames() []models.Game {
	games := dl.GetAllGames()

	return games
}

func GetGame(gameID int) models.Game {
	game := dl.GetGame(gameID)

	return game
}
