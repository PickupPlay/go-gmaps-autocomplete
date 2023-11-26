package gogoogleplaces

import "time"

const (
	// This should be 1 request every 35 seconds to stay within the $200 credit
	AC_COIN_RATE = time.Duration(20 * time.Second)
	DT_COIN_RATE = time.Duration(30 * time.Second)
)

func (s *Service) genCoins() {
	go genAPICoins(s.autocompleteCoins, AC_COIN_RATE)
	go genAPICoins(s.detailCoins, DT_COIN_RATE)
}

func genAPICoins(coinChan chan<- struct{}, rate time.Duration) {
	curr := time.Now()

	for {
		now := time.Now()
		if now.Sub(curr) > rate {
			coinChan <- struct{}{}
			curr = curr.Add(rate)
		} else {
			time.Sleep(curr.Add(rate).Sub(now))
		}
	}
}
