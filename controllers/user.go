package controllers

import (
	"encoding/json"
	"gbu-go-postgresql/config"
	"gbu-go-postgresql/models"
	"net/http"

	"github.com/gorilla/mux"
)

//utils untuk response JSON
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

//create user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	if user.Role == "" {
		user.Role = "user" // default value
	}

	//query insert data user ke database
	sqlStatement := `INSERT INTO users (name, email, role) VALUES ($1, $2, $3) RETURNING id`
	err := config.DB.QueryRow(sqlStatement, user.Name, user.Email, user.Role).Scan(&user.ID)
	
	//response err dan success
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "User created successfully",
		"id":      user.ID,
	})
}

//get all user
func GetUsers(w http.ResponseWriter, r *http.Request) {
	//query select data user dari database
	rows, err := config.DB.Query("SELECT id, name, email, role, created_at FROM users")
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	//iterasi data user dan masukkan ke slice users agar menampilkan semua data user
	var users []models.User
	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Name, &u.Email, &u.Role, &u.CreatedAt)
		users = append(users, u)
	}

	respondJSON(w, http.StatusOK, users)
}

//get one user by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	//params by id
	params := mux.Vars(r)
	id := params["id"]

	//query select data user by id dari database
	var u models.User
	sqlStatement := `SELECT id, name, email, role, created_at FROM users WHERE id=$1`
	err := config.DB.QueryRow(sqlStatement, id).Scan(&u.ID, &u.Name, &u.Email, &u.Role, &u.CreatedAt)

	if err != nil {
		respondJSON(w, http.StatusNotFound, map[string]string{"error": "User not found"})
		return
	}

	respondJSON(w, http.StatusOK, u)
}

//update user by id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	//query update data user by id ke database
	sqlStatement := `UPDATE users SET name=$1, role=$2 WHERE id=$3`
	res, err := config.DB.Exec(sqlStatement, user.Name, user.Role, id)

	if err != nil {
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		respondJSON(w, http.StatusNotFound, map[string]string{"error": "User not found"})
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "User updated successfully"})
}

//delete user by id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	//query delete data user by id ke database
	sqlStatement := `DELETE FROM users WHERE id=$1`
	res, err := config.DB.Exec(sqlStatement, id)

	if err != nil {
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		respondJSON(w, http.StatusNotFound, map[string]string{"error": "User not found"})
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "User deleted successfully"})
}