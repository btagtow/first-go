package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	//capital D means it's exported
	//:= creates a new variable (only inside a function)
	//go likes single letter variables:
	// r := gin.Default()
	router := gin.Default()
	// router.GET("/ping", func(c *gin.Context) {
	router.GET("/puppy", handlePuppy)
	router.Run()
}
