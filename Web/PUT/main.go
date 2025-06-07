package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TODO struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
	UserId int `json:"userId"`
}

func UPDATE() {
	// get the data to update
	todoList,err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err!=nil {
		fmt.Println(err);
		return;
	}
	defer todoList.Body.Close()
	jsonData,_ := io.ReadAll(todoList.Body)
	fmt.Println("Data received from the api is ",string(jsonData))

	// convert the json data into struct and update it 
	var structData TODO;
	_ = json.Unmarshal(jsonData,&structData)
	fmt.Println("Data after conversion ",structData)
	structData.Title = "This title is updated"
	fmt.Println("Data after updation in struct format",structData)

	// convert the updated data (struct) into the json data 
	updatedJsonData,_ := json.Marshal(structData)
	// fmt.Println("Data after updation in json format",updatedJsonData)

	// Create the PUT request 
	req,_ := http.NewRequest("PUT","https://jsonplaceholder.typicode.com/todos/1",bytes.NewBuffer(updatedJsonData))
	req.Header.Set("Content-type","application/json");

	// Create the client to send the request
	client := http.Client{};
	res,err := client.Do(req);
	if err != nil {
		fmt.Println(err)
		return;
	}
	defer res.Body.Close()

	fmt.Println(res.Status)

	// Convert the json byte array into readable format
	var data TODO;
	_ = json.NewDecoder(res.Body).Decode(&data);
	fmt.Println(data)
}


func main() {
	fmt.Println("Server Started Successfully")

	UPDATE()
}