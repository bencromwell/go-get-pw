package pwextract

import (
	"fmt"
)

func SelectedCharactersFromPassword(password string, characters []int) (map[int]string, error) {
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
