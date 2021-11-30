package handlers

import (
	"strconv"

	"github.com/paij0se/ymp3cli/src/client/validators"
)

// This function handles the questions.
func QuestionHandler(url string, questions []BaseQuestion) {

	// Create a question about what you want to do, like listening to music, downloading music or other.
	option := CreateQuestion(Question{
		Label: "What do you want to do?",
		Options: (func() []string {
			var options []string

			// The created questions, it stores them so that you can select them.
			for _, question := range questions {
				options = append(options, question.Name)

			}

			if len(options) == 0 {

				return []string{}
			}

			return options
		})(),
		Validator: validators.Number,
	})

	if option == nil {

		return
	}

	number, _ := strconv.Atoi(*option)

	question := questions[number-1]

	// If it contains subquestions, it calls the question handler and returns.
	if question.SubQuestion != nil {
		QuestionHandler(url, *question.SubQuestion)

		return

	}

	// If the question has no sub-questions, the question must have options.
	// Call the question creator with the options.
	option = CreateQuestion(Question{
		Label:     question.Option.Label,
		Options:   question.Option.Options(url),
		Validator: question.Option.Validator,
	})

	question.Option.Handler(url, *option)
}
