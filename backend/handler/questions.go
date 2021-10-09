package handler

import (
	"net/http"

	"github.com/SystemEngineeringTeam/ryuou-manager/dboperation"
)

func QuestionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getQuestion(w, r)
	}
}

func getQuestion(w http.ResponseWriter, r *http.Request) {
	questions, err := dboperation.SelectAllQuestions()
	if err != nil {
		w.Write([]byte("error"))
		return
	}

}
