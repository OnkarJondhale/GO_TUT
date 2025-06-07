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
	Todo string `json:"todo"`
	Completed bool `json:"completed"`
	UserId int `json:"userId"`
}

func GetRequest(URL string) (* http.Response,error) {
	res,err := http.Get(URL);
	if err!=nil {
		return nil, fmt.Errorf("error has occurred: %v", err)
	}

	return res,nil
}

func GetResponseData(res * http.Response) ([]byte,error) {
	data,err := io.ReadAll(res.Body)
	if err!= nil {
		return nil,fmt.Errorf("error has occured: %v",err)
	}

	return data,nil
}

func ConvertJsonToStruct(data []byte) {
	var todoList TODO;
	err := json.Unmarshal(data,&todoList);
	if err!=nil {
		fmt.Print(err)
		return;
	}
	
	fmt.Println("The json data is converted into the string using Unmarshalling",todoList)
	fmt.Println("ID ",todoList.Id)
	fmt.Println("TODO ",todoList.Todo)
	fmt.Println("COMPLETED ",todoList.Completed)
	fmt.Println("USERID ",todoList.UserId)
}

func Decoder(data []byte) {
	var todoList TODO;
	err := json.NewDecoder(bytes.NewReader(data)).Decode(&todoList)
	if err != nil {
		fmt.Println(err);
		return;
	}
	fmt.Println("The json data is converted into the string using json.NewDecoder().Decode() method",todoList)
	fmt.Println("ID ",todoList.Id)
	fmt.Println("TODO ",todoList.Todo)
	fmt.Println("COMPLETED ",todoList.Completed)
	fmt.Println("USERID ",todoList.UserId)
}

func GET() {
	res,err := GetRequest("https://dummyjson.com/todos/1");
	if err != nil {
		fmt.Println(err);
		return;
	}
	defer res.Body.Close();

	if(res.StatusCode!=http.StatusOK) {
		fmt.Println(res.Status)
		return;
	}

	// Convert to struct without converting it to data
	/* 
		var todoList TODO;
		err := json.NewDecoder(res.Body).Decode(&todoList);
		if err!=nil {
			fmt.Println(err)
			return;
		}
		fmt.Println("Data is ",todoList)
	*/

	data,err := GetResponseData(res)
	if err != nil {
		fmt.Println(err)
		return;
	}
	fmt.Println("Data received from the API call is ",string(data));

	ConvertJsonToStruct(data)
	fmt.Printf("\n");
	Decoder(data);
}

func main() {
	fmt.Println("Server Started Successfully")

	GET();
}
