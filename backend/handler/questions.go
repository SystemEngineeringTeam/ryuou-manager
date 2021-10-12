package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SystemEngineeringTeam/ryuou-manager/dboperation"
)

func QuestionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getQuestion(w, r)
	}
}

func getQuestion(w http.ResponseWriter, r *http.Request) {

	// Get the team id from the url path
	// The request path is '/questions/{team_id}'
	teamID := r.URL.Path[len("/questions/"):]

	numericTeamID, err := strconv.Atoi(teamID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	questions, err := dboperation.SelectAllQuestionsByTeamID(numericTeamID)
	if err != nil {
		w.Write([]byte("error"))
		return
	}

	response, err := json.Marshal(questions)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(response))
}
