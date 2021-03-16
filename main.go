package main

import (
	"fmt"

	"github.com/bencromwell/go-get-pw/pwextract"
)

func main() {
	password, _ := pwextract.Password()
	characters, _ := pwextract.Characters()

	selected := pwextract.SelectedCharactersFromPassword(password, characters)

	for index, character := range selected {
		fmt.Printf("%d: %s\n", index, character)
	}
}
