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

	body, err := decompress(response)

	if err != nil {
		return err
	}

	return decodeJson(body, &output)
}

func checkErrors(response *http.Response) error {
	if response.StatusCode != 200 {
		body, _ := ioutil.ReadAll(response.Body)
		return errors.New("Bad response: " + string(body))
	}

	return nil
}

func decompress(response *http.Response) (io.Reader, error) {
	header := response.Header.Get("Content-Encoding")

	if len(header) < 1 {
		return response.Body, nil
	}

	reader, err := gzip.NewReader(response.Body)

	if err != nil {
		return nil, err
	}

	return reader, nil
}

func decodeJson(body io.Reader, into interface{}) error {
	jsonDecoder := json.NewDecoder(body)

	return jsonDecoder.Decode(&into)
}
