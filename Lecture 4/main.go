package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter your input: ")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Your Enter number: ", input)
	fmt.Printf("Your Enter number: %T", input)
	input = strings.TrimSpace(input)

	sqr , err := strconv.ParseInt(input,10,64)
		
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Printf("Your Enter number: %T", sqr)
	}

}