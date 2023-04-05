package controllers

import (
	"Combine-Gorm-Mux-Jwt/models"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

type Person struct {
	Id    int
	Name  string
	Phone string
}

func ReadPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := models.ConnectDB()
	if err != nil {
		http.Error(w, "Failed connect to database", http.StatusInternalServerError)
		return
	}

	var read []models.Person
	if err = db.Find(&read).Error; err != nil {
		return
	}

	response, err := json.Marshal(read)
	if err != nil {
		panic(err)
	}

	w.Write(response)
}

func ReadPersonById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := models.ConnectDB()
	if err != nil {
		http.Error(w, "failed connect to database", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Id not found", http.StatusBadRequest)
		return
	}

	var person models.Person
	if err = db.First(&person, id).Error; err != nil {
		http.Error(w, "Data person not found", http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(person)
	if err != nil {
		http.Error(w, "Data can't convert to JSON", http.StatusBadRequest)
		return
	}

	w.Write(response)
}

func AddPerson(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := models.ConnectDB()
	if err != nil {
		http.Error(w, "failed connect to database", http.StatusInternalServerError)
		return
	}

	var person models.Person
	err = json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, "failed decode to JSON", http.StatusBadRequest)
		return
	}
if err = db.Create(&person).Error; err != nil {
	http.Error(w, "failed add data", http.StatusBadRequest)
	return
}

w.Write([]byte("Success add new data"))


}


func EditPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := models.ConnectDB()
	if err != nil {
		http.Error(w, "failed connect to database", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "cannot find id person", http.StatusBadRequest)
		return
	}

	var person models.Person
	err = json.NewDecoder(r.Body).Decode(&person)
	if db.Where("id = ?", id).Updates(&person).RowsAffected == 0 {
		http.Error(w, "cannot update data", http.StatusBadRequest)
		return
	}

	w.Write([]byte("Update Data Success"))
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appliaction/json")

	vars := mux.Vars(r)
	id := vars["id"]

	db, err := models.ConnectDB()
	if err != nil {
		http.Error(w, "failed connect to database", http.StatusInternalServerError)
		return
	}

	var person models.Person

	if db.Delete(&person, "id = ?", id).RowsAffected == 0 {
		http.Error(w, "Delete Data Failed", http.StatusBadRequest)
		return
	}

	w.Write([]byte("Delete Data Success"))
}
