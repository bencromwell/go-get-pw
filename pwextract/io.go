package pwextract

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/term"
)

func Password(fileDescriptor int) (string, error) {
	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(fileDescriptor)
	fmt.Println()

	if err != nil {
		return "", err
	}

	password := string(bytePassword)

	return strings.TrimSpace(password), nil
}

func Characters(reader io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(reader)

	fmt.Print("Enter character numbers, space separated: ")

	scanner.Scan()

	characterString := scanner.Text()

	characters := strings.Fields(characterString)

	var characterInts []int

	for _, character := range characters {
		if i, err := strconv.Atoi(character); err == nil {
			characterInts = append(characterInts, i)
		}
	}

	sort.Ints(characterInts)

	return characterInts, nil
}
