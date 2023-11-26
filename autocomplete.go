package gogoogleplaces

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (s *Service) Autocomplete(
	query string,
	lat, lon float64,
	dry bool,
) (*PlaceAutocompleteResponse, error) {
	<-s.autocompleteCoins

	url, err := url.Parse(AC_ENDPOINT)
	if err != nil {
		return nil, err
	}

	params := url.Query()
	params.Add("input", query)
	params.Add("location", fmt.Sprintf("%v,%v", lat, lon))
	params.Add("origin", fmt.Sprintf("%v,%v", lat, lon))
	params.Add("key", s.apiKey)
	url.RawQuery = params.Encode()

	if dry {
		params.Del("key")
		url.RawQuery = params.Encode()
		fmt.Println(url)
		return nil, nil
	}

	response, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	autocompleteResponse := &PlaceAutocompleteResponse{}

	err = json.NewDecoder(response.Body).Decode(autocompleteResponse)
	if err != nil {
		return nil, err
	}

	return autocompleteResponse, nil
}

const (
	AC_ENDPOINT = "https://maps.googleapis.com/maps/api/place/autocomplete/json"
)
