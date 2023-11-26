package gogoogleplaces

import "os"

type Service struct {
	apiKey            string
	autocompleteCoins chan struct{}
	detailCoins       chan struct{}
}

func New() *Service {
	service := &Service{
		apiKey:            os.Getenv("GOOGLEAPIKEY"),
		autocompleteCoins: make(chan struct{}),
		detailCoins:       make(chan struct{}),
	}

	go service.genCoins()

	return service
}
