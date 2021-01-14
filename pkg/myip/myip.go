package myip

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const (
	gcURLJSONIP   = "https://jsonip.com"
	gcURLIFConfig = "https://ifconfig.me/ip"

	gcTimeout = time.Second * 5
)

// GetByJSONIP by https://jsonip.com
func GetByJSONIP() (string, error) {
	client := http.Client{
		Timeout: gcTimeout,
	}

	resp, err := client.Get(gcURLJSONIP)
	if err != nil {
		return "", errors.Wrap(err, "get ip by "+gcURLJSONIP)
	}

	defer resp.Body.Close()
	jsonBody := json.NewDecoder(resp.Body)

	var mapBody map[string]string
	err = jsonBody.Decode(&mapBody)
	if err != nil {
		return "", errors.Wrap(err, "parse json from "+gcURLJSONIP)
	}

	ip, ok := mapBody["ip"]
	if !ok {
		return "", errors.Wrap(err, "not found ip from "+gcURLJSONIP)
	}

	return ip, nil
}

// GetByIFConfig by https://ifconfig.me/ip
func GetByIFConfig() (string, error) {
	client := http.Client{
		Timeout: gcTimeout,
	}

	resp, err := client.Get(gcURLIFConfig)
	if err != nil {
		return "", errors.Wrap(err, "get ip by "+gcURLIFConfig)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "read body from "+gcURLIFConfig)
	}

	return string(b), nil
}
