package main

import (
	"bufio"
	"fmt"
	"os"

	mapset "github.com/deckarep/golang-set"
)

func main() {
	fmt.Print("Enter the Connection List: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	input := scanner.Text()
	fmt.Println(input)
	size := len(input)

	var adjList = make(map[byte]string)
	var lval byte

	active := true

	for i := 0; i < size; i++ {
		switch input[i] {
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
				active = false
			} else {
				adjList[lval] = adjList[lval] + string(input[i])
				active = true
			}
		}
	}

	var friend1, friend2 byte

	fmt.Print("Enter the pair of users (Format : A B): ")
	fmt.Scanf("%c %c", &friend1, &friend2)

	toDoSet := mapset.NewSet()
	doneSet := mapset.NewSet()

	toDoSet.Add(friend1)
	var elem byte

	for toDoSet.Cardinality() != 0 {
		it := toDoSet.Iterator()
		for temp := range it.C {
			elem = temp.(byte)
			doneSet.Add(elem)
			toDoSet.Remove(elem)
			it.Stop()
		}
		reachable := adjList[elem]
		for i := 0; i < len(reachable); i++ {
			if reachable[i] == friend2 {
				fmt.Println("True")
				return
			}
			if !doneSet.Contains(reachable[i]) {
				toDoSet.Add(reachable[i])
			}
		}
	}
	fmt.Println("False")
}
