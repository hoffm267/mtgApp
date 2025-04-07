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
func CreatePlayers(c *gin.Context) {
	players := mapPlayers(c)
	err := bl.CreatePlayers(players)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	c.IndentedJSON(http.StatusOK, nil)
}

// READ
func GetAllPlayers(c *gin.Context) {
	players := bl.GetAllPlayers()

	c.IndentedJSON(http.StatusOK, players)
}

func GetPlayer(c *gin.Context) {
	playerID, _ := strconv.Atoi(c.Param("playerid"))
	player := bl.GetPlayer(playerID)

	c.IndentedJSON(http.StatusOK, player)
}

// UPDATE
func UpdatePlayers(c *gin.Context) {
	players := mapPlayers(c)
	err := bl.UpdatePlayers(players)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	c.IndentedJSON(http.StatusOK, nil)
}

// DELETE
func DeletePlayers(c *gin.Context) {
	players := mapPlayers(c)
	err := bl.DeletePlayers(players)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	c.IndentedJSON(http.StatusOK, nil)
}

// HELPER
func mapPlayers(c *gin.Context) []models.Player {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	var players []models.Player
	err = json.Unmarshal(jsonData, &players)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	return players
}
