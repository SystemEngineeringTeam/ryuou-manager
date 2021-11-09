package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/SystemEngineeringTeam/ryuou-manager/dboperation"
)

func AdminTeamHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAdminTeams(w, r)
	case http.MethodPost:
		postAdminTeams(w, r)
	case http.MethodDelete:
		deleteAdminTeams(w, r)
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

func postAdminTeams(w http.ResponseWriter, r *http.Request) {
	// admin/teams/{team_id}/{user_id}
	teamID := strings.Split(r.URL.Path, "/")[3]
	userID := strings.Split(r.URL.Path, "/")[4]

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

func deleteAdminTeams(w http.ResponseWriter, r *http.Request) {
	// admin/teams/{team_id}/{user_id}
	teamID := strings.Split(r.URL.Path, "/")[3]
	userID := strings.Split(r.URL.Path, "/")[4]

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
