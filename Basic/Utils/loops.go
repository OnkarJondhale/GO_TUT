package Utils

import "fmt"

func Loop() {
	arr := []int{1, 2, 3, 4, 4, 5, 6, 7, 7, 6, 4, 3, 2}

	fmt.Println(arr)

	for i:=0;i<len(arr);i++ {
		fmt.Print(arr[i]," ")
	}
	fmt.Print("\n")
	fmt.Print("\n")

	for index,value := range arr {
		fmt.Print(index,value,"\n")
	}
	fmt.Print("\n");

	str := "Hello,World!"
	for index,value := range str {
		fmt.Printf("%d %c\n",index,value);
	}

	var count int = 10;
	
	for {
		if count<0 {
			break;
		}
		fmt.Print(count," ");
		count--;
	}
}