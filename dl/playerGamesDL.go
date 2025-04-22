package dl

import (
	"context"
	"example/mtgApp/database"
	"example/mtgApp/models"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
)

// CREATE
func CreatePlayerGames(playerGames []models.PlayerGame) ([][2]int, error) {
	var query strings.Builder

	query.WriteString(`INSERT INTO playergames ("GameID", "PlayerID", "Placing", "CommanderName") VALUES `)
	for i := 0; i < len(playerGames); i++ {
		query.WriteString("(")
		query.WriteString(strconv.Itoa(playerGames[i].GameID))
		query.WriteString(", ")
		query.WriteString(strconv.Itoa(playerGames[i].PlayerID))
		query.WriteString(", ")
		query.WriteString(strconv.Itoa(playerGames[i].Placing))
		query.WriteString(", '")
		query.WriteString(playerGames[i].CommanderName)
		query.WriteString("')")

		if i+1 < len(playerGames) {
			query.WriteString(", ")
		}
	}
	query.WriteString(` RETURNING "PlayerID", "GameID";`)
	var test = query.String()
	print(test)
	rows, err := database.Conn.Query(context.Background(), query.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	playerGameIDs := [][2]int{}
	for rows.Next() {
		var playerGameID [2]int
		err = rows.Scan(&playerGameID[0], &playerGameID[1])
		if err != nil {
			return nil, err
		}
		playerGameIDs = append(playerGameIDs, playerGameID)
	}

	return playerGameIDs, nil
}

// READ
func GetAllPlayerGames() []models.PlayerGame {
	rows, err := database.Conn.Query(context.Background(), `SELECT * FROM playergames`)
	if err != nil {
		return nil
	}
	defer rows.Close()

	playergames := GetPlayerGamesList(rows)

	return playergames
}

func GetPlayerGame(idArr []int) models.PlayerGame {
	query := fmt.Sprintf("%s%d%s%d", `SELECT * FROM playergames WHERE "PlayerID"=`, idArr[0], ` AND "GameID"=`, idArr[1])
	rows, err := database.Conn.Query(context.Background(), query)
	if err != nil {
		return models.PlayerGame{}
	}
	defer rows.Close()

	rows.Next()
	var pg models.PlayerGame
	err = rows.Scan(&pg.PlayerID, &pg.GameID, &pg.Placing, &pg.CommanderName)
	if err != nil {
		return models.PlayerGame{}
	}

	return pg
}

// UPDATE
func UpdatePlayerGames(playergames []models.PlayerGame) error {
	var query strings.Builder

	query.WriteString(`UPDATE playergames SET "Placing" = tmp."Placing", "CommanderName" from ( values `)
	for i := 0; i < len(playergames); i++ {
		query.WriteString("(")
		query.WriteString(strconv.Itoa(playergames[i].PlayerID))
		query.WriteString(", ")
		query.WriteString(strconv.Itoa(playergames[i].Placing))
		if i+1 < len(playergames) {
			query.Write([]byte{44, 32})
			query.WriteString("\n")
		}
	}
	query.WriteString(`) as tmp ("PlayerID", "GameID", "Placing", "CommanderName) `)
	query.WriteString(`WHERE playergames."PlayerID" = tmp."PlayerID" AND playergames."GameID" = tmp."GameID";`)

	_, err := database.Conn.Query(context.Background(), query.String())
	if err != nil {
		return err
	}

	return nil
}

// DELETE
func DeletePlayerGames(playergames []models.PlayerGame) error {
	var query strings.Builder

	query.WriteString(`DELETE FROM playergames WHERE ("PlayerID", "GameID") IN (`)
	for i := 0; i < len(playergames); i++ {
		query.WriteString("(")
		query.WriteString(strconv.Itoa(playergames[i].PlayerID))
		query.WriteString(",")
		query.WriteString(")")

		if i+1 < len(playergames) {
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

// HELPERS
func GetPlayerGamesList(rows pgx.Rows) []models.PlayerGame {
	playergames := []models.PlayerGame{}
	x := 0

	for rows.Next() {
		var pg models.PlayerGame
		err := rows.Scan(&pg.PlayerID, &pg.GameID, &pg.Placing, &pg.CommanderName)
		if err != nil {
			// return nil
		}
		playergames = append(playergames, pg)
		x++
	}
	return playergames
}
