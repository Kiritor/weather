package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	requestUrlTemplate = "http://api.openweathermap.org/data/2.5/find?q=%s&units=metric"
)

type Response struct {
	List []struct {
		Main struct {
			Temp           float32 `json:"temp"`
			Max_temp       float32 `json:"temp_max"`
			Min_temp       float32 `json:"temp_min"`

		} `json:"main"`
	} `json:"list"`
}

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c Client) Request(city string) (*Weather, error) {
	requestUrl := getUrl(city)
	resp, err := http.Get(requestUrl)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	var responseObject Response
	json.Unmarshal(body, &responseObject)
	return &Weather{Temperature: responseObject.List[0].Main.Temp,Max_temp:responseObject.List[0].Main.Max_temp,Min_temp:responseObject.List[0].Main.Min_temp}, nil
}

func getUrl(city string) string {
	return fmt.Sprintf(requestUrlTemplate, city)
}
