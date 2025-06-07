package Utils

import (
	"fmt"
	"strconv"
)

func StringToInteger(a string) {
	num, _ := strconv.Atoi(a)
	fmt.Printf("%T %T %s %d\n",a,num,a,num)
}

func StringTOFloat(a string) {
	num, _ := strconv.ParseFloat(a,64)
	fmt.Printf("%T %T %s %f\n",a,num,a,num)
}

func StringToBoolean(a string) {
	num, _ := strconv.ParseBool(a)
	fmt.Printf("%T %T %s %t\n",a,num,a,num)
}

func IntegerToString(a int) {
	num := strconv.Itoa(a)
	fmt.Printf("%T %T %d %s\n",a,num,a,num)
}


func FloatToString(a float64) {
	num := fmt.Sprintf("%f",a)
	fmt.Printf("%T %T %f %s\n",a,num,a,num)
}


func BoolToString(a bool) {
	num := strconv.FormatBool(a)
	fmt.Printf("%T %T %t %s\n",a,num,a,num)
}


func Dataconversion() {

	// Interger to Float
	var a int = 4
	var b float32 = float32(a)
	fmt.Printf("%T %T %d %f\n",a,b,a,b);

	// Float to integer
	var c float32 = 3.2 
	var d int = int(c)
	fmt.Printf("%T %T %f %d\n",c,d,c,d)

	// string to integer
	StringToInteger("123213")

	// string to float64
	StringTOFloat("2342.423")

	// string to bool
	StringToBoolean("false")

	// integer to string
	IntegerToString(1232)

	// float to string
	FloatToString(2342.3423)

	// bool to string 
	BoolToString(true)
}