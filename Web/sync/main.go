package main

import (
	"fmt"
	"sync"
)

func Task(Task int,wg *sync.WaitGroup) {
	defer wg.Done();
	fmt.Println("Started ",Task);
	fmt.Println("Ended ",Task);
}

func main() {

	var wg sync.WaitGroup;
	for i:=1;i<=3;i++ {
		wg.Add(1)
		go Task(i,&wg);
	}

	wg.Wait();
	fmt.Println("Task Completed")
}

