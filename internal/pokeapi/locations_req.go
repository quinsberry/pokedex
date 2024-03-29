package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageUrl *string) (LocationAreas, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	dat, ok := c.cache.Get(url)
	if ok {
		locationAreasResponse := LocationAreas{}
		err := json.Unmarshal(dat, &locationAreasResponse)
		if err != nil {
			return LocationAreas{}, err
		}
		return locationAreasResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreas{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 399 {
		return LocationAreas{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreas{}, err
	}
	locationAreasResponse := LocationAreas{}
	err = json.Unmarshal(data, &locationAreasResponse)
	if err != nil {
		return LocationAreas{}, err
	}
	c.cache.Add(url, data)
	return locationAreasResponse, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	url := baseUrl + "/location-area/" + locationAreaName

	dat, ok := c.cache.Get(url)
	if ok {
		locationAreaResp := LocationArea{}
		err := json.Unmarshal(dat, &locationAreaResp)
		if err != nil {
			return LocationArea{}, err
		}
		return locationAreaResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}
	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(url, data)
	return locationArea, nil
}
