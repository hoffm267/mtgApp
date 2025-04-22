package main

import (
	"example/mtgApp/controllers"
	"example/mtgApp/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	database.InitDB(os.Getenv("DB"))

	router := gin.Default()

	router.GET("/player", controllers.GetAllPlayers)
	router.GET("/player/:playerid", controllers.GetPlayer)
	router.POST("/player", controllers.CreatePlayers)
	router.POST("/player/update", controllers.UpdatePlayers)
	router.DELETE("/player", controllers.DeletePlayers)

	router.GET("/game", controllers.GetAllGames)
	router.GET("/game/:gameid", controllers.GetGame)
	router.POST("/game", controllers.CreateGames)
	router.POST("/game/update", controllers.UpdateGames)
	router.DELETE("/game", controllers.DeleteGames)

	router.GET("/playergame", controllers.GetAllPlayerGames)
	router.GET("/playergame/:playerid/:gameid", controllers.GetPlayerGame)
	router.POST("/playergame", controllers.CreatePlayerGames)
	router.POST("/playergame/update", controllers.UpdatePlayerGames)
	router.DELETE("/playergame", controllers.DeletePlayerGames)

	router.Run("localhost:8080")
}
