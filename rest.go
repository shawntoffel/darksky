package darksky

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func Get(url string, headers map[string]string, output interface{}) error {

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)

	defer response.Body.Close()

	return DecodeJson(response, &output)
}

func DecodeJson(r *http.Response, into interface{}) error {
	var decoder = json.NewDecoder(r.Body)

	if r.StatusCode != 200 {
		body, _ := ioutil.ReadAll(r.Body)
		return errors.New("Bad response: " + string(body))
	}

	return decoder.Decode(&into)
}
