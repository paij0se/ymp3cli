package helpers

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// That prints a welcome message to the user
func Welcome() {
	file, err := os.Open("src/client/welcome.txt")

	if err != nil {
		log.Println(err)

		return
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)

	if err != nil {
		log.Println(err)

		return
	}

	fmt.Println(string(data))
	fmt.Printf("Version 0.0.5\n\n")

	fmt.Println("Type 'help' to see the list of commands.")
	fmt.Printf("\nType <ctrl + c> to exit.\n\n")
}
