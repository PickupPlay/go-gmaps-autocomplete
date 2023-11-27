# go-gmaps-autocomplete
Service code for fetching autocomplete results from Google Maps.

Comes built in with a rate limiter you can tweek to your needs

# Usage

[!WARNING]
Refer to the code in `example/` as this will stay up to date.

```go
import (
    gogoogleplaces "github.com/PickupPlay/go-google-places"
)

func main() {
	service := gogoogleplaces.New()
	resp, err := service.Autocomplete("starbu", 40, -100, false)
	if err != nil {
		log.Fatal(err)
	}

	for _, prediction := range resp.Predictions {
		fmt.Println(prediction)
	}

    details, err := service.Details(resp.Predictions[0].PlaceID)
}
```

# Updates

## Nov 26

- Support for Places (New) details api
  - Retrieve `location` field for a `place_id`
- Service code moved to its own home
- Per-API rate limiting and coin generation

## Nov 16
MVP done

## Nov 14
Implementation to come
