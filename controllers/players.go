package controllers

import (
	"example/mtgApp/bl"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllPlayers(c *gin.Context) {
	players := bl.GetAllPlayers()

	c.IndentedJSON(http.StatusOK, players)
}

func GetPlayer(c *gin.Context) {
	playerID, _ := strconv.Atoi(c.Param("playerid"))
	player := bl.GetPlayer(playerID)

	c.IndentedJSON(http.StatusOK, player)
}
