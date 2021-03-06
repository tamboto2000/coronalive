package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	scrapper "github.com/tamboto2000/coronalive/services/scrapper"
)

type Data struct {
	Errors []Error `json:"errors"`
	*scrapper.COVIDData
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//GetAll menampilkan semua data
//path: /getAll
func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

//GetByProvince menampilkan data berdasarkan provinsi
//path: /getAllByProvince
//query: prov (string)
func GetByProvince(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	covidData, err := scrapper.GetAllDataByProvince()
	if err != nil {
		json.NewEncoder(w).Encode(Error{
			Code:    500,
			Message: err.Error(),
		})

		return
	}

	prov := r.URL.Query().Get("prov")
	if prov != "" {
		for _, item := range covidData {
			if strings.ToLower(item.Province) == strings.ToLower(prov) {
				json.NewEncoder(w).Encode(item)
				return
			}
		}

		json.NewEncoder(w).Encode(covidData)
		return
	}

	json.NewEncoder(w).Encode(covidData)
}

//GetNationalSummary menampilkan rangkuman data nasional
//path: /getNationalSummary
func GetNationalSummary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	nationalSummary, err := scrapper.GetNationalSummary()
	if err != nil {
		json.NewEncoder(w).Encode(Error{
			Code:    500,
			Message: err.Error(),
		})

		return
	}

	json.NewEncoder(w).Encode(nationalSummary)
}
