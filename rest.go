package darksky

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type RestClient interface {
	Get(url string, headers map[string]string, output interface{}) error
}

type restClient struct {
	Client *http.Client
}

func NewRestClient() RestClient {
	client := &http.Client{}

	return &restClient{client}
}

func NewRestClientWithClient(client *http.Client) RestClient {
	return &restClient{client}
}

func (r *restClient) Get(url string, headers map[string]string, output interface{}) error {

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

	return decodeJson(response, &output)
}

func decodeJson(r *http.Response, into interface{}) error {
	var decoder = json.NewDecoder(r.Body)

	if r.StatusCode != 200 {
		body, _ := ioutil.ReadAll(r.Body)
		return errors.New("Bad response: " + string(body))
	}

	return decoder.Decode(&into)
}
