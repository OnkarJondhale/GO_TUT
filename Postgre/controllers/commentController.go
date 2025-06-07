package controller

import (
	auth "Blog/middlewares"
	model "Blog/models"
	"Blog/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	// Get the postId
	qparams := mux.Vars(r)
	postId := qparams["postId"]

	// Get the userId from JWT
	payload, ok := r.Context().Value(auth.Payload).(jwt.MapClaims)
	if !ok {
		fmt.Println("Error: Missing payload in context")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "Unauthorized access",
		})
		return
	}

	// Get the comment content properly
	var input struct {
		Content string `json:"content"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	var comment model.Comment
	query := `INSERT INTO comment(userId, postId, content) VALUES ($1, $2, $3) RETURNING id, postId, userId, content, createdAt`

	err = utils.DB.QueryRow(query, payload["id"], postId, input.Content).Scan(
		&comment.Id, &comment.PostId, &comment.UserId, &comment.Content, &comment.CreatedAt,
	)
	if err != nil {
		fmt.Println("Error inserting comment:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("Comment created successfully")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "Comment created successfully",
		"data":    comment,
	})
}

func GetComment(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Fetching comments...")

    qparams := mux.Vars(r)
    postId := qparams["postId"]

    query := `SELECT x.content, x.createdAt, z.fname, z.lname, z.email
    FROM comment AS x 
    INNER JOIN blog AS y ON x.postId = y.id
    INNER JOIN users AS z ON x.userId = z.id
    WHERE y.id = $1`

    data, _ := utils.DB.Query(query, postId)
    defer data.Close()

    type response struct {
        Content   string    `json:"content,omitempty" db:"content"`
        CreatedAt time.Time `json:"createdAt,omitzero" db:"createdAt"`
        Fname     string    `json:"fname,omitempty" db:"fname"`
        Lname     string    `json:"lname,omitempty" db:"lname"`
        Email     string    `json:"email,omitempty" db:"email"`
    }

    var comments []response
    for {
        var comment response
        _ = data.Scan(&comment.Content, &comment.CreatedAt, &comment.Fname, &comment.Lname, &comment.Email)
        comments = append(comments, comment)
        if !data.Next() {
            break
        }
    }

    fmt.Println("Comments fetched successfully")

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]any{
        "success": true,
        "message": "Comments fetched successfully",
        "data":    comments,
    })
}
