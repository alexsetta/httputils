package httputils

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	client = &http.Client{
		Timeout: time.Second * 15,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
)

func GetBody(url string, debug bool) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("GetBody: NewRequest: %v", err)
	}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("GetBody: Do: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("GetBody: status error: %s", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("GetBody: ReadAll: %v", err)
	}
	data := string(body)
	if debug {
		fmt.Println(data)
	}
	return data, nil
}
