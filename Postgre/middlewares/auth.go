package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	// "github.com/golang-jwt/jwt"
)

// it creates a distinct data type as contextKey which behaves similar to string it is essential to do to avoid collisions
type contextKey string

const Payload contextKey = "claims"

// func Auth(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("Auth middleware...")

// 		// Retrive the jwt token string from the request
// 		var token map[string]any
// 		err := json.NewDecoder(r.Body).Decode(&token)
// 		if err != nil {
// 			fmt.Println("Error in auth middleware")
// 			w.WriteHeader(http.StatusBadRequest)
// 			json.NewEncoder(w).Encode(map[string]any{
// 				"success": false,
// 				"message": "Invalid credentials",
// 			})
// 			return
// 		}

// 		// load the secret key from environment variables
// 		_ = godotenv.Load(".env")
// 		jwtSecrete := os.Getenv("JWT_SECRET_KEY")

// 		// parse the jwt token to get the jwt token data such as headers,claims,row,methods,valid
// 		jwtToken, _ := token["token"].(string)
// 		data, err := jwt.Parse(jwtToken, func(token *jwt.Token) (any, error) {
// 			return []byte(jwtSecrete), nil
// 		})
// 		if err != nil || !data.Valid {
// 			fmt.Println("Error in auth middleware")
// 			w.WriteHeader(http.StatusBadRequest)
// 			json.NewEncoder(w).Encode(map[string]any{
// 				"success": false,
// 				"message": "Invalid jwt token",
// 			})
// 			return
// 		}

// 		// retrieve the Payload
// 		claim := data.Claims.(jwt.MapClaims)

// 		// add Payload to the context to send it to next handler
// 		ctx := context.WithValue(r.Context(), Payload, claim)

// 		fmt.Println("Auth validation successful...")

// 		// call the next handler
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Auth middleware...")

		// Read and restore the body
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading request body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// restore the body for next handler
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Parse token from request body
		var token map[string]any
		if err := json.Unmarshal(bodyBytes, &token); err != nil {
			fmt.Println("Error decoding request body:", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"success": false,
				"message": "Invalid credentials",
			})
			return
		}

		jwtToken, ok := token["token"].(string)
		if !ok || jwtToken == "" {
			fmt.Println("Error: Missing JWT token")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Parse JWT
		jwtSecret := os.Getenv("JWT_SECRET_KEY")
		data, err := jwt.Parse(jwtToken, func(token *jwt.Token) (any, error) {
			return []byte(jwtSecret), nil
		})
		if err != nil || !data.Valid {
			fmt.Println("Error in auth middleware")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]any{
				"success": false,
				"message": "Invalid jwt token",
			})
			return
		}

		// Extract claims
		claim := data.Claims.(jwt.MapClaims)

		// Add payload to request context
		ctx := context.WithValue(r.Context(), Payload, claim)

		fmt.Println("Auth validation successful...")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
