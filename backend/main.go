package main

import (
	"net/http"

	"github.com/SystemEngineeringTeam/ryuou-manager/handler"
)

func main() {
	http.HandleFunc("/questions/", handler.QuestionHandler)
	http.HandleFunc("/admin/questions/", handler.AdminQuestionHandler)
	http.HandleFunc("/admin/teams/", handler.AdminTeamHandler)
	http.HandleFunc("/users", handler.UserHandler)
	http.HandleFunc("/users/login", handler.LoginHandler)

	http.ListenAndServe(":8080", nil)
}
