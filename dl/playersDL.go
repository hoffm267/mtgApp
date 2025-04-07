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
func CreatePlayers(players []models.Player) ([]int, error) {
	var query strings.Builder

	query.WriteString(`INSERT INTO players ("Name") VALUES `)
	for i := 0; i < len(players); i++ {
		query.Write([]byte{40, 39})
		query.WriteString(players[i].Name)
		query.Write([]byte{39, 41})
		if i+1 < len(players) {
			query.Write([]byte{44, 32})
		}
	}
	query.WriteString(` RETURNING "PlayerID";`)

	rows, err := database.Conn.Query(context.Background(), query.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	playerIDs := []int{}
	for rows.Next() {
		var playerID int
		err = rows.Scan(&playerID)
		if err != nil {
			return nil, err
		}
		playerIDs = append(playerIDs, playerID)
	}

	return playerIDs, nil
}

// READ
func GetAllPlayers() []models.Player {
	rows, err := database.Conn.Query(context.Background(), "SELECT * FROM players")
	if err != nil {
		return nil
	}
	defer rows.Close()

	players := GetPlayerList(rows)

	return players
}

func GetPlayer(playerID int) models.Player {
	query := fmt.Sprintf("%s%d", `SELECT * FROM players WHERE "PlayerID"=`, playerID)
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

// UPDATE
func UpdatePlayers(players []models.Player) error {
	var query strings.Builder

	query.WriteString(`UPDATE players SET "Name" = tmp."Name" from ( values `)
	for i := 0; i < len(players); i++ {
		query.WriteString("(")
		query.WriteString(strconv.Itoa(players[i].PlayerID))
		query.WriteString(", ")
		query.Write([]byte{39})
		query.WriteString(players[i].Name)
		query.Write([]byte{39, 41})
		if i+1 < len(players) {
			query.Write([]byte{44, 32})
			query.WriteString("\n")
		}
	}
	query.WriteString(`) as tmp ("PlayerID", "Name") `)
	query.WriteString(`WHERE players."PlayerID" = tmp."PlayerID";`)

	_, err := database.Conn.Query(context.Background(), query.String())
	if err != nil {
		return err
	}

	return nil
}

// DELETE
func DeletePlayers(players []models.Player) error {
	var query strings.Builder

	query.WriteString(`DELETE FROM players WHERE "PlayerID" IN (`)
	for i := 0; i < len(players); i++ {
		query.WriteString(strconv.Itoa(players[i].PlayerID))

		if i+1 < len(players) {
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
func GetPlayerList(rows pgx.Rows) []models.Player {
	players := []models.Player{}
	x := 0

	for rows.Next() {
		var p models.Player
		err := rows.Scan(&p.PlayerID, &p.Name)
		if err != nil {
			// return nil
		}
		players = append(players, p)
		x++
	}
	return players
}
