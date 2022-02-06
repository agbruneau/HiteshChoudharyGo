package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; no super or parent

	agb := User{"André-Guy Bruneau", "agbruneau@gmail.com", true, 53}
	fmt.Println(agb)
	fmt.Printf("AGB details are: %+v\n", agb)
	fmt.Printf("Name is %v and email is %v\n", agb.Name, agb.Email)

}
