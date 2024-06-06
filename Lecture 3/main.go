package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("User Input: ")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:")


	//comma ok || err ok

	// input , _ := reader.ReadString('\n')
	input , err := reader.ReadString('\n')
	fmt .Println("Thanks for rating, ", input)
	fmt .Printf("Thanks for rating, %T", input)
}
