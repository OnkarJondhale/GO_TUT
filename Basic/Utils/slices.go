package Utils

import "fmt"

func Slice() {
	slice1 := []int{1, 2, 3, 4, 5}

	fmt.Printf("%T\n",slice1)
	fmt.Println(slice1)

	slice2 := make([]int,5);
	fmt.Println(slice2)

	fmt.Println(slice1[4]);

	slice1 = append(slice1,6,7,8,9,10);
	fmt.Println(len(slice1));
	fmt.Println(slice1)

	matrix := [][]int{{1,2,3},{4,5}}
	fmt.Println(matrix);

	var slice3 []int;
	fmt.Println(slice3)
	slice3 = append(slice3,4);
	fmt.Println(slice3);
	slice3 = append(slice3,6);
	fmt.Println(slice3);
	slice3 = append(slice3,5);
	fmt.Println(slice3);
	slice3 = append(slice3,8);
	fmt.Println(slice3);

	index := 1
	slice3 = append(slice3[:index],slice3[index+1:]...)
	fmt.Println(slice3);
}