package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type DirectionsCalculateService struct{}

type Result struct {
	Distance float64 `json:"distance"`
	Time     int64   `json:"time"`
}

type path struct {
	Distance float64 `json:"distance"`
	Time     int64   `json:"time"`
}

type paths []path

type graphhopperResponse struct {
	Paths paths `json:"paths"`
}

func (_ DirectionsCalculateService) Call(startPointLat, startPointLng, endPointLat, endPointLng float64) Result {
	client := &http.Client{}

	req, err := http.NewRequest("GET", os.Getenv("GRAPHHOPPER_URL"), nil)
	if err != nil {
		log.Print(err)
		panic(err)
	}

	req.Header.Add("Accept", "application/json")

	q := req.URL.Query()
	q.Add("key", os.Getenv("GRAPHHOPPER_KEY"))
	q.Add("point", fmt.Sprintf("%f,%f", startPointLat, startPointLng))
	q.Add("point", fmt.Sprintf("%f,%f", endPointLat, endPointLng))
	q.Add("vehicle", "car")
	q.Add("calc_points", "false")
	q.Add("points_encoded", "false")
	q.Add("instructions", "false")
	q.Add("optimize", "true")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	if err != nil {
		log.Print(err)
		panic(err)
	}

	defer resp.Body.Close()
	rawBody, _ := ioutil.ReadAll(resp.Body)

	var graphhopperResponse graphhopperResponse
	err = json.Unmarshal(rawBody, &graphhopperResponse)
	if err != nil {
		log.Print(err)
		panic(err)
	}

	return Result{
		Distance: graphhopperResponse.Paths[0].Distance,
		Time:     graphhopperResponse.Paths[0].Time,
	}
}
