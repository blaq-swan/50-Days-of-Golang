package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What do you think about our services?")

	// comma ok || err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)
	fmt.Printf("Type of this rating is %T", input)

}
