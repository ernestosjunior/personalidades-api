package personalities

import (
	"encoding/json"
	"net/http"
	"personalidades-api/database"
	"personalidades-api/models"

	"github.com/gorilla/mux"
)

func List(w http.ResponseWriter, r *http.Request) {
	var personalities []models.Personality
	database.DB.Find(&personalities)
	json.NewEncoder(w).Encode(personalities)
}

func Find(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality

	database.DB.First(&personality, id)

	json.NewEncoder(w).Encode(personality)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	json.NewDecoder(r.Body).Decode(&personality)

	database.DB.Create(&personality)

	json.NewEncoder(w).Encode(personality)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	database.DB.Delete(&models.Personality{}, id)

	w.WriteHeader(http.StatusNoContent)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality

	database.DB.First(&personality, id)

	json.NewDecoder(r.Body).Decode(&personality)

	database.DB.Save(&personality)

	json.NewEncoder(w).Encode(personality)
}
