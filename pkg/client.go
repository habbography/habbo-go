package habbo

import (
	"fmt"
	"net/http"
)

type Hotel string

const (
	HotelDE Hotel = ".de"
	HotelUS Hotel = ".com"
	HotelBR Hotel = ".com.br"
	HotelNL Hotel = ".nl"
	HotelFR Hotel = ".fr"
)

type BaseClient struct {
	Hotel      Hotel
	BaseUrl    string
	HttpClient *http.Client
}

func NewClient(hotel Hotel) *BaseClient {
	return &BaseClient{
		Hotel:      hotel,
		BaseUrl:    fmt.Sprintf("https://habbo%s/api/public", hotel),
		HttpClient: &http.Client{},
	}
}
