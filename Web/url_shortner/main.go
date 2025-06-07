package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

type URL struct {
	Original_Url string    `json:"original_url"`
	Date         time.Time `json:"date"`
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var mp = make(map[string]string)

func RandomString(length int) string {
    result := make([]byte, length)
    for i := range result {
        result[i] = charset[rand.Intn(len(charset))]
    }
    return string(result)
}

func DefaultRoute(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintln(w,"Hello,World!")
}

func CreateShortUrl(w http.ResponseWriter,r * http.Request) {
	// the res is in json format 
	res,err := io.ReadAll(r.Body);
	if err!=nil {
		fmt.Println("Error has occured at CreateShortUrl ",err);
		return;
	}

	// fmt.Printf("%T\n",res)
	fmt.Println("Data received from the api is ",string(res))
	
	var URL URL;
	_ = json.Unmarshal(res,&URL)

	URL.Date = time.Now()

	fmt.Println("Data in struct format",URL)

	randomString := RandomString(5);
	fmt.Println("Shortened Url is ",randomString);
	mp[randomString] = URL.Original_Url

	fmt.Println("Mapping is as ",mp[randomString])

	finalString := `http://localhost:3000/redirect?url=` + randomString
	fmt.Println("Shortened url is ",finalString)


	// Returns the response in JSON Format
	response := struct {
		Short_url string `json:"Short_url"`
	}{Short_url: finalString }

	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(response)
}

func redirect(w http.ResponseWriter,r * http.Request) {
	// https://localhost:3000/api?name=John&id=12

	// https://localhost:3000/redirect?url=ab3VeX
	currentUrl := r.URL.Query().Get("url")
	originalUrl := mp[currentUrl]
	fmt.Println(currentUrl,originalUrl)
	http.Redirect(w,r,originalUrl,http.StatusFound)
	fmt.Fprintf(w,currentUrl,originalUrl)
}


func main() {

	http.HandleFunc("/",DefaultRoute)

	http.HandleFunc("/createshorturl",CreateShortUrl)
	http.HandleFunc("/redirect",redirect)

	fmt.Println("Server Started Successfully")
	http.ListenAndServe(":3000",nil)
}
