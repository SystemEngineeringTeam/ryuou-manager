package main

import (
	"net/http"

	"github.com/SystemEngineeringTeam/ryuou-manager/handler"
)

func main() {
	http.HandleFunc("/questions/", handler.QuestionHandler)

	http.ListenAndServe(":8080", nil)
}
