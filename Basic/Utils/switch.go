package Utils

import "fmt"

func Switch() {
	var day string;

	fmt.Println("Enter a day")
	fmt.Scan(&day);

	switch day {
	case "monday" : {
		fmt.Println("Monday")
	}
	case "tuesday" : {
		fmt.Println("Tuesday")
	}
	case "wednesday" : {
		fmt.Println("Wednesday")
	}
	case "thursday" : {
		fmt.Println("Thursday")
	}
	case "friday" : {
		fmt.Println("Friday")
	}
	case "saturday" : {
		fmt.Println("Saturday")
	}
	case "sunday" : {
		fmt.Println("Sunday")
	}
	default : {
		fmt.Println("Not a valid day")
	}
	}


	a := 2
	switch a {
	case 1 : 
		fmt.Println(1)
		fallthrough
	case 2 : 
		fmt.Println(2)
		fallthrough
	case 3 : 
		fmt.Println(3)
		fallthrough
	case 4 : 
		fmt.Println(4)
		fallthrough
	case 5 : 
		fmt.Println(5)
		fallthrough
	case 6 : 
		fmt.Println(6)
		fallthrough
	default : 
		fmt.Println(7)
	}
}