package controller

import (
	"coronalive/services/scrapper"
	"encoding/json"
	"net/http"
)

type Data struct {
	Errors []Error `json:"errors"`
	*scrapper.COVIDData
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	var data Data
	covidData, errs := scrapper.GetAllData()
	if len(errs) > 0 {
		for _, err := range errs {
			data.Errors = append(data.Errors, Error{
				Code:    500,
				Message: err.Error(),
			})
		}

		json.NewEncoder(w).Encode(data)
		return
	}

	data.COVIDData = covidData
	json.NewEncoder(w).Encode(data)
}