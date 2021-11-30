package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SystemEngineeringTeam/ryuou-manager/dboperation"
	"github.com/gorilla/mux"
)

func AdminTeamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	switch r.Method {
	case http.MethodPost:
		joinTeamHandler(w, r)
	case http.MethodDelete:
		leftTeamHandler(w, r)
	}
}

func SendAllTeamsHandler(w http.ResponseWriter, r *http.Request) {
	// allow cors
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	teams := dboperation.SelectAllTeams()

	jsonTeams, err := json.Marshal(teams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonTeams))
}

func joinTeamHandler(w http.ResponseWriter, r *http.Request) {
	// admin/teams/{team_id}/{user_id}
	vars := mux.Vars(r)
	teamID := vars["team_id"]
	userID := vars["user_id"]

	numericTeamID, err := strconv.Atoi(teamID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	numericUserID, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !dboperation.TeamExists(numericTeamID) || !dboperation.UserExists(numericUserID) {
		http.Error(w, "Team does not exist", http.StatusBadRequest)
		return
	}

	if err := dboperation.JoinTeam(numericTeamID, numericUserID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func leftTeamHandler(w http.ResponseWriter, r *http.Request) {
	// admin/teams/{team_id}/{user_id}
	vars := mux.Vars(r)
	teamID := vars["team_id"]
	userID := vars["user_id"]

	numericTeamID, err := strconv.Atoi(teamID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	numericUserID, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !dboperation.TeamExists(numericTeamID) || !dboperation.UserExists(numericUserID) {
		http.Error(w, "Team does not exist", http.StatusBadRequest)
		return
	}

	if err := dboperation.LeaveTeam(numericTeamID, numericUserID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
