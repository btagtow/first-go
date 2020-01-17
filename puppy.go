package main

import (
	"github.com/gin-gonic/gin"
)

type Puppy struct {
	Name    string
	Breed   string
	Age     int
	Naughty bool
}

func handlePuppy(context *gin.Context) {
	puppy := getPuppy()
	context.JSON(200, puppy)
}

func getPuppy() Puppy {
	return Puppy{
		Name:  "Bixby@@@2",
		Breed: "Golden Retriever",
		Age:   1,
	}
}
