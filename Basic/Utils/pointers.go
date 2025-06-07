package Utils

import (
	"fmt"
	"unsafe"
)

func swapWithoutPointer(a,b int) {
	fmt.Println("Values before swap",a,b)
	t := a 
	a = b 
	b = t 
}

func swapWithPointer(a,b *int) {
	t := *a 
	*a = *b 
	*b = t 
}

func Pointers() {
	var a *int
	fmt.Printf("%T\n",a)

	fmt.Println(unsafe.Sizeof(a))

	var num int = 2;
	var p * int  = &num 

	// Gives same values 
	fmt.Println("value of num is",num,"value of num represented by pointer is",*p)
	fmt.Println("address of num is",&num,"address of num represented by pointer is",p);

	var q ** int = &p;
	fmt.Println("Value of num represented by double pointer",**q,"address of num",*q,"address of pointer p",&p,"address of pointer p represented by double pointer",q)

	fmt.Println("Swap by using pass by value will not work")
	num1 := 10
	num2 := 20
	swapWithoutPointer(num1,num2)
	fmt.Println("After swap",num1,num2)

	fmt.Println("Before swap",num1,num2)
	swapWithPointer(&num1,&num2)
	fmt.Println("After swap",num1,num2)
}