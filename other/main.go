package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter the Connection List: ")
	var input string
	size, _ := fmt.Scanln(&input)

	var adjList = make(map[byte]string)
	var lval, rval byte
	
	for i := 0; i < size; i++ {
		switch active := false; input[i] {
		case '(':
		case ')':
		case '[':
		case ']':
		case ' ':
		case ',':
			break
		default:
			if active {
				lval = input[i]
			}
			if !active {
				adjList[lval] += input[i]
			}
			active != active
		}
	}

	var friend1, friend2 byte

	fmt.Print("Enter the pair of users (Format : A B): ")
	fmt.Scanf("%c %c", &friend1, &friend2)


	for ;


}