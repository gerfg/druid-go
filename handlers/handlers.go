package handlers

import (
	"net/http"

	"github.com/gerfg/go-druid/druid"
	"github.com/gerfg/go-druid/utils"
)

func DruidQuery(w http.ResponseWriter, r *http.Request) {

	reqJSON, err := utils.ExtractRequestBody(r)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{ \"error\": " + err.Error()))
	}

	responseJSON, err := druid.Query(reqJSON)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{ \"error\": " + err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(*responseJSON)
}
