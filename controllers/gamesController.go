package controllers

import (
	"encoding/json"
	"example/mtgApp/bl"
	"example/mtgApp/models"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CREATE
func CreateGames(c *gin.Context) {
	games := mapGames(c)
	gameIDs, err := bl.CreateGames(games)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	c.IndentedJSON(http.StatusOK, gameIDs)
}

// READ
func GetAllGames(c *gin.Context) {
	games := bl.GetAllGames()

	c.IndentedJSON(http.StatusOK, games)
}

func GetGame(c *gin.Context) {
	gameID, _ := strconv.Atoi(c.Param("gameid"))
	game := bl.GetGame(gameID)

	c.IndentedJSON(http.StatusOK, game)
}

// UPDATE
func UpdateGames(c *gin.Context) {
	games := mapGames(c)
	err := bl.UpdateGames(games)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	c.IndentedJSON(http.StatusOK, nil)
}

// DELETE
func DeleteGames(c *gin.Context) {
	games := mapGames(c)
	err := bl.DeleteGames(games)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	c.IndentedJSON(http.StatusOK, nil)
}

// HELPER
func mapGames(c *gin.Context) []models.Game {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	var games []models.Game
	err = json.Unmarshal(jsonData, &games)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	return games
}
