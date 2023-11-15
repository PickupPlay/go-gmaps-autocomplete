package gogmapsautocomplete

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

func Autocomplete(query string) error {
	return nil
}
