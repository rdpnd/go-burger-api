package routes

import (
	"burger-api/internal/model"
	"encoding/json"
	"github.com/gorilla/pat"
	"net/http"
)

func CreateHandlers() *pat.Router {
	mux := pat.New()

	mux.Get("/api/v1/burgers/random", makeJsonHandler(viewRandomHandler))
	mux.Get("/api/v1/burgers/{id}", makeJsonHandler(viewOneHandler))
	mux.Get("/api/v1/burgers", makeJsonHandler(viewAllHandler))
	mux.Post("/api/v1/burgers", makeJsonHandler(saveOneHandler))

	return mux
}

func makeJsonHandler(fn func(http.ResponseWriter, *http.Request) interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := fn(w, r)
		if result == nil {
			return
		}
		js, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(js)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

func viewAllHandler(w http.ResponseWriter, r *http.Request) interface{} {
	p := model.ExtractPage(r)
	nameEq := r.URL.Query().Get("burger_name")
	burgers, err := model.FindAll(p, nameEq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	return burgers
}

func viewRandomHandler(w http.ResponseWriter, r *http.Request) interface{} {
	burger, err := model.FindRandom()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	return burger
}

func viewOneHandler(w http.ResponseWriter, r *http.Request) interface{} {
	id := r.URL.Query().Get(":id")
	burgers, err := model.FindOne(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	return burgers
}

func saveOneHandler(w http.ResponseWriter, r *http.Request) interface{} {
	var reqBurger model.Burger
	err := json.NewDecoder(r.Body).Decode(&reqBurger)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	insrBurger, err := model.InsertOne(&reqBurger)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	return insrBurger
}
