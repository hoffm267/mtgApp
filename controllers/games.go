package controllers

import (
	"example/mtgApp/bl"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllGames(c *gin.Context) {
	games := bl.GetAllGames()

	c.IndentedJSON(http.StatusOK, games)
}

func GetGame(c *gin.Context) {
	gameID, _ := strconv.Atoi(c.Param("gameid"))
	game := bl.GetGame(gameID)

	c.IndentedJSON(http.StatusOK, game)
}
