package main

import (
	"AP1/config"
	"AP1/middleware"
	"AP1/routers"
	"context"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config.ConnectDB()

	err := config.DB.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	router := gin.Default()

	router.LoadHTMLFiles(
		"templates/index.html",
		"templates/musicapp/all_songs.html",
		"templates/authentification/login.html",
		"templates/authentification/signup.html",
		"templates/musicapp/mymusic.html",
		"templates/musicapp/favourite.html",
		"templates/musicapp/playlist.html",
	)

	router.Static("/css", "./static/musicapp/css")
	router.Static("/scripts", "./static/musicapp/scripts")
	router.Static("/music", "./music")
	router.Use(middleware.Logger())

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"server_url": "http://localhost:8080",
		})
	})

	router.GET("/index.html", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"server_url": "http://localhost:8080",
		})
	})

	router.GET("/musicapp/all_songs.html", func(c *gin.Context) {
		c.HTML(200, "all_songs.html", gin.H{
			"title": "All Songs",
		})
	})

	router.GET("/authentification/login.html", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{
			"title": "Login",
		})
	})

	router.GET("/authentification/signup.html", func(c *gin.Context) {
		c.HTML(200, "signup.html", gin.H{
			"title": "Sign Up",
		})
	})

	router.GET("/mymusic.html", func(c *gin.Context) {
		c.HTML(200, "mymusic.html", gin.H{
			"title": "My Music",
		})
	})
	router.GET("/favourite.html", func(c *gin.Context) {
		c.HTML(200, "favourite.html", gin.H{
			"title": "Favourite music",
		})
	})

	router.GET("/playlist.html", func(c *gin.Context) {
		c.HTML(200, "playlist.html", gin.H{
			"title": "Playlists",
		})
	})

	db := config.DB.Database("music-app")
	routers.InitRouter(router, db)

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
