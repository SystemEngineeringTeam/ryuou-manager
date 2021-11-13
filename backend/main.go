package main

import (
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/SystemEngineeringTeam/ryuou-manager/handler"
)

func main() {
	r := mux.NewRouter()

	// /questions
	r.HandleFunc("/questions/{team_id:[0-9]+}", nil)
	r.HandleFunc("/questions/{team_id:[0-9]+}/{question_id:[0-9]+}", nil)

	// /users
	r.HandleFunc("/users", nil)
	r.HandleFunc("/users/login", nil)

	// /admin/questions
	r.HandleFunc("/admin/questions", nil)
	r.HandleFunc("/admin/questions/{question_id:[0-9]+}", nil)
	r.HandleFunc("/admin/questions/{team_id:[0-9]+}/{question_id:[0-9]+}", nil)

	// /admin/teams
	r.HandleFunc("/admin/teams", nil)
	r.HandleFunc("/admin/teams/{team_id:[0-9]+}/{user_id:[0-9]+}", nil)

	http.ListenAndServe(":8080", nil)
}
