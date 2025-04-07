package dl

import (
	"context"
	"example/mtgApp/database"
	"example/mtgApp/models"
	"fmt"
	"strconv"
	"strings"
)

// CREATE
func CreateGames(games []models.Game) ([]int, error) {
	var query strings.Builder

	query.WriteString(`INSERT INTO games ("PlayerCount", "Date") VALUES `)
	for i := 0; i < len(games); i++ {
		query.WriteString("(")
		query.WriteString(strconv.Itoa(games[i].PlayerCount))
		query.WriteString(", '")
		time := games[i].Date.Format("2006-01-02T15:04:05")
		query.WriteString(time)
		query.WriteString("')")
		if i+1 < len(games) {
			query.WriteString(", ")
		}
	}
	query.WriteString(`RETURNING "GameID";`)
	tmp := query.String()
	print(tmp)

	rows, err := database.Conn.Query(context.Background(), query.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	gameIDs := []int{}
	for rows.Next() {
		var gameID int
		err = rows.Scan(&gameID)
		if err != nil {
			return nil, err
		}
		gameIDs = append(gameIDs, gameID)
	}

	return gameIDs, nil
}

// READ
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

// UPDATE
func UpdateGames(games []models.Game) error {
	var query strings.Builder

	query.WriteString(`UPDATE games SET "Date" = tmp."Date", "PlayerCount" = tmp."PlayerCount" from ( values `)
	for i := 0; i < len(games); i++ {
		query.WriteString("(")
		query.WriteString(strconv.Itoa(games[i].GameID))
		query.WriteString(", ")
		query.WriteString(strconv.Itoa(games[i].PlayerCount))
		query.WriteString(", '")
		// query.WriteString(games[i].Date.String())
		time := games[i].Date.Format("2006-01-02T15:04:05")
		query.WriteString(time)
		query.WriteString("'::date)")

		if i+1 < len(games) {
			query.WriteString(", ")
			query.WriteString("\n")
		}
	}
	query.WriteString(` ) as tmp ("GameID", "PlayerCount", "Date") `)
	query.WriteString(`WHERE games."GameID" = tmp."GameID";`)
	str := query.String()
	print(str)

	_, err := database.Conn.Query(context.Background(), query.String())
	if err != nil {
		return err
	}

	return nil
}

// DELETE
func DeleteGames(games []models.Game) error {
	var query strings.Builder

	query.WriteString(`DELETE FROM games WHERE "GameID" IN (`)
	for i := 0; i < len(games); i++ {
		query.WriteString(strconv.Itoa(games[i].GameID))

		if i+1 < len(games) {
			query.WriteString(", ")
		}
	}
	query.WriteString(");")

	_, err := database.Conn.Query(context.Background(), query.String())
	if err != nil {
		return err
	}

	return nil
}
