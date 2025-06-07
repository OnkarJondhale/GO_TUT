package main

import (
	"fmt"
	"time"
)


func Hello() {
	fmt.Println("2) Hello");

	time.Sleep(2000*time.Millisecond)
}

func Hi() {
	fmt.Println("3) Hi");
}


func withoutGoRoutines() {
	fmt.Println("1) Welcome")
	Hello();
	Hi();
	fmt.Println("4) How are you")
}

func withGoRoutines() {
	fmt.Println("1) Welcome")
	go Hello();
	Hi();
	fmt.Println("4) How are you")

	time.Sleep(1000*time.Millisecond)
}


func main() {
	withoutGoRoutines();
	fmt.Printf("\n");
	withGoRoutines();
}