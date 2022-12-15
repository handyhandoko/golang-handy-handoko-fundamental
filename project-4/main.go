package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"strconv"

	"project-2/repository"
	"project-2/model"
	"project-2/sort"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	
	"github.com/go-playground/validator/v10"
)

var vegetables = []model.Vegetable{
	model.Vegetable {1, "sawi", 4500},
	model.Vegetable {3, "bayam", 2000},
	model.Vegetable {2, "kangkung", 1000},
	model.Vegetable {4, "kol", 5000},
	model.Vegetable {5, "pare", 3000},
}

var validate *validator.Validate

type Server struct {
	Router *chi.Mux
	// Db, config can be added here
}

func pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong!"))
}

func list(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vegetables)
}

func get(w http.ResponseWriter, r *http.Request) {
  idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	index, err := repository.FindIndexById(vegetables, uint(id))
	var vegetable model.Vegetable = vegetables[index]

  if err != nil {
    w.WriteHeader(500)
    w.Write([]byte(fmt.Sprintf("error fetching data %d: %v", idParam, err)))
    return
  }

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vegetable)
}

func create(w http.ResponseWriter, r *http.Request) {

	var newVegetable model.Vegetable
	var err error = json.NewDecoder(r.Body).Decode(&newVegetable)
	if err != nil {
		panic(err)
	}
	err = validate.Struct(newVegetable)
	if err != nil {
		http.Error(w, err.Error(), 422)
	}
	vegetables = repository.AddItem(vegetables, newVegetable)
}

func update(w http.ResponseWriter, r *http.Request) {
  idParam := chi.URLParam(r, "id")
	id, _ := strconv.ParseUint(idParam, 10, 32)

	var updateVegetable model.Vegetable
	err := json.NewDecoder(r.Body).Decode(&updateVegetable)
	if err != nil {
		panic(err)
	}
	err = validate.Struct(updateVegetable)
	if err != nil {
		http.Error(w, err.Error(), 422)
	}
	vegetables, _ = repository.UpdateById(vegetables, updateVegetable, uint(id))
}

func delete(w http.ResponseWriter, r *http.Request) {
  idParam := chi.URLParam(r, "id")
	id, _ := strconv.ParseUint(idParam, 10, 32)

	index, err := repository.FindIndexById(vegetables, uint(id))
	if err != nil {
		http.Error(w, err.Error(), 404)
	}
	var vegetable model.Vegetable = vegetables[index]

	vegetables, _ = repository.RemoveById(vegetables, vegetable)
}

func (s *Server) MountHandlers() {
	// Mount all Middleware here
	s.Router.Use(middleware.Logger)

	// Mount all handlers here
	s.Router.Get("/ping", pong)
	s.Router.Get("/vegetable", list)
	s.Router.Get("/vegetable/{id}", get)
	s.Router.Post("/vegetable", create)
	s.Router.Put("/vegetable/{id}", update)
	s.Router.Delete("/vegetable/{id}", delete)
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

func main() {
	validate = validator.New()
	vegetables = sort.Selection(vegetables)

	s := CreateNewServer()
	s.MountHandlers()
	http.ListenAndServe(":3000", s.Router)
}