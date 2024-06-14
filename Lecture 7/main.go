package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome")


	var number = []int{1,2,3,4,5,67}
	fmt.Println(number)
	fmt.Printf("Type of fruitlist is %T\n",number)

	number = append(number,1,5)
	fmt.Println(number)

	number = append(number[3:])
	fmt.Println(number)

	highScore := make([]int,4)

	highScore[0] = 22
	highScore[1] = 23
	highScore[2] = 254
	highScore[3] = 25
	// highScore[4] = 125 illegal

	

	highScore = append(highScore,1,2,3,)
	fmt.Println(highScore)
	
	sort.Ints(highScore)
}