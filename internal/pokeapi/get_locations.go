package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageUrl *string) (LocationAreasResponse, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	dat, ok := c.cache.Get(url)
	if ok {
		locationAreasResponse := LocationAreasResponse{}
		err := json.Unmarshal(dat, &locationAreasResponse)
		if err != nil {
			return LocationAreasResponse{}, err
		}
		return locationAreasResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	locationAreasResponse := LocationAreasResponse{}
	err = json.Unmarshal(data, &locationAreasResponse)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	c.cache.Add(url, data)
	return locationAreasResponse, nil
}
