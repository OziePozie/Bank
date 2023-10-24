package handlers

import (
	"Bank/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (h handler) RegisterAccount(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var account models.Account
	json.Unmarshal(body, &account)

	if result := h.DB.Create(&account); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
