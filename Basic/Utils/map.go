package Utils

import "fmt"

func Map() {
	mp := make(map[int]int)

	mp[5]++;
	mp[5]++;
	fmt.Println(mp[5])

	value , exist := mp[5];
	fmt.Println(value,exist)

	value , exist = mp[4]
	fmt.Println(value,exist)

	var scores = map[string]int {  "a" : 98, "b" : 97, "c" : 96, "d" : 95}
	fmt.Println(scores)

	for key,value:= range scores {
		fmt.Println(key,value)
	}

	delete(scores,"c")
	fmt.Println(scores)

	value,exist = scores["c"];
	if exist {
		fmt.Println("key exist",value)
	} else {
		fmt.Println("key does not exist")
	}

	value,exist = scores["d"];
	if exist {
		fmt.Println("key exist",value)
	} else {
		fmt.Println("key does not exist")
	}
}
