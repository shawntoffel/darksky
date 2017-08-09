package darksky

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

func get(url string, output interface{}) error {

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept-Encoding", "gzip")

	client := http.Client{}
	response, err := client.Do(req)

	defer response.Body.Close()

	err = checkErrors(response)

	if err != nil {
		return err
	}

	return decodeCompressedJson(response.Body, &output)
}

func checkErrors(response *http.Response) error {
	if response.StatusCode != 200 {
		body, _ := ioutil.ReadAll(response.Body)
		return errors.New("Bad response: " + string(body))
	}

	return nil
}

func decodeCompressedJson(body io.Reader, into interface{}) error {
	reader, err := gzip.NewReader(body)

	if err != nil {
		return err
	}

	jsonDecoder := json.NewDecoder(reader)

	return jsonDecoder.Decode(&into)
}
