package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	case http.MethodGet:
		sendAllTeamsHandler(w, r)
	case http.MethodPost:
		createTeamHandler(w, r)
	}
}

func AdminTeamRemoveHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	vars := mux.Vars(r)
	teamID := vars["team_id"]

	numericTeamID, err := strconv.Atoi(teamID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !dboperation.TeamExists(numericTeamID) {
		http.Error(w, "Team does not exist", http.StatusBadRequest)
		return
	}

	if err := dboperation.RemoveTeam(numericTeamID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func sendAllTeamsHandler(w http.ResponseWriter, r *http.Request) {
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

func JoinTeamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	// admin/teams/{team_id}/{user_id}
	vars := mux.Vars(r)
	teamID := vars["team_id"]
	userID := vars["user_id"]

	log.Println("hogehoge")

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

func LeftTeamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	// admin/teams/{team_id}/{user_id}
	vars := mux.Vars(r)
	userID := vars["user_id"]

	numericUserID, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !dboperation.UserExists(numericUserID) {
		http.Error(w, "Team does not exist", http.StatusBadRequest)
		return
	}

	if err := dboperation.LeaveTeam(numericUserID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createTeamHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var t struct {
		Name string `json:"name"`
	}
	err = json.Unmarshal(b, &t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dboperation.InsertNewTeam(t.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
