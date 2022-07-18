/*

GO Output Functions

Go has three functions to output text:

- print()
- Println()
- Printf()

*/

// The print() Function

// The print function prints its arguments with their defualt format
/*
package main

import "fmt"

func main() {
	var i, j string = "hello", "world"

	fmt.Print(i, "\n", j)
	// fmt.Print(j, "\n")
}

package main

import "fmt"

func main() {
	var i, j = "Hello", "World"

	fmt.Print(i, " ", j, "\n")

	var a, b = 12, 14

	fmt.Print(a, b)
}


// Println() Function
// Println() function is similiar

package main

import "fmt"

func main() {
	var i, j string = "Hello", "World"

	fmt.Println(i, j)
}

*/

/*
The printf() Function

It first formats its arguments based on the given formatting verb and then prints them.

Here we will use two formatting verbs:

- %v  value
- %T  type

*/

package main

import "fmt"

func main() {
	var i string = "Hello"
	var j int = 16

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T", j, j)
}
