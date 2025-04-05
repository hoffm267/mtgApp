package dl

import (
	"context"
	"example/mtgApp/database"
	"example/mtgApp/models"
	"fmt"
)

func GetAllGames() []models.Game {
	rows, err := database.Conn.Query(context.Background(), "SELECT * FROM games")
	if err != nil {
		return nil
	}
	defer rows.Close()

	games := []models.Game{}
	for rows.Next() {
		var g models.Game
		err = rows.Scan(&g.GameID, &g.PlayerCount, &g.Date)
		if err != nil {
			return nil
		}
		games = append(games, g)
	}

	return games
}

func GetGame(gameID int) models.Game {
	query := fmt.Sprintf("%s%d", "SELECT * FROM games WHERE \"GameID\"=", gameID)
	rows, err := database.Conn.Query(context.Background(), query)
	if err != nil {
		return models.Game{}
	}
	defer rows.Close()

	rows.Next()
	var g models.Game
	err = rows.Scan(&g.GameID, &g.PlayerCount, &g.Date)
	if err != nil {
		return models.Game{}
	}

	return g
}
