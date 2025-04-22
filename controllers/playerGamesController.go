package controllers

import (
	"encoding/json"
	"example/mtgApp/bl"
	"example/mtgApp/models"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CREATE
func CreatePlayerGames(c *gin.Context) {
	playerGames := mapPlayerGames(c)
	var playerGameIDs [][2]int
	playerGameIDs, err := bl.CreatePlayerGames(playerGames)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	c.IndentedJSON(http.StatusOK, playerGameIDs)
}

func CreateFullGame(c *gin.Context) {
	playerGames := mapPlayerGames(c)
	var size = len(playerGames)
	game := models.Game{GameID: -1, Date: time.Now(), PlayerCount: size}
	gameIDs, err1 := bl.CreateGames([]models.Game{game})
	game.GameID = gameIDs[0]
	for i := range len(playerGames) {
		playerGames[i].GameID = gameIDs[0]
	}
	_, err2 := bl.CreatePlayerGames(playerGames)
	if err1 != nil || err2 != nil {
		if err1 != nil {
			c.IndentedJSON(http.StatusBadGateway, err1)
		} else {
			c.IndentedJSON(http.StatusBadGateway, err2)
		}
	}

	fullGame := models.FullGame{PlayerGames: playerGames, Game: game}
	c.IndentedJSON(http.StatusOK, fullGame)
}

// READ
func GetAllPlayerGames(c *gin.Context) {
	playergames := bl.GetAllPlayerGames()

	c.IndentedJSON(http.StatusOK, playergames)
}

func GetPlayerGame(c *gin.Context) {
	playerID, _ := strconv.Atoi(c.Param("playerid"))
	gameID, _ := strconv.Atoi(c.Param("gameid"))

	playerGame := bl.GetPlayerGame([]int{playerID, gameID})

	c.IndentedJSON(http.StatusOK, playerGame)
}

// UPDATE
func UpdatePlayerGames(c *gin.Context) {
	playergames := mapPlayerGames(c)
	err := bl.UpdatePlayerGames(playergames)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	c.IndentedJSON(http.StatusOK, nil)
}

// DELETE
func DeletePlayerGames(c *gin.Context) {
	playergames := mapPlayerGames(c)
	err := bl.DeletePlayerGames(playergames)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	c.IndentedJSON(http.StatusOK, nil)
}

// HELPER
func mapPlayerGames(c *gin.Context) []models.PlayerGame {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	var playergames []models.PlayerGame
	err = json.Unmarshal(jsonData, &playergames)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	return playergames
}
