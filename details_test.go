package gogoogleplaces

import (
	"fmt"
	"testing"
)

func TestDetails(t *testing.T) {
	service := New()

	place, err := service.Autocomplete("North Boulder P", 40, -100, false)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(place.Predictions[0])

	dt, err := service.Details(place.Predictions[0].PlaceID)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(dt)
}
