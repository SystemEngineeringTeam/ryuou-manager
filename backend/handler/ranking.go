package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SystemEngineeringTeam/ryuou-manager/dboperation"
)

func RankingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	ranking, err := dboperation.SelectRankingData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(ranking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(jsonBytes))
}
