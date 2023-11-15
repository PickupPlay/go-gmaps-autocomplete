package gogmapsautocomplete

type PlaceAutocompleteResponse struct {
	Predictions []PlaceAutocompletePrediction `json:"predictions"`
}

type PlaceAutocompletePrediction struct {
	Description    string `json:"description"`
	PlaceID        string `json:"place_id"`
	DistanceMeters int    `json:"distance_meters"`
}
