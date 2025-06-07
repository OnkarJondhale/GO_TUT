package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TODO struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
	UserId int `json:"userId"`
}

func DELETE() {
	// create a request 
	req,_ := http.NewRequest("DELETE","https://jsonplaceholder.typicode.com/todos/1",nil)

	// create a client
	client := http.Client{}

	// send the request
	res,err := client.Do(req);
	if err!= nil {
		fmt.Println(err)
		return;
	}

	// close the resource
	defer res.Body.Close()
	fmt.Println(res.Status)

	// convert the json byte format into struct format 
	var todoList TODO;
	data := json.NewDecoder(res.Body).Decode(&todoList);
	fmt.Println(data)
}


func main() {
	fmt.Println("Server Started Successfully")

	DELETE();
}