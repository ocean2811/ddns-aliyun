package myip

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/pkg/errors"
)

const (
	// DetectorNum means have DetectorNum way to get ip
	DetectorNum = 3
)

const (
	gcURLJSONIP   = "https://jsonip.com"
	gcURLIFConfig = "https://ifconfig.me/ip"
	gcURLIPIP     = "https://myip.ipip.net"

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

// GetByIPIP by https://myip.ipip.net
func GetByIPIP() (string, error) {
	client := http.Client{
		Timeout: gcTimeout,
	}

	resp, err := client.Get(gcURLIPIP)
	if err != nil {
		return "", errors.Wrap(err, "get ip by "+gcURLIPIP)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "read body from "+gcURLIPIP)
	}

	reg := regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`)

	return reg.FindAllString(string(b), -1)[0], nil
}
