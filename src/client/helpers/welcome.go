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
	fmt.Printf("\nwelcome to ymp3cli!\n\n")
	fmt.Printf("Version 0.0.6\n\n")
	fmt.Printf("\nType <ctrl + c> to exit.\n\n")
}
