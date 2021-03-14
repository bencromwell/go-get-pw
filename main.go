package main

import (
	"fmt"
	"os"

	"github.com/bencromwell/go-get-pw/pwextract"
)

func main() {
	password, _ := pwextract.Password()
	characters, _ := pwextract.Characters()

	selected, err := pwextract.SelectedCharactersFromPassword(password, characters)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for index, character := range selected {
		fmt.Printf("%d: %s\n", index, character)
	}
}
