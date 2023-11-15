package gogmapsautocomplete

import "time"

const (
	// This should be 1 request every 35 seconds to stay within the $200 credit
	COIN_RATE = time.Duration(20 * time.Second)
)

func (service *Service) genCoins() {
	curr := time.Now()

	for {
		now := time.Now()
		if now.Sub(curr) > COIN_RATE {
			service.coins <- struct{}{}
			curr = curr.Add(COIN_RATE)
		} else {
			time.Sleep(curr.Add(COIN_RATE).Sub(now))
		}
	}
}
