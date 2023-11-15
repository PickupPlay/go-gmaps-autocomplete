package gogmapsautocomplete

import "encoding/json"

type PlaceAutocompleteResponse struct {
	Predictions  []PlaceAutocompletePrediction `json:"predictions"`
	Status       interface{}                   `json:"status"`
	ErrorMessage string                        `json:"error_message"`
}

type PlacesAutocompleteStatus struct {
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
