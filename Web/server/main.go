package main

import (
	"fmt"
	"net/http"
)

func main() {
	
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		fmt.Println("Hello, World!")
		fmt.Fprint(w,"Hello, World!")
	})

	err := http.ListenAndServe(":3000",nil)
	if err!= nil {
		fmt.Println(err);
		return;
	}
}