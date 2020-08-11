package druid

import (
	"bytes"
	"net/http"

	"github.com/gerfg/go-druid/utils"
)

func Query(reqJSON *[]byte) (*[]byte, error) {
	req, err := http.NewRequest("POST", "http://localhost:8888/druid/v2/", bytes.NewBuffer(*reqJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := utils.ExtractResponse(resp)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
