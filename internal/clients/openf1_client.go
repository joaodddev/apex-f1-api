package clients

import (
	"github.com/go-resty/resty/v2"
)

func GetDrivers() (string, error) {
	client := resty.New()

	resp, err := client.R().
		Get("https://api.openf1.org/v1/drivers")

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}
