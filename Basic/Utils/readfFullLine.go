package Utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFullLine() string {

	fmt.Println("Enter a full string")
	reader := bufio.NewReader(os.Stdin)
	result, _ := reader.ReadString('\n')
	return result;
}