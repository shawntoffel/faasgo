package faasgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func (g Gateway) request(method string, endpoint string, data interface{}, output interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	if g.user != "" || g.pass != "" {
		req.SetBasicAuth(g.user, g.pass)
	}

	req.Header.Add("Content-Type", "application/json")

	return g.doRequest(req, output)
}

func (g Gateway) doRequest(req *http.Request, output interface{}) error {
	resp, err := g.Client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = checkErrors(resp)
	if err != nil {
		return err
	}

	return decodeJson(resp.Body, output)
}

func (g Gateway) simpleRequest(method string, endpoint string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := g.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = checkErrors(resp)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func checkErrors(response *http.Response) error {
	if response.StatusCode >= 200 && response.StatusCode <= 299 {
		return nil
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return fmt.Errorf("%d: %s", response.StatusCode, string(body))
}

func decodeJson(body io.Reader, into interface{}) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	if b == nil || len(b) < 1 {
		return nil
	}

	return json.Unmarshal(b, &into)
}
