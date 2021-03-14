package pwextract

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func Password() (string, error) {
	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()

	if err != nil {
		return "", err
	}

	password := string(bytePassword)

	return strings.TrimSpace(password), nil
}

func Characters() ([]int, error) {
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

	sort.Ints(characterInts)

	return characterInts, nil
}
