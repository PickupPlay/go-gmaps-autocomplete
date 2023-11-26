package gogoogleplaces

import (
	"testing"
)

func TestAndPrint(t *testing.T) {
	service := New()
	_, err := service.Autocomplete(
		"north boul",
		40.014984,
		-105.270546,
		false,
	)
	if err != nil {
		t.Error(err)
	}
}
