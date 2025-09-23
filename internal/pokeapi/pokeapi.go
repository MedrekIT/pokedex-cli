package pokeapi

import (
	"io"
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

func GetLocations(url string, conf *Config) (Locations, error) {
	if url == "" {
		url = urlToAPI + "location-area"
	}

	if dataBytes, ok := conf.Cache.Get(url); ok {
		var data Locations
		if err := json.Unmarshal(dataBytes, &data); err != nil {
			return Locations{}, fmt.Errorf("error decoding response body - %w", err)
		}

		fmt.Println("(data loaded from cache)")
		return data, nil
	}

	res, err := getRequest(url)
	if err != nil {
		return Locations{}, err
	}
	defer res.Body.Close()

	dataBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return Locations{}, fmt.Errorf("error reading data from body: %w", err)
	}
	conf.Cache.Add(url, dataBytes)

	var data Locations
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		return Locations{}, fmt.Errorf("error decoding response body - %w", err)
	}


	return data, nil
}

func GetPokemons(url string, conf *Config) (Area, error) {
	fullURL := urlToAPI + "location-area/" + url

	if dataBytes, ok := conf.Cache.Get(fullURL); ok {
		var data Area
		if err := json.Unmarshal(dataBytes, &data); err != nil {
			return Area{}, fmt.Errorf("error decoding response body - %w", err)
		}

		fmt.Println("(data loaded from cache)")
		return data, nil
	}

	res, err := getRequest(fullURL)
	if err != nil {
		return Area{}, err
	}
	defer res.Body.Close()

	dataBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return Area{}, fmt.Errorf("error reading data from body: %w", err)
	}
	conf.Cache.Add(fullURL, dataBytes)

	var data Area
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		return Area{}, fmt.Errorf("error decoding response body - %w", err)
	}


	return data, nil
}

func GetPokemon(url string, conf *Config) (Pokemon, error) {
	fullURL := urlToAPI + "pokemon/" + url

	if dataBytes, ok := conf.Cache.Get(fullURL); ok {
		var data Pokemon
		if err := json.Unmarshal(dataBytes, &data); err != nil {
			return Pokemon{}, fmt.Errorf("error decoding response body - %w", err)
		}

		fmt.Println("(data loaded from cache)")
		return data, nil
	}

	res, err := getRequest(fullURL)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	dataBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error reading data from body: %w", err)
	}
	conf.Cache.Add(fullURL, dataBytes)

	var data Pokemon
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		return Pokemon{}, fmt.Errorf("error decoding response body - %w", err)
	}


	return data, nil
}
