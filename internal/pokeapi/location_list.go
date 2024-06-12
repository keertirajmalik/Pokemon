package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)


type Client struct {
   httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
    return Client{
        httpClient: http.Client{
            Timeout: timeout,
        },
    }
}

func (c *Client) ListLocations(pageURL *string) (RespShallowLocation, error) {

    url := baseURL + "/location-area"
    if pageURL != nil {
        url = *pageURL
    }

    req, err := http.NewRequest("GET", url, nil)

    if err != nil {
        return RespShallowLocation{}, err
    }

    res, err := c.httpClient.Do(req)
    if err != nil {
        return RespShallowLocation{}, err
    }
    defer res.Body.Close()

    data, err := io.ReadAll(res.Body)
    if err != nil {
        return RespShallowLocation{}, err
    }

    locationResp := RespShallowLocation{}
    err = json.Unmarshal(data, &locationResp)

    if err != nil {
        return RespShallowLocation{}, err
    }

    return locationResp, nil
}
