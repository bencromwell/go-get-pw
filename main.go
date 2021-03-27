package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/bencromwell/go-get-pw/pwextract"
)

func main() {
	password, _ := pwextract.Password(int(syscall.Stdin))
	characters, _ := pwextract.Characters(os.Stdin)

	selected := pwextract.SelectedCharactersFromPassword(password, characters)

	for index, character := range selected {
		fmt.Printf("%d: %s\n", index, character)
	}
}
