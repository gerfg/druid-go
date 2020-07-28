package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func ExtractResponse(resp *http.Response) (*[]byte, error) {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil
	}
	return &bodyBytes, err
}

func ExtractRequestBody(req *http.Request) (*[]byte, error) {
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	return &bodyBytes, nil
}

func OpenJSON(path string, print bool) (*[]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	jsonBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	if print {
		fmt.Println(string(jsonBytes))
	}
	return &jsonBytes, nil
}
