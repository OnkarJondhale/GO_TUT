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
)

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating a blog...")

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

	fmt.Println("User payload:", payload)

	var blog model.Blog
	err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		fmt.Println("Error in create blog", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "Invalid credentials",
		})
		return
	}

	fmt.Println("Received data is ", blog)

	var response model.Blog
	_ = utils.DB.QueryRow(`INSERT INTO blog(title,content,userId) values($1,$2,$3) RETURNING id,title,content,userId,createdAt`, blog.Title, blog.Content, payload["id"]).Scan(&response.Id, &response.Title, &response.Content, &response.UserId, &response.CreatedAt)

	fmt.Println("Blog created successfully")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "Blog created successfully",
		"data":    response,
	})
}

func GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all blogs...")

	response, _ := utils.DB.Query(`SELECT * from blog`)
	defer response.Close()

	var blogs []model.Blog
	for response.Next() {
		var blog model.Blog
		_ = response.Scan(&blog.Id, &blog.Title, &blog.Content, &blog.UserId, &blog.CreatedAt)
		blogs = append(blogs, blog)

		if !response.Next() {
			break
		}
	}

	fmt.Println("Blog fetched successfully")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "Blog fetched successfully",
		"data":    blogs,
	})
}

func GetBlogById(w http.ResponseWriter, r *http.Request) {
	qparams := mux.Vars(r)
	id := qparams["postId"]

	query := `
        SELECT 
            y.id, y.title, y.content, y.userId, y.createdAt,
            x.id, x.email, x.fname, x.lname
        FROM users AS x
        INNER JOIN blog AS y
        ON x.id = y.userId 
        WHERE y.id = $1
    `

	var blog model.Blog
	var user model.Users

	err := utils.DB.QueryRow(query, id).Scan(
		&blog.Id, &blog.Title, &blog.Content, &blog.UserId, &blog.CreatedAt,
		&user.Id, &user.Email, &user.FirstName, &user.LastName,
	)

	if err != nil {
		fmt.Println("Database error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("Blog by ID fetched successfully")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "Blog fetched successfully",
		"data": map[string]any{
			"blog": blog,
			"user": user,
		},
	})
}

func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting user details for deleting it...")

	qparams := mux.Vars(r)
	id := qparams["postId"]

	query := `DELETE FROM blog
				WHERE id = $1
				RETURNING id,title,content,userId`

	var blog model.Blog
	_ = utils.DB.QueryRow(query, id).Scan(&blog.Id, &blog.Title, &blog.Content, &blog.UserId)

	fmt.Println("blog details deleted successfully")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "Blog deleted successfully",
		"data":    blog,
	})
}

func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting user details for updating it...")

	qparams := mux.Vars(r)
	id := qparams["postId"]

	var data model.Blog
	json.NewDecoder(r.Body).Decode(&data)

	fmt.Println(data);

	query := `UPDATE blog
				SET content = $1 , title = $2
				WHERE id = $3
				RETURNING id,title,content,userId`

	var blog model.Blog
	_ = utils.DB.QueryRow(query, data.Content, data.Title, id).Scan(&blog.Id, &blog.Title, &blog.Content, &blog.UserId)

	fmt.Println("blog details updated successfully")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "Blog updated successfully",
		"data":    blog,
	})
}
