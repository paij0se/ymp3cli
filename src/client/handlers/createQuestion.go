package handlers

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

// Create a new question.
func CreateQuestion(question Question) *string {
	var sizeOptions uint32 = 1

	// Create the dynamic options.
	for _, option := range question.Options {
		fmt.Printf("[%d] %s\n", sizeOptions, option)

		sizeOptions++
	}

	// Create the static options
	for _, option := range [1]string{"Exit"} {
		fmt.Printf("[%d] %s\n", sizeOptions, option)

		sizeOptions++
	}

	// Create the prompt.
	prompt := promptui.Prompt{
		Label: question.Label,
		// Validate the question with the granted controller.
		Validate: func(input string) error {
			return question.Validator(input, sizeOptions)
		},
	}

	// Get the answer.
	input, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)

		return nil
	}

	return &input
}
