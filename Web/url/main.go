package main  

import ( "fmt"
		"net/url"
	)

func main() {
	fmt.Println("Server Started Successfully")

	URL := "https://dummyjson.com/posts/1"
	fmt.Println("URL is ",URL);
	fmt.Printf("%T \n",URL);

	parsedUrl,err := url.Parse(URL);
	if err!=nil {
		fmt.Println(err)
		return;
	}

	fmt.Printf("Parsed url is %v \n %T \n",parsedUrl,parsedUrl);

	fmt.Println("Scheme",parsedUrl.Scheme)
	fmt.Println("Host",parsedUrl.Host)
	fmt.Println("Path",parsedUrl.Path)
	fmt.Println("RawQuery",parsedUrl.RawQuery)

	// modify the url
	parsedUrl.RawQuery = "username=user123"
	fmt.Println(parsedUrl)

	// convert url to string
	str := parsedUrl.String();
	fmt.Printf("string is %s %T \n",str,str);
}