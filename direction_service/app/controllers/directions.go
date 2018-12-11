package controllers

import (
	"direction_service/app/services"
	"encoding/json"
	"net/http"
)

type Point struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type DirectionsCalculateRequest struct {
	StartPoint Point `json:"start_point"`
	EndPoint   Point `json:"end_point"`
}

func DirectionsCalculate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var request DirectionsCalculateRequest
	err := decoder.Decode(&request)

	if err != nil {
		panic(err)
	}

	service := services.DirectionsCalculateService{}
	result := service.Call(request.StartPoint.Lat, request.StartPoint.Lng, request.EndPoint.Lat, request.EndPoint.Lng)
	js, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
}
