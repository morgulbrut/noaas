package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func randomNo() string {
	nopes := make([]string, 0)
	nopes = append(nopes,
		"How 'bout no?",
		"NO",
		"Nope",
		"Nada",
		"Njet!",
		"Yesn't",
		"Nein, Nein, Nein!")
	return nopes[rand.Intn(len(nopes))]
}

func main() {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"noPhrase": randomNo(),
		})
	})

	router.GET("/txt", func(c *gin.Context) {
		c.HTML(http.StatusOK, "how 'bout no?", nil)
	})

	router.Run(":" + port)
}
