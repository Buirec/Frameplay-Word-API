package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var words = [9]string{
	"Alpha", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel", "India",
}

type Id struct {
	Id string `json:"id"`
}

func digitSum(num int) int {
	output := 0
	for num > 0 {
		output = output + num%10
		num = num / 10
	}

	if output > 10 {
		output = digitSum(output)
	}

	return output - 1
}

func getWord(c *gin.Context) {
	var id Id
	var output string

	if err := c.BindJSON(&id); err != nil {
		fmt.Println(err)
	}

	fmt.Println(c)

	num, err := strconv.Atoi(id.Id)
	if err != nil {
		fmt.Println(err)
	}

	if num < 10 {
		output = words[num]
	} else {
		output = words[digitSum(num)]
	}

	c.IndentedJSON(http.StatusCreated, output)
	return
}

func main() {
	router := gin.Default()

	router.POST("/word", getWord)

	router.Run("localhost:8080")
}
