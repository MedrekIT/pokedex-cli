package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func getRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request - %w", err)
	}

	req.Header.Set("content-type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching response - %w", err)
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status - %s", res.Status)
	}
	
	return res, nil
}

func getLocations(url string) (Locations, error) {
	if url == "" {
		url = urlToAPI + "/location-area"
	}
	
	res, err := getRequest(url)
	if err != nil {
		return Locations{}, err
	}
	defer res.Body.Close()

	var data Locations
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return Locations{}, fmt.Errorf("error decoding response body - %w", err)
	}
	

	return data, nil
}
