package scrapper

import "sync"

func GetAllData() (*COVIDData, []error) {
	wg := new(sync.WaitGroup)
	mutex := new(sync.Mutex)
	wg.Add(3)

	updateJSON := new(DataFromUpdateJSON)
	provJSON := new(DataFromProvJSON)
	dataJSON := new(DataFromDataJSON)
	errs := make([]string, 0)

	//get data from udate.json
	go func(rest *DataFromUpdateJSON, mutex *sync.Mutex, wg *sync.WaitGroup) {
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
	go func(rest *DataFromProvJSON, mutex *sync.Mutex, wg *sync.WaitGroup) {
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
	go func(rest *DataFromDataJSON, mutex *sync.Mutex, wg *sync.WaitGroup) {
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
