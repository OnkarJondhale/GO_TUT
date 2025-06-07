package controller

import (
	model "Blog/models"
	"Blog/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("User registration initiated...")

	// Get data from request body
	var user model.Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Error in register")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "Invalid credentials",
		})
		return
	}

	// check credentials
	if user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Password == "" {
		fmt.Println("Error in register")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "All fields are required",
		})
		return
	}

	fmt.Println("Received data is ", user)

	// check if user exists
	var count int
	_ = utils.DB.QueryRow(`SELECT COUNT(*) FROM users WHERE email = $1`, user.Email).Scan(&count)

	if count > 0 {
		fmt.Println("Error in register")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "Account already exist",
		})
		return
	}

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error in register")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "Error while hashing password",
		})
		return
	}

	// create the entry for the user
	var res model.Users
	_ = utils.DB.QueryRow(`INSERT INTO users(fname,lname,email,password) values($1,$2,$3,$4) RETURNING id,fname,lname,email,createdAt`, user.FirstName, user.LastName, user.Email, hashedPassword).Scan(&res.Id, &res.FirstName, &res.LastName, &res.Email, &res.CreatedAt)

	fmt.Println("User registrered successfully...", res)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "User registered successfully",
		"data":    res,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("User login initiated...")

	// get the data from request body
	var data model.Users
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error in login")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "Invalid credentials",
		})
		return
	}

	fmt.Println("Received data is ", data)

	// check the credentials
	if data.Email == "" || data.Password == "" {
		fmt.Println("Error in login")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "All fields are required",
		})
		return
	}

	// check user exists or not if it is then get its data
	var user model.Users
	err = utils.DB.QueryRow(`SELECT id, fname, lname, email, password FROM users WHERE email = $1`, data.Email).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		fmt.Println("Error in login")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "Account does not exist",
		})
		return
	}

	// compare the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		fmt.Println("Error in login")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "Password is incorrect",
		})
		return
	}

	payload := jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	_ = godotenv.Load(".env")
	jwtSecrete := os.Getenv("JWT_SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	jwtToken, _ := token.SignedString([]byte(jwtSecrete));

	fmt.Println("USer login successful")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "User login successful",
		"data":    jwtToken,
	})
}
