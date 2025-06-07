package Utils

import "fmt"

func Array() {
	var arr1 [5]int;
	fmt.Println(arr1);

	var arr2 = [5]int{3,2,54,6,2};
	fmt.Println(arr2);

	arr3 := [5]int{1,2,3,4,5};
	fmt.Println(arr3);

	fmt.Println(arr3[0]);
	fmt.Println(arr3[1]);
	fmt.Println(arr3[2]);
	fmt.Println(arr3[3]);
	fmt.Println(arr3[4]);

	fmt.Println(len(arr1),len(arr2),len(arr3))

	fmt.Println("Enter array elements")
	var arr4 [5]int;
	for i:=0;i<5;i++ {
		fmt.Scan(&arr4[i]);
	}

	for i:=0;i<5;i++ {
		fmt.Print(arr4[i]," ")
	}
	fmt.Printf("\n");

	var name[5]string;
	fmt.Printf("%q\n",name);

	name[0] = "abc";
	name[3] = "abc";
	fmt.Println(name);
	fmt.Printf("%q",name);

	matrix := [2][3]int {{1,2,3},{4,5,6}}
	fmt.Println(matrix)
}