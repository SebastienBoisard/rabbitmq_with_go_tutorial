package main

import (
	"fmt"
	"github.com/SebastienBoisard/rabbitmq_with_go_tutorial/tutorial"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Error: the parameter indicating the tutorial id is needed (between 1 and 4)")
		return
	}

	tutorialId := os.Args[1]

	switch tutorialId {
	case "1":
		tutorial.PlayTutorial01()

	case "2":
		tutorial.PlayTutorial02()

	case "3":
		tutorial.PlayTutorial03()

	case "4":
		tutorial.PlayTutorial04()

	default:
		fmt.Println("Error: tutorial ID must be between 1 and 4")
	}
}
