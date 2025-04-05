package dl

import (
	"context"
	"example/mtgApp/database"
	"example/mtgApp/models"
	"fmt"
)

func GetAllPlayers() []models.Player {
	rows, err := database.Conn.Query(context.Background(), "SELECT * FROM players")
	if err != nil {
		return nil
	}
	defer rows.Close()

	players := []models.Player{}
	for rows.Next() {
		var p models.Player
		err = rows.Scan(&p.PlayerID, &p.Name)
		if err != nil {
			return nil
		}
		players = append(players, p)
	}

	return players
}

func GetPlayer(playerID int) models.Player {
	query := fmt.Sprintf("%s%d", "SELECT * FROM players WHERE \"PlayerID\"=", playerID)
	rows, err := database.Conn.Query(context.Background(), query)
	if err != nil {
		return models.Player{}
	}
	defer rows.Close()

	rows.Next()
	var p models.Player
	err = rows.Scan(&p.PlayerID, &p.Name)
	if err != nil {
		return models.Player{}
	}

	return p
}
