package bl

import (
	"example/mtgApp/dl"
	"example/mtgApp/models"
)

// CREATE
func CreateGames(games []models.Game) ([]int, error) {
	gamesIDs, err := dl.CreateGames(games)
	if err != nil {
		return nil, err
	}
	return gamesIDs, nil
}

// READ
func GetAllGames() []models.Game {
	games := dl.GetAllGames()

	return games
}

func GetGame(gameID int) models.Game {
	game := dl.GetGame(gameID)

	return game
}

// UPDATE
func UpdateGames(games []models.Game) error {
	err := dl.UpdateGames(games)
	if err != nil {
		return err
	}

	return nil
}

// DELETE
func DeleteGames(games []models.Game) error {
	err := dl.DeleteGames(games)
	if err != nil {
		return err
	}

	return nil
}
