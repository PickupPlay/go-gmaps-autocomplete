package gogmapsautocomplete

import (
	"fmt"
	"testing"
)

func TestAndPrint(t *testing.T) {
	ac, err := Autocomplete(
		"north boulder pa",
		40.014984,
		-105.270546,
		false,
	)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(ac)
}
