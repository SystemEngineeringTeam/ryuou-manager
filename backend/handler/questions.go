package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/SystemEngineeringTeam/ryuou-manager/dboperation"
	"github.com/SystemEngineeringTeam/ryuou-manager/model"
	"github.com/gorilla/mux"
)

// 問題全件取得
func SendAllQuestionHandler(w http.ResponseWriter, r *http.Request) {

	// allow cors
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

func QuestionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		questionDetailHandler(w, r)
	case http.MethodPost:
		submitQuestionHandler(w, r)
	case http.MethodPut:
		openQuestionHandler(w, r)
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		// allow cors
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
	}
}

// 問題詳細を取得するためのハンドラ
func questionDetailHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID := vars["team_id"]
	questionID := vars["question_id"]

	numericTeamID, err := strconv.Atoi(teamID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	numericQuestionID, err := strconv.Atoi(questionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !dboperation.QuestionExists(numericQuestionID) || !dboperation.TeamExists(numericTeamID) {
		http.Error(w, "Invalid TeamID or QuestionID", http.StatusBadRequest)
		return
	}

	question, err := dboperation.SelectQuestionDetailByTeamID(numericTeamID, numericQuestionID)
	if err != nil {
		http.Error(w, "Question Not Opened", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(question)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(response))
}

// 問題提出用のハンドラ
func submitQuestionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID := vars["team_id"]
	questionID := vars["question_id"]

	numericTeamID, err := strconv.Atoi(teamID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	numericQuestionID, err := strconv.Atoi(questionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !dboperation.QuestionExists(numericQuestionID) || !dboperation.TeamExists(numericTeamID) {
		http.Error(w, "Invalid TeamID or QuestionID", http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var answer model.SubmitRequest
	err = json.Unmarshal(body, &answer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dboperation.SubmitQuestion(numericTeamID, numericQuestionID, answer.Answer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// 問題をオープンするためのハンドラ
func openQuestionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID := vars["team_id"]
	questionID := vars["question_id"]

	numericTeamID, err := strconv.Atoi(teamID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	numericQuestionID, err := strconv.Atoi(questionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !dboperation.QuestionExists(numericQuestionID) || !dboperation.TeamExists(numericTeamID) {
		http.Error(w, "Invalid TeamID or QuestionID", http.StatusBadRequest)
		return
	}

	err = dboperation.OpenQuestion(numericTeamID, numericQuestionID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
