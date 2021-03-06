package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/SystemEngineeringTeam/ryuou-manager/dboperation"
	"github.com/SystemEngineeringTeam/ryuou-manager/model"
	"github.com/gorilla/mux"
)

func CollectHandler(w http.ResponseWriter, r *http.Request) {
	// allow cors
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	vars := mux.Vars(r)
	teamID := vars["team_id"]
	questionID := vars["question_id"]

	numericTeamID, err := strconv.Atoi(teamID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	numericQuestionID, err := strconv.Atoi(questionID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := dboperation.ChangeToCollect(numericTeamID, numericQuestionID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := dboperation.AddScore(numericQuestionID, numericTeamID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func AdminQuestionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	switch r.Method {
	case http.MethodGet:
		sendAllSubmitsHandler(w, r)
	case http.MethodPost:
		createNewQuestionHandler(w, r)
	}
}

func sendAllSubmitsHandler(w http.ResponseWriter, r *http.Request) {
	submits, err := dboperation.SelectAllSubmit()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(submits)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(jsonBytes))
}

func createNewQuestionHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var q model.QuestionCreateRequest

	if err := json.Unmarshal(b, &q); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	var question model.Question
	question.Title = q.Title
	question.Description = q.Description
	question.Score = q.Score

	if err := dboperation.InsertNewQuestion(question); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func AdminQuestionIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	switch r.Method {
	case http.MethodDelete:
		deleteQuestionHandler(w, r)
	}
}

func deleteQuestionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	questionID := vars["question_id"]

	numericQuestionID, err := strconv.Atoi(questionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := dboperation.DeleteQuestion(numericQuestionID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
