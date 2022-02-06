package main

import "fmt"

const LoginToken = "aseljfiasdlkfjasdlfjkiasdflkjasdflkjfd"

func main() {
	// Used of the const variable that is public because of the First Capitalize Letter in the Name of this variable
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

	// type of variables
	var username string = "agbruneau"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallInt uint8 = 255
	fmt.Println(smallInt)
	fmt.Printf("Variable is of type: %T \n", smallInt)

	var smallFloat float64 = 255.45544511254451885
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	// no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T \n", numberOfUser)
}
