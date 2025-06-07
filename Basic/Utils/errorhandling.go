package Utils

import "fmt"

func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator can not be zero");
	}

	return a/b,nil;
}

func ErrorHandling() {
	data1,_ := divide(10,20);
	fmt.Println(data1)

	data2,err2 := divide(10,0);
	if(err2!=nil){
		fmt.Println(err2);
	}
	fmt.Println(data2)

}