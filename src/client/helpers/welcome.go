package helpers

import (
	"fmt"
)

// That prints a welcome message to the user
func Welcome() {
	ymp3cli := `

██╗   ██╗███╗   ███╗██████╗ ██████╗  ██████╗██╗     ██╗
╚██╗ ██╔╝████╗ ████║██╔══██╗╚════██╗██╔════╝██║     ██║
 ╚████╔╝ ██╔████╔██║██████╔╝ █████╔╝██║     ██║     ██║
  ╚██╔╝  ██║╚██╔╝██║██╔═══╝  ╚═══██╗██║     ██║     ██║
   ██║   ██║ ╚═╝ ██║██║     ██████╔╝╚██████╗███████╗██║
   ╚═╝   ╚═╝     ╚═╝╚═╝     ╚═════╝  ╚═════╝╚══════╝╚═╝
                                                       


	`
	fmt.Println(ymp3cli)
	fmt.Printf("\nwelcome to ymp3cli!\n\n")
	fmt.Printf("Version 0.2.0\n")
	fmt.Printf("\nType <ctrl + c> to exit.\n\n")
}
