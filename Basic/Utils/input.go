package Utils

import "fmt"

func Input() int{
	var a int;

	fmt.Println("Enter a number");
	fmt.Scan(&a);

	return a;
}

