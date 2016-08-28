package dominos

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RequestPayload struct {
	Email             string `json:"Email"`
	EmailOptIn        bool   `json:"EmailOptIn"`
	ReasonDescription string `json:"ReasonDescription"`
}

type ResponsePayload struct {
	Status      int                 `json:"Status"`
	Email       string              `json:"Email"`
	Type        string              `json:"Type"`
	StatusItems []map[string]string `json:"StatusItems"`
}

func Unsub(email string) error {
	// JSON payload
	payload := RequestPayload{
		email,
		false,
		"0005",
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Formulate request based on tests
	req, err := http.NewRequest(
		"POST",
		"https://us-external-api.dominos.com/power/opt-in-and-opt-out",
		bytes.NewReader(data),
	)
	if err != nil {
		return err
	}
	req.Header.Add(
		"referer",
		"https://us-external-api.dominos.com/en/assets/build/xdomain/proxy.html",
	)
	req.Header.Add(
		"origin",
		"https://us-external-api.dominos.com",
	)
	req.Header.Add(
		"content-type",
		"text/plain;charset=UTF-8",
	)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// Verify response
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	respPayload := ResponsePayload{}
	if err = json.Unmarshal(respBytes, &respPayload); err != nil {
		return err
	}

	return nil
}
