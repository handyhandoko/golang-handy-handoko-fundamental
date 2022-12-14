package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"project-2/model"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	
	"github.com/go-playground/validator/v10"
)

const (
	username = "root"
	password = ""
	hostname = "127.0.0.1:3306"
	dbname = "belajargolang"
)

var validate *validator.Validate
var db *sql.DB

type Server struct {
	Router *chi.Mux
	// Db, config can be added here
}

// return Data Source Name string
func dsn () string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}

func pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong!"))
}

func list(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("select id, name, price from vegetables")
	if (err != nil) {
		fmt.Printf("Error select: %s", err.Error())
	}
	defer rows.Close()

	var vegetables []model.Vegetable
	for rows.Next(){
		var vegetable model.Vegetable
		err = rows.Scan(&vegetable.Id, &vegetable.Name, &vegetable.Price)

		if (err != nil) {
			fmt.Printf("Error scan: %s", err.Error())
		}
		vegetables = append(vegetables, vegetable)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vegetables)
}

func get(w http.ResponseWriter, r *http.Request) {
  id := chi.URLParam(r, "id")
	var vegetable model.Vegetable
	
	err := db.
				QueryRow("SELECT id, name, price FROM vegetables WHERE id = ?", id).
				Scan(&vegetable.Id, &vegetable.Name, &vegetable.Price)
	if (err != nil) {
		fmt.Printf("Error query row: %s", err.Error())
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
	_, err = db.Exec("INSERT INTO vegetables (name, price) VALUES (?, ?)", newVegetable.Name, newVegetable.Price)
	if (err != nil) {
		fmt.Printf("Error Insert Row: %s", err.Error())
	}
}

func update(w http.ResponseWriter, r *http.Request) {
	var updateVegetable model.Vegetable
	err := json.NewDecoder(r.Body).Decode(&updateVegetable)
	if err != nil {
		panic(err)
	}
	err = validate.Struct(updateVegetable)
	if err != nil {
		http.Error(w, err.Error(), 422)
	}

  id := chi.URLParam(r, "id")

	_, err = db.Exec("UPDATE vegetables SET name = ?, price = ? WHERE id = ?", updateVegetable.Name, updateVegetable.Price, id)
	if (err != nil) {
		fmt.Printf("Error Update Data: %s", err.Error())
	}
}

func delete(w http.ResponseWriter, r *http.Request) {
  id := chi.URLParam(r, "id")
	_, err := db.Exec("DELETE FROM vegetables WHERE id = ?", id)
	if (err != nil){
		fmt.Printf("Error Delete Data: %s", err.Error())
	}
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

	var err error
	db, err = sql.Open("mysql", dsn())
	if err != nil {
		log.Printf("Error %s when opening db.", err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if(err != nil){
		fmt.Printf("error connecting to database: %s \n",err.Error())
	}

	s := CreateNewServer()
	s.MountHandlers()
	http.ListenAndServe(":3000", s.Router)
}