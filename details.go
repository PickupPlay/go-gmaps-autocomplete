package gogoogleplaces

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	DT_ENDPOINT = "https://places.googleapis.com/v1/places/"
)

func (s *Service) Details(
	placeID string,
) (*PlaceDetailsResponse, error) {
	<-s.detailCoins

	url, err := url.Parse(DT_ENDPOINT + placeID)
	if err != nil {
		return nil, err
	}

	params := url.Query()
	params.Add("fields", "location")
	params.Add("key", s.apiKey)
	url.RawQuery = params.Encode()

	response, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	dtResponse := &PlaceDetailsResponse{}

	if err := json.NewDecoder(response.Body).Decode(dtResponse); err != nil {
		return nil, err
	}

	return dtResponse, nil
}
