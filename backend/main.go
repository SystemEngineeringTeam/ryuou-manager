package main

import (
	"net/http"

	"github.com/SystemEngineeringTeam/ryuou-manager/handler"
	"github.com/gorilla/mux"
	// "github.com/SystemEngineeringTeam/ryuou-manager/handler"
)

func main() {
	r := mux.NewRouter()

	// /questions
	r.HandleFunc("/questions/{team_id:[0-9]+}", handler.SendAllQuestionHandler)
	r.HandleFunc("/questions/{team_id:[0-9]+}/{question_id:[0-9]+}", handler.QuestionHandler)

	// /users
	r.HandleFunc("/users", handler.UserHandler)
	r.HandleFunc("/users/login", handler.LoginHandler)

	// /admin/questions
	r.HandleFunc("/admin/questions", handler.AdminQuestionHandler)
	r.HandleFunc("/admin/questions/{question_id:[0-9]+}", handler.AdminQuestionIDHandler)
	r.HandleFunc("/admin/questions/{team_id:[0-9]+}/{question_id:[0-9]+}", handler.CollectHandler)

	// /admin/teams
	r.HandleFunc("/admin/teams", handler.AdminTeamHandler)
	r.HandleFunc("/admin/teams/{team_id:[0-9]+}", handler.AdminTeamRemoveHandler)
	r.HandleFunc("/admin/teams/{team_id:[0-9]+}/{user_id:[0-9]+}", handler.AdminTeamWithIDHandler)

	http.ListenAndServe(":8080", r)
}
