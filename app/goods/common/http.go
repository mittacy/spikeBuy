package common

import (
	"bytes"
	"encoding/json"
	"goods/app/model"
	"io/ioutil"
	"net/http"
)

func PostRequest(url string, jsonData []byte) (*model.Result, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var r model.Result
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
