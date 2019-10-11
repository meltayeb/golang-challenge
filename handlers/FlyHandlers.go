package handlers

import (
	"encoding/json"
	"github.com/go-task/models"
	"github.com/go-task/services"
	"github.com/gorilla/schema"
	"net/http"
)

func TransactionIndex(w http.ResponseWriter, r *http.Request) {

	filter := new(models.Filter)
	if err := schema.NewDecoder().Decode(filter, r.URL.Query()); err != nil {
		// Handle error
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

	result:= services.Filter(filter)

    if err := json.NewEncoder(w).Encode(result); err != nil {
       panic(err)
    }
}

