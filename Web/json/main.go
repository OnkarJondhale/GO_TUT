package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	IsMale bool `json:"gender"` // after marshalling this field is referred by gender while after unmarshalling it is referred as IsMale
}


func main() {
	fmt.Println("Server Started Successfully")

	var a = Person{"abc",34,true};
	fmt.Println(a)

	fmt.Printf("structure %T\n",a);

	jsonData,err := json.Marshal(a)
	if err!=nil {
		fmt.Println(err)
		return;
	}

	fmt.Println(string(jsonData))
	fmt.Printf("object %T\n",jsonData)

	var new_data Person
	err = json.Unmarshal([]byte(jsonData),&new_data);
	if err!=nil {
		fmt.Println(err)
		return;
	}

	fmt.Println(new_data)
	fmt.Printf("structure %T\n",new_data)

}