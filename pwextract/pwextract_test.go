package pwextract

import (
	"reflect"
	"testing"
)

func TestSelectedCharactersFromPassword_InRange(t *testing.T) {
	password := "password42"
	characters := []int{1, 4, 5}

	result := SelectedCharactersFromPassword(password, characters)
	expected := map[int]string{1: "p", 4: "s", 5: "w"}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("did not retrieve correct characters from password")
	}
}

func TestSelectedCharactersFromPassword_OutOfRange(t *testing.T) {
	password := "password42"
	characters := []int{99}

	result := SelectedCharactersFromPassword(password, characters)

	expected := map[int]string{}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("out of range should have resulted in an empty map")
	}
}

func TestSelectedCharactersFromPassword_IndexZeroIsIgnored(t *testing.T) {
	password := "password42"
	characters := []int{0}

	result := SelectedCharactersFromPassword(password, characters)

	expected := map[int]string{}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("index 0 being requested should not cause a panic")
	}
}
