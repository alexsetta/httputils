package httputils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetBody(url string, debug bool) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("GetBody: %w", err)
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("GetBody: %w", err)
	}
	data := string(b)
	if debug {
		fmt.Println(data)
	}
	return data, nil
}
