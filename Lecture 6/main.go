package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	var name [4]string

	name[0]="apple0"
	name[1]="apple1"
	name[2]="apple2"
	name[3]="apple3"

	fmt.Println("Name list :",name)

	var vegList = [3]int{1,5,3}
	fmt.Println(vegList)
}