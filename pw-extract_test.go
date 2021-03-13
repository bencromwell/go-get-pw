package main

import (
	"reflect"
	"testing"
)

func TestSelectedCharactersFromPassword_InRange(t *testing.T) {
	password := "password42"
	characters := []int{1, 4, 5}

	result, _ := selectedCharactersFromPassword(password, characters)
	expected := map[int]string{1: "p", 4: "s", 5: "w"}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("did not retrieve correct characters from password")
	}
}

func TestSelectedCharactersFromPassword_OutOfRange(t *testing.T) {
	password := "password42"
	characters := []int{99}

	_, err := selectedCharactersFromPassword(password, characters)

	if err == nil {
		t.Fatalf("expected an error")
	}
}
