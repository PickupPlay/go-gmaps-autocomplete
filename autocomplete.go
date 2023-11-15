package gogmapsautocomplete

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type Service struct {
	coins chan struct{}
}

func New() *Service {
	service := &Service{
		coins: make(chan struct{}),
	}

	go service.genCoins()

	return service
}

func Autocomplete(
	query string,
	lat, lon float64,
	dry bool,
) (*PlaceAutocompleteResponse, error) {
	API_KEY := os.Getenv("GOOGLEAPIKEY")

	url, err := url.Parse(ENDPOINT)
	if err != nil {
		return nil, err
	}

	params := url.Query()
	params.Add("input", query)
	params.Add("location", fmt.Sprintf("%v,%v", lat, lon))
	params.Add("origin", fmt.Sprintf("%v,%v", lat, lon))
	params.Add("key", API_KEY)
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
	ENDPOINT = "https://maps.googleapis.com/maps/api/place/autocomplete/json"
)
