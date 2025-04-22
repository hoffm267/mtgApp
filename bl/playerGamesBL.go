package bl

import (
	"errors"
	"example/mtgApp/dl"
	"example/mtgApp/models"
)

// CREATE
func CreatePlayerGames(playergames []models.PlayerGame) ([][2]int, error) {
	if !checkPlayerGameData(playergames) {
		return nil, errors.New("invalid player data")
	}

	var playerGameIDs [][2]int
	playerGameIDs, err := dl.CreatePlayerGames(playergames)
	if err != nil {
		return nil, err
	}

	return playerGameIDs, nil
}

// READ
func GetAllPlayerGames() []models.PlayerGame {
	playergames := dl.GetAllPlayerGames()

	return playergames
}

func GetPlayerGame(playerGameID []int) models.PlayerGame {
	playergame := dl.GetPlayerGame(playerGameID)

	return playergame
}

// UPDATE
func UpdatePlayerGames(playergames []models.PlayerGame) error {
	if !checkPlayerGameData(playergames) {
		return errors.New("invalid player data")
	}

	err := dl.UpdatePlayerGames(playergames)
	if err != nil {
		return err
	}

	return nil
}

// DELETE
func DeletePlayerGames(playergames []models.PlayerGame) error {
	if !checkPlayerGameData(playergames) {
		return errors.New("invalid player data")
	}

	err := dl.DeletePlayerGames(playergames)
	if err != nil {
		return err
	}

	return nil
}

//HELPER

func checkPlayerGameData(players []models.PlayerGame) bool {
	for _, playergame := range players {
		if playergame.CommanderName == "" {
			//TODO turn this into logging
			print("Player name cannot be nil.")
			return false
		}

		if playergame.Placing < 1 {
			//TODO turn this into logging
			print("Player must be greater than 0.")
			return false
		}

		if playergame.PlayerID < 1 {
			//TODO turn this into logging
			print("PlayerID must be greater than 0.")
			return false
		}

		if playergame.GameID < 1 {
			//TODO turn this into logging
			print("GameID must be greater than 0.")
		}

	}
	return true
}
