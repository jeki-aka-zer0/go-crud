package controllers

import (
	"github.com/jeki-aka-zer0/go-crud/models"

	"encoding/json"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	models.Test()

	res, _ := json.Marshal([]string{"Hello", "world"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
