package main

import (
	"fmt"
	"Basic/Utils"
)

// The Basic is module name (verify it from go.mod file)

func Calc(ch rune,a int, b int) float32{
	if(ch=='+'){
		return float32(a+b);
	} else if(ch=='-') {
		return float32(a-b)
	} else if(ch=='/') {
		return float32(a/b)
	} else{
		return float32(a*b)
	}
}


func main(){
	fmt.Println("_________________________")
	// Utils.PrintMessage("Hello,World!");
	// fmt.Println(Utils.ReadFullLine());
	// Utils.Add();
	// Utils.Input();

	// sum := Calc('+',4,5);
	// fmt.Println("sum is ",sum)

	// diff := Calc('-',4,5);
	// fmt.Println("substraction is ",diff)

	// multiplication := Calc('*',4,5);
	// fmt.Println("multiplication is ",multiplication)

	// division := Calc('/',4,5);
	// fmt.Println("division is ",division)

	// fmt.Println(Calc('+',10,20))

	// Utils.ErrorHandling()
	// Utils.Array();

	// Utils.Slice();

	// Utils.IfElse();

	// Utils.Switch();

	// Utils.Loop();

	// Utils.Map();

	// Utils.Structure();

	// Utils.Pointers();

	// Utils.Dataconversion();

	// Utils.Strings();

	// Utils.Time();

	// Utils.Defer();

	Utils.Filehandling();
}