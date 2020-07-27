package scrapper

import (
	"sync"

	"github.com/tamboto2000/coronalive/services/scrapper/raw"
)

//GetAllData controller untuk menampilkan data
func GetAllData() (*COVIDData, []error) {
	wg := new(sync.WaitGroup)
	mutex := new(sync.Mutex)
	wg.Add(3)

	updateJSON := new(raw.DataFromUpdateJSON)
	provJSON := new(raw.DataFromProvJSON)
	dataJSON := new(raw.DataFromDataJSON)
	errs := make([]string, 0)

	//get data from udate.json
	go func(rest *raw.DataFromUpdateJSON, mutex *sync.Mutex, wg *sync.WaitGroup) {
		updateJSON, err := requestUpdateJSON()
		if err != nil {
			mutex.Lock()
			errs = append(errs, err.Error())
			mutex.Unlock()
			wg.Done()

			return
		}

		*rest = *updateJSON

		wg.Done()
	}(updateJSON, mutex, wg)

	//get data from prov.json
	go func(rest *raw.DataFromProvJSON, mutex *sync.Mutex, wg *sync.WaitGroup) {
		provJSON, err := requestProvJSON()
		if err != nil {
			mutex.Lock()
			errs = append(errs, err.Error())
			mutex.Unlock()
			wg.Done()

			return
		}

		*rest = *provJSON

		wg.Done()
	}(provJSON, mutex, wg)

	//get data from data.json
	go func(rest *raw.DataFromDataJSON, mutex *sync.Mutex, wg *sync.WaitGroup) {
		dataJSON, err := requestDataJSON()
		if err != nil {
			mutex.Lock()
			errs = append(errs, err.Error())
			mutex.Unlock()
			wg.Done()

			return
		}

		*rest = *dataJSON

		wg.Done()
	}(dataJSON, mutex, wg)

	wg.Wait()

	covidData := new(COVIDData)
	fromUpdateJSON(covidData, updateJSON)
	fromProvJSON(covidData, provJSON)
	fromDataJSON(covidData, dataJSON)

	return covidData, nil
}

func GetAllDataByProvince() ([]ByProvince, error) {
	provJSON, err := requestProvJSON()
	if err != nil {
		return nil, err
	}

	covidData := new(COVIDData)
	fromProvJSON(covidData, provJSON)

	return covidData.ByProvince, nil
}

func GetNationalSummary() (*Item, error) {
	updateJSON, err := requestUpdateJSON()
	if err != nil {
		return nil, err
	}

	covidData := new(COVIDData)
	fromUpdateJSON(covidData, updateJSON)

	return covidData.NationalSummary, nil
}
