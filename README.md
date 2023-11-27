# go-google-places

Middleware code for making Google Places API calls.
This module is meant to be imported by a service and run in a secure environment that has access to your Google API key.

Comes built in with a rate limiter than can be customized per API/endpoint/pricing.

# Usage

>[!WARNING]
>Refer to the code in `example/` as it is guaranteed to compile

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

## Nov 27

- Readme 'usage' and example binary

## Nov 26

- Support for Places (New) details api
  - Retrieve `location` field for a `place_id`
- Service code moved to its own home
- Per-API rate limiting and coin generation

## Nov 16
MVP done

## Nov 14
Implementation to come
