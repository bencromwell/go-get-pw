package pwextract

func SelectedCharactersFromPassword(password string, characters []int) map[int]string {
	selected := make(map[int]string)

	for _, index := range characters {
		if index >= 1 && len(password) > index-1 {
			selected[index] = string([]rune(password)[index-1])
		}
	}

	return selected
}
