package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func parseTopAwait() map[string]interface{} {

	// url := "https://kinopoiskapiunofficial.tech/api/v2.2/films/top?type=TOP_AWAIT_FILMS&page=1"

	// url := "https://kinopoiskapiunofficial.tech/api/v2.2/films/premieres?year=2023&month=JANUARY"
	url := "https://kinopoiskapiunofficial.tech/api/v2.2/films/top?type=TOP_250_BEST_FILMS&page=1"

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//Handle Error
	}

	req.Header.Set("X-API-KEY", "4bc9e0d4-62ff-417c-b947-30137719aef9")
	req.Header.Set("Content-Type", "application/json")

	// req.Header = http.Header{
	// 	"X-API-KEY":    {"264c8737-1442-470e-9ba0-ac598995deaf"},
	// 	"Content-Type": {"application/json"},
	// }

	res, err := client.Do(req)
	if err != nil {
		//Handle Error
	}

	var result map[string]interface{}

	json.NewDecoder(res.Body).Decode(&result)
	fmt.Println(res.Status)
	fmt.Println(result)

	return result
}

func parseActors(id string) []interface{} {

	url := "https://kinopoiskapiunofficial.tech/api/v1/staff?filmId=" + id

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//Handle Error
	}

	req.Header.Set("X-API-KEY", "4bc9e0d4-62ff-417c-b947-30137719aef9")
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		//Handle Error
	}

	var result []interface{}

	json.NewDecoder(res.Body).Decode(&result)
	fmt.Println(res.Status)
	fmt.Println(result)

	fmt.Println(url)
	return result
}

func parseFilm(id string) map[string]interface{} {

	url := "https://kinopoiskapiunofficial.tech/api/v2.2/films/" + id

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//Handle Error
	}

	req.Header.Set("X-API-KEY", "4bc9e0d4-62ff-417c-b947-30137719aef9")
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		//Handle Error
	}

	var result map[string]interface{}

	json.NewDecoder(res.Body).Decode(&result)
	fmt.Println(res.Status)
	fmt.Println(result)

	fmt.Println(url)
	return result
}
