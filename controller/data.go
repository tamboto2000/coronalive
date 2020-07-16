package controller

import (
	"encoding/json"
	"net/http"

	"github.com/tamboto2000/coronalive/services/scrapper"
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
