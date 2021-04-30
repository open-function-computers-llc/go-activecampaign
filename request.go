package activecampaign

import (
	"bytes"
	"net/http"
)

func getRequest(route string) (*http.Response, error) {
	url := apiBase + route
	log.Info("GET: " + url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	req.Header.Add("Api-Token", apiKey)

	client := &http.Client{}
	return client.Do(req)
}

func postRequest(route string, body []byte) (*http.Response, error) {
	url := apiBase + route
	log.Info("POST: "+url, "POSTBODY: "+string(body))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	req.Header.Add("Api-Token", apiKey)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}
