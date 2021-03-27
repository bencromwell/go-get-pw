package pwextract

import (
	"reflect"
	"strings"
	"testing"
)

func TestCharacters_readsStringIntoSortedInts(t *testing.T) {
	reader := strings.NewReader("3 1 5 4 9 7")
	ints, _ := Characters(reader)

	expected := []int{1, 3, 4, 5, 7, 9}

	if !reflect.DeepEqual(ints, expected) {
		t.Errorf("should have given ascending array of ints from the input")
	}
}
