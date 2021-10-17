package api

import (
	"encoding/json"
	"log"
	"net/http"
	
	"github.com/EmilGeorgiev/training/go-testing/userservice"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	UserRepo userservice.UserRepository
}

func (h UserHandler) Create(w http.ResponseWriter, req *http.Request) {
	var nu userservice.NewUser
	if err := json.NewDecoder(req.Body).Decode(&nu); err != nil {
		log.Println("malformed json: ", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := h.UserRepo.Create(req.Context(), nu)
	if err != nil {
		log.Println("Failed to create a new user: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("User is successfully created.")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h UserHandler) Get(w http.ResponseWriter, req *http.Request) {
	userID := chi.URLParam(req, "id")
	u, err := h.UserRepo.Get(req.Context(), userID)
	if err != nil {
		log.Println("Filed to find user by ID: ", userID)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Found the user by ID: ", userID)
	if err = json.NewEncoder(w).Encode(u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h UserHandler) Update(w http.ResponseWriter, req *http.Request) {
	var nu userservice.NewUser
	if err := json.NewDecoder(req.Body).Decode(&nu); err != nil {
		log.Println("Malformed json: ", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := chi.URLParam(req, "id")

	u, err := h.UserRepo.Update(req.Context(), userID, nu)
	if err != nil {
		log.Println("Failed to update the user: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("User is successfully updated.")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h UserHandler) Delete(w http.ResponseWriter, req *http.Request) {
	userID := chi.URLParam(req, "id")
	if err := h.UserRepo.Delete(req.Context(), userID); err != nil {
		log.Println("Filed to delete user by ID: ", userID)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("User is deleted")
	w.WriteHeader(http.StatusNoContent)
}