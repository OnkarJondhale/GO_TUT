package main

import (
	"fmt"
	"io"
	"net/http"
)

func sendRequest(url string) *http.Response {
	res,err := http.Get(url)
	if err!=nil {
		fmt.Println(err);
		return nil;
	}

	return res
}

func readJsonResponse(res * http.Response) []uint8{
	data,err := io.ReadAll(res.Body)
	if err!=nil {
		fmt.Println(err)
		return nil;
	}

	return data;
}

func main() {
	fmt.Println("Server Started Successfully")

	res := sendRequest("https://dummyjson.com/posts/1");
	defer res.Body.Close()

	data := readJsonResponse(res);
	fmt.Println("Data in string",string(data));
}

