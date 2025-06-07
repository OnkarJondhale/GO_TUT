package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TODO struct {
	Id int `json:"id"`
	Todo string `json:"todo"`
	Completed bool `json:"completed"`
	UserId int `json:"userId"`
}

func POST() {
	todoList := TODO{Todo:"this todo is for testing the api",Completed : true,UserId : 123}
	// fmt.Printf("%T\n",todoList)
	fmt.Println("The todoList is ",todoList);

	jsonData,_ := json.Marshal(todoList)
	fmt.Println(string(jsonData))

	res,err := http.Post("https://dummyjson.com/todos/add","application/json",bytes.NewBuffer(jsonData))
	if err!=nil {
		fmt.Println(err)
		return;
	}

	defer res.Body.Close()

	fmt.Println(res.Status)

	var todoListResponse TODO;
	_ = json.NewDecoder(res.Body).Decode(&todoListResponse);
	fmt.Println("Data is ",todoListResponse)
}



func main() {
	fmt.Println("Server Satrted Successfully")

	POST()
}