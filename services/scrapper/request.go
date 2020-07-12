package scrapper

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	userAgent  = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0"
	cookie     = "__cfduid=d04b8423becb002e1ec2da89231f589011594142566; _ga=GA1.3.1660964307.1594142570; _gid=GA1.3.993296428.1594403012"
	updateJSON = "update.json"
	provJSON   = "prov.json"
	dataJSON   = "data.json"
)

func httpHeader() http.Header {
	httpHeader := http.Header{}
	httpHeader.Set("User-Agent", userAgent)
	httpHeader.Set("Accept", "*/*")
	httpHeader.Set("Accept-Language", "en-US,en;q=0.5")
	httpHeader.Set("X-Requested-With", "XMLHttpRequest")
	httpHeader.Set("Connection", "keep-alive")
	httpHeader.Set("Cookie", cookie)
	//just in case if error occured, try add this headers
	//If-Modified-Since: Fri, 10 Jul 2020 17:47:57 GMT
	//If-None-Match: W/"5f08a9cd-c1f1"

	return httpHeader
}

func requestAPI(file string) ([]byte, error) {
	req, err := http.NewRequest("GET", "https://data.covid19.go.id/public/api/"+file, nil)
	if err != nil {
		return nil, err
	}

	req.Header = httpHeader()

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 200 {
		return nil, errors.New(string(respBody))
	}

	return respBody, nil
}

func requestUpdateJSON() (*DataFromUpdateJSON, error) {
	raw, err := requestAPI(updateJSON)
	if err != nil {
		return nil, err
	}

	result := new(DataFromUpdateJSON)
	if err = json.Unmarshal(raw, result); err != nil {
		return nil, err
	}

	return result, nil
}

func requestProvJSON() (*DataFromProvJSON, error) {
	raw, err := requestAPI(provJSON)
	if err != nil {
		return nil, err
	}

	result := new(DataFromProvJSON)
	if err = json.Unmarshal(raw, result); err != nil {
		return nil, err
	}

	return result, nil
}

func requestDataJSON() (*DataFromDataJSON, error) {
	raw, err := requestAPI(dataJSON)
	if err != nil {
		return nil, err
	}

	result := new(DataFromDataJSON)
	if err = json.Unmarshal(raw, result); err != nil {
		return nil, err
	}

	return result, nil
}
