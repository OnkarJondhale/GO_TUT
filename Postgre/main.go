package main

import (
	controller "Blog/controllers"
	auth "Blog/middlewares"
	"Blog/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Server started successfully")

	router := mux.NewRouter()

	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register", controller.Register).Methods("POST")
	authRouter.HandleFunc("/login", controller.Login).Methods("POST")

	blogRouter := router.PathPrefix("/blog").Subrouter()
	blogRouter.HandleFunc("/get", controller.GetAllBlogs).Methods("GET")
	blogRouter.HandleFunc("/get/{postId}", controller.GetBlogById).Methods("GET")

	protectedBlogRouter := router.PathPrefix("/blog").Subrouter()
	protectedBlogRouter.Use(auth.Auth)
	protectedBlogRouter.HandleFunc("/create", controller.CreateBlog).Methods("POST")
	protectedBlogRouter.HandleFunc("/update/{postId}", controller.UpdateBlog).Methods("PUT")
	protectedBlogRouter.HandleFunc("/delete/{postId}", controller.DeleteBlog).Methods("DELETE")

	commentRouter := router.PathPrefix("/comment").Subrouter()
	commentRouter.HandleFunc("/get/{postId}", controller.GetComment).Methods("GET")

	protectedCommentRouter := router.PathPrefix("/comment").Subrouter()
	protectedCommentRouter.Use(auth.Auth)
	protectedCommentRouter.HandleFunc("/create/{postId}", controller.CreateComment).Methods("POST")
	protectedCommentRouter.HandleFunc("/update/{commentId}", nil).Methods("PUT")
	protectedCommentRouter.HandleFunc("/delete/{commentId}", nil).Methods("DELETE")

	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/get/{userId}", controller.GetUser).Methods("GET")

	protectedUserRouter := router.PathPrefix("/user").Subrouter()
	protectedUserRouter.Use(auth.Auth)
	protectedUserRouter.HandleFunc("/update/details", controller.UpdateUserDetails).Methods("PUT")
	protectedUserRouter.HandleFunc("/update/password", controller.UpdateUserPassword).Methods("PUT")
	protectedUserRouter.HandleFunc("/delete", controller.DeleteUser).Methods("DELETE")

	utils.ConnectDb()

	http.ListenAndServe(":3000", router)
}
