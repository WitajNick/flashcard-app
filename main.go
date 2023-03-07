package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type flashcard struct {
	ID           string      `json:"id"`
	SpeechPart   string      `json:"speechPart"`
	English      string      `json:"english"`
	Chinese      string      `json:"chinese"`
	Pinyin       string      `json:"pinyin"`
}

var flashcards = []flashcard{
	{ID: "1", SpeechPart: "Noun", English: "Word", Chinese: "单词", Pinyin: "Dāncí"},
	{ID: "2", SpeechPart: "Adjective", English: "Related", Chinese: "相关的", Pinyin: "Xiāngguān de"},
	{ID: "3", SpeechPart: "Adverb", English: "Maybe", Chinese: "或许", Pinyin: "Huòxǔ"},
}

func getFlashcards(c *gin.Context){
	c.IndentedJSON(http.StatusOK, flashcards)
}

func postFlashcards(c *gin.Context) {
	var newFlashcard flashcard

	if err := c.BindJSON(&newFlashcard); err != nil {
		return
	}

	flashcards = append(flashcards, newFlashcard)
	c.IndentedJSON(http.StatusCreated, newFlashcard)
}

func getFlashcardById(c *gin.Context) {
	id := c.Param("id")
	for _, a := range flashcards {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/flashcards", getFlashcards)
	r.POST("/flashcards", postFlashcards)
	r.GET("/flashcards/:id", getFlashcardById)

	r.Run() // listen on 0.0.0.0:8080
}
