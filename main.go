package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type TimingsResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   struct {
		Timings struct {
			Fajr     string `json:"Fajr"`
			Sunrise  string `json:"Sunrise"`
			Dhuhr    string `json:"Dhuhr"`
			Asr      string `json:"Asr"`
			Sunset   string `json:"Sunset"`
			Maghrib  string `json:"Maghrib"`
			Isha     string `json:"Isha"`
			Imsak    string `json:"Imsak"`
			Midnight string `json:"Midnight"`
		} `json:"timings"`
	} `json:"data"`
}

const (
	CITY    = "Tunis"
	COUNTRY = "Tunisia"

	FILENAME      = "timings.json"
	COMP_USERNAME = "joe"
	PATH          = "/home/" + COMP_USERNAME + "/.local/share/"

	FULLPATH = PATH + FILENAME
)

func main() {
	url := fmt.Sprintf("http://api.aladhan.com/v1/timingsByCity?city=%s&country=%s&method=3&adjustment=1", CITY, COUNTRY)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	var timingResponse TimingsResponse
	err = json.Unmarshal(body, &timingResponse)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	times := timingResponse.Data.Timings
	file, err := os.Create(FULLPATH)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	json, _ := json.Marshal(times)
	file.Write(json)

	os.Exit(0)

}

/* vim: set noai ts=4 sw=4: */
