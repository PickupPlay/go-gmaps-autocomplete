package main

import (
	"fmt"

	gogoogleplaces "github.com/PickupPlay/go-google-places"
)

func main() {
	service := gogoogleplaces.New()
	resp, _ := service.Autocomplete("starbu", 40, -100, false)
	details, _ := service.Details(resp.Predictions[0].PlaceID)
	fmt.Println(resp.Predictions[0])
	fmt.Println(details)
}
