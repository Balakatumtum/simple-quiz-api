package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Balakatumtum/simple-quiz-api/database"
	"github.com/gin-gonic/gin"
)

// show user list of questions and answers
func ShowPaper(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.ExamPaper)
}

// user able to add extra questions
func AddQuestion(c *gin.Context) {
	var newq database.Question
	if err := c.BindJSON(&newq); err != nil {
		return
	}
	database.ExamPaper = append(database.ExamPaper, newq)
	c.IndentedJSON(http.StatusCreated, newq)
}

// query user's answers in accordance to the number of questions
func Query_UserAnswers(c *gin.Context) {
	var User_input []string
	leng := len(database.ExamPaper)
	x := c.Request.URL.Query()
	for i := 1; i <= leng; i++ {
		User_input = append(User_input, x.Get(fmt.Sprint("answer", i)))
		log.Println(x.Get(fmt.Sprint("answer", i)))
	}
	CheckAnswers(c, User_input)
}

// loop user's answers to match keyanswers, if so score will be added
func CheckAnswers(c *gin.Context, jawaban []string) {
	score := 0
	leng := len(database.ExamPaper)
	for i := 0; i < leng; i++ {
		if jawaban[i] == database.ExamPaper[i].Answers[database.ExamPaper[i].CorrectAnswerID] {
			score += 1
		}
	}
	c.String(http.StatusOK, "Your correct answer is %d", score)
}

func main() {
	router := gin.Default()
	router.GET("/exampaper", ShowPaper)
	router.POST("/exampaper", AddQuestion)
	router.POST("/input", Query_UserAnswers)
	router.Run("localhost:3000") //runs on post 3000
}
