package main

import (
	"Bank/internal/db"
	"Bank/internal/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()
	router.HandleFunc("/register", h.RegisterAccount).Methods(http.MethodPost)
	//router.HandleFunc("/accounts", h.GetAllAccounts).Methods(http.MethodGet)
	//router.HandleFunc("/accounts/{id}", h.GetAccount).Methods(http.MethodGet)
	//
	//router.HandleFunc("/accounts/{id}", h.UpdateAccount).Methods(http.MethodPut)
	//router.HandleFunc("/accounts/{id}", h.DeleteAccount).Methods(http.MethodDelete)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)

}
