package main

import "fmt"

const LoginToken string = "nothingimportant"

func main() {
	var username string = "blaqswan"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isActive bool = true
	fmt.Println(isActive)
	fmt.Printf("Variable is of type: %T \n", isActive)

	var smallFloat float32 = 255.6336359999
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	// default values and some aliases
	var randomVariable int
	fmt.Println(randomVariable)
	fmt.Printf("Variable is of type: %T \n", randomVariable)

	var stringVariable string
	fmt.Println(stringVariable)
	fmt.Printf("Variable is of type: %T \n", stringVariable)

	// Another way of declaring Variable
	fmt.Println()
	numberOfUsers := 250000
	fmt.Println(numberOfUsers)

	fmt.Println()
	fmt.Printf("%s is of type %T \n", LoginToken, LoginToken)
}
