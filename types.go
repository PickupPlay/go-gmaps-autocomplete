package gogoogleplaces

import "encoding/json"

type PlaceAutocompleteResponse struct {
	Predictions  []PlaceAutocompletePrediction `json:"predictions"`
	Status       interface{}                   `json:"status"`
	ErrorMessage string                        `json:"error_message"`
}

type PlaceAutocompleteStatus struct {
}

type PlaceAutocompletePrediction struct {
	Description    string `json:"description"`
	PlaceID        string `json:"place_id"`
	DistanceMeters int    `json:"distance_meters"`
}

func (ac PlaceAutocompleteResponse) String() string {
	bs, _ := json.MarshalIndent(ac, "", "  ")
	return string(bs)
}

type PlaceDetailsResponse struct {
	Location Location `json:"location"`
}

type Location struct {
	Lat float64 `json:"latitude"`
	Lon float64 `json:"longitude"`
}

func (r PlaceDetailsResponse) String() string {
	bs, _ := json.MarshalIndent(r, "", "  ")
	return string(bs)
}
