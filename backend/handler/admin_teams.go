package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SystemEngineeringTeam/ryuou-manager/dboperation"
)

func AdminTeamHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAdminTeams(w, r)
	}
}

func getAdminTeams(w http.ResponseWriter, r *http.Request) {
	teams := dboperation.SelectAllTeams()

	jsonTeams, err := json.Marshal(teams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonTeams))
}
