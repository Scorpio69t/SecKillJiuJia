package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	platform := ""
	prompt := &survey.Select{
		Message: "What platform do you want to use?",
		Options: []string{"seckill", "jinniu", "yuemiao"},
		Default: "seckill",
	}

	survey.AskOne(prompt, &platform)

	fmt.Printf("You chose %s\n", platform)
}
