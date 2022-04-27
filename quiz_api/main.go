package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type question struct { // the question object
	ID       int    `json:"id"`
	Question string `json:"question"`
	Answer1  string `json:"answer1"`
	Answer2  string `json:"answer2"`
	Answer3  string `json:"answer3"`
	Answer4  string `json:"answer4"`
}

var questions = []question{ // all our different objects
	{
		ID:       1,
		Question: "How long was Barack Obama president for?",
		Answer1:  "6 Years",
		Answer2:  "24 Years",
		Answer3:  "8 Years",
		Answer4:  "16 Years",
	}, {
		ID:       2,
		Question: "When did the programming language python release",
		Answer1:  "1989",
		Answer2:  "2000",
		Answer3:  "1995",
		Answer4:  "2005",
	},
}

func getRandomQuestion() int {
	rand.Seed(time.Now().UnixNano()) // randomises the rand item
	in := len(questions)             // gets the length of the questions variable
	pick := rand.Intn(in-1+1) + 1    // randomises the numbers rand.Intn(max - min + 1) + min)
	fmt.Println(in)
	fmt.Println(pick)
	return pick
}

func getQuestion(c *gin.Context) {
	var curr_question = getRandomQuestion() - 1 // gets the ID of a random question
	var response = questions[curr_question]     // gets the rest of the data from the ID
	c.IndentedJSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()              // Starts gin
	router.GET("/question", getQuestion) // our endpoint
	router.Run("localhost:8080")         // where we run the server
}
