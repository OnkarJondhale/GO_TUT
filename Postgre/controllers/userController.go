package controller

import (
	auth "Blog/middlewares"
	model "Blog/models"
	"Blog/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting user details...")

	qparams := mux.Vars(r)
	id := qparams["userId"]

	query := ` SELECT id,fname,lname,email
				FROM users
				WHERE id = $1`

	var user model.Users
	_ = utils.DB.QueryRow(query, id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email)

	fmt.Println("User details fetched successfully")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "user details fetched successfully",
		"data":    user,
	})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting user details for deleting it...")

	payload, _ := r.Context().Value(auth.Payload).(jwt.MapClaims)
	id := payload["id"]

	query := `DELETE FROM users
				WHERE id = $1
				RETURNING id,fname,lname,email`

	var user model.Users
	_ = utils.DB.QueryRow(query, id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email)

	fmt.Println("User details deleted successfully")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "user deleted successfully",
		"data":    user,
	})
}

func UpdateUserDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting user details for updating it...")

	payload, _ := r.Context().Value(auth.Payload).(jwt.MapClaims)
	id := payload["id"]

	var user model.Users
	_ = json.NewDecoder(r.Body).Decode(&user)

	query := `UPDATE users
				SET fname = $1,lname = $2
				WHERE id = $3
				RETURNING id,fname,lname,email`

	var response model.Users
	_ = utils.DB.QueryRow(query, user.FirstName, user.LastName, id).Scan(&response.Id, &response.FirstName, &response.LastName, &response.Email)

	fmt.Println("user details deleted successfully")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "user deleted successfully",
		"data":    response,
	})
}

func UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting user details for updating it...")

	payload, _ := r.Context().Value(auth.Payload).(jwt.MapClaims)
	id := payload["id"]

	type Response struct {
		OldPassword string `json:"oldPassword"`
		NewPassword string `json:"newPassword"`
		Password    string `json:"password"`
	}

	var user Response
	_ = json.NewDecoder(r.Body).Decode(&user)

	_ = utils.DB.QueryRow(`SELECT password from users WHERE id = $1`, id).Scan(&user.Password)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.OldPassword))
	if err != nil {
		fmt.Println("Error in updating password")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "Password is does not match with previous password",
		})
		return
	}

	query := `UPDATE users
				SET password = $1
				WHERE id = $2
				RETURNING id,fname,lname,email`

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.NewPassword), bcrypt.DefaultCost)

	var response model.Users
	_ = utils.DB.QueryRow(query, hashedPassword, id).Scan(&response.Id, &response.FirstName, &response.LastName, &response.Email)

	fmt.Println("user password updated successfully")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "user updated successfully",
		"data":    response,
	})
}
