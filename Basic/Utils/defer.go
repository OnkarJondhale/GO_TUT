package Utils 

import "fmt"

func Defer() {
	fmt.Println("Start");
	defer fmt.Println("Deferred execution--> This will be executed at last as defer follows LIFO manner");
	fmt.Println("End");

	fmt.Print("\n")

	fmt.Println("Start");
	defer fmt.Println("Deferred execution 1");
	defer fmt.Println("Deferred execution 2");
	defer fmt.Println("Deferred execution 3");
	defer fmt.Println("Deferred execution 4");
	defer fmt.Println("Deferred execution 5");
	defer fmt.Println("Deferred execution 6");
	fmt.Println("End")
}