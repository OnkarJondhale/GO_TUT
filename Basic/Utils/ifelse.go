package Utils

import "fmt"

func IfElse() {
	fmt.Println("Print greatest of three numbers");

	var a,b,c int;

	fmt.Scan(&a,&b,&c);

	// Ladder 
	if a>=b && a>=c {
		fmt.Printf("%d is greater\n",a);
	} else if b>=a && b>=c {
		fmt.Printf("%d is greater\n",b);
	} else {
		fmt.Printf("%d is greater\n",c);
	}

	// Nested 
	if a>=b {
		if a>=c {
			fmt.Printf("%d is greater\n",a);
		} else {
			fmt.Printf("%d is greater\n",c)
		}
	} else if b>=a {
		if b>=c {
			fmt.Printf("%d is greater\n",b);
		} else {
			fmt.Printf("%d is greater\n",c);
		}
	}
}