package handlers

type BaseQuestion struct {
	Name        string
	Option      *QuestionOption
	SubQuestion *[]BaseQuestion
}

type QuestionOption struct {
	Label string

	Handler func(url string, data string)
	Options func(url string) []string

	Validator func(input string, size uint32) error
}

type SubQuestion struct {
	Name string

	Questions []BaseQuestion

	Validator func(input string, size uint32) error
}

type Question struct {
	Label string

	Options []string

	Validator func(input string, size uint32) error
}
