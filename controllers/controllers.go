package controllers

import (
	"github.com/jeki-aka-zer0/go-crud/models"

	"encoding/json"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal("Welcome to simple GoLang CRUD application")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func ReadAllEngineers(w http.ResponseWriter, r *http.Request) {
	engineers := models.GetAllEngineers()
	res, _ := json.Marshal(engineers)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
