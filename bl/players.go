package bl

import (
	"example/mtgApp/dl"
	"example/mtgApp/models"
)

func GetAllPlayers() []models.Player {
	players := dl.GetAllPlayers()

	return players
}

func GetPlayer(playerID int) models.Player {
	player := dl.GetPlayer(playerID)

	return player
}
