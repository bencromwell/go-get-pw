package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	password, _ := password()
	characters, _ := characters()

	selected, err := selectedCharactersFromPassword(password, characters)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for index, character := range selected {
		fmt.Printf("%d: %s\n", index, character)
	}
}

func password() (string, error) {
	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))

	if err != nil {
		return "", err
	}

	password := string(bytePassword)

	return strings.TrimSpace(password), nil
}

func characters() ([]int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter character numbers, space separated: ")

	characterString, err := reader.ReadString('\n')

	if err != nil {
		return []int{}, err
	}

	characters := strings.Fields(characterString)

	var characterInts []int

	for _, character := range characters {
		var i int
		if i, err = strconv.Atoi(character); err == nil {
			characterInts = append(characterInts, i)
		}
	}

	return characterInts, nil
}

func selectedCharactersFromPassword(password string, characters []int) (map[int]string, error) {
	selected := make(map[int]string)

	for _, index := range characters {
		if len(password) > index-1 {
			selected[index] = string([]rune(password)[index-1])
		} else {
			return nil, fmt.Errorf("requested character %d out of range", index)
		}
	}

	return selected, nil
}
