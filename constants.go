package main

type config struct {
	Next string
	Previous string
}

type CliCommand struct {
	name string
	description string
	callback func(*config) error
}

type Result struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type Locations struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []Result `json:"results"`
}

var urlToAPI = "https://pokeapi.co/api/v2"
