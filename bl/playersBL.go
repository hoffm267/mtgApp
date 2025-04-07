package bl

import (
	"errors"
	"example/mtgApp/dl"
	"example/mtgApp/models"
)

// CREATE
func CreatePlayers(players []models.Player) ([]int, error) {
	if !checkPlayerData(players) {
		return nil, errors.New("invalid player data")
	}

	var playerIDs []int
	playerIDs, err := dl.CreatePlayers(players)
	if err != nil {
		return nil, err
	}

	return playerIDs, nil
}

// READ
func GetAllPlayers() []models.Player {
	players := dl.GetAllPlayers()

	return players
}

func GetPlayer(playerID int) models.Player {
	player := dl.GetPlayer(playerID)

	return player
}

// UPDATE
func UpdatePlayers(players []models.Player) error {
	if !checkPlayerData(players) {
		return errors.New("invalid player data")
	}

	err := dl.UpdatePlayers(players)
	if err != nil {
		return err
	}

	return nil
}

// DELETE
func DeletePlayers(players []models.Player) error {
	if !checkPlayerData(players) {
		return errors.New("invalid player data")
	}

	err := dl.DeletePlayers(players)
	if err != nil {
		return err
	}

	return nil
}

//HELPER

func checkPlayerData(players []models.Player) bool {
	for _, player := range players {
		if player.Name == "" {
			//TODO turn this into logging
			print("Player name cannot be nil.")
			return false
		}

		if player.PlayerID < 1 {
			//TODO turn this into logging
			print("PlayerID must be greater than 0.")
			return false
		}

	}
	return true
}
