package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/SystemEngineeringTeam/ryuou-manager/dboperation"
)

func QuestionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// net/httpにはpathでパラメータを処理するリッチな機能が備わっていないため，
		// 以下のように分割して処理する
		// /questions/{team_id}←len == 3
		// /questions/{team_id}/{question_id}←len == 4
		if len(strings.Split(r.URL.Path, "/")) == 3 {
			getQuestions(w, r)
		} else if len(strings.Split(r.URL.Path, "/")) == 4 {
			getQuestionDetail(w, r)
		}

	case http.MethodPut:

	}
}

func getQuestions(w http.ResponseWriter, r *http.Request) {

	// Get the team id from the url path
	// The request path is '/questions/{team_id}'
	teamID := strings.Split(r.URL.Path, "/")[2]

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

func getQuestionDetail(w http.ResponseWriter, r *http.Request) {
	teamID := strings.Split(r.URL.Path, "/")[2]
	questionID := strings.Split(r.URL.Path, "/")[3]

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
