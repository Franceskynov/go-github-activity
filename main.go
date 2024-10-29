package main

import (
	"os"
	"github.com/Franceskynov/go-github-activity/actions"
	"github.com/Franceskynov/go-github-activity/utils"
)

func main() {

	if utils.ArgsChecker(os.Args) {
		actions.FormatUserData(os.Args[1])
	}
}
