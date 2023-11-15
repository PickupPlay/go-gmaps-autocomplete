package gogmapsautocomplete

import "time"

const (
	COIN_RATE = time.Duration(1 * time.Second)
)

func (service *Service) genCoins() {
	curr := time.Now()

	for {
		now := time.Now()
		if now.Sub(curr) >= COIN_RATE {
			service.coins <- struct{}{}
		}
		curr = curr.Add(COIN_RATE)
	}
}
