package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Methods in golang")

	agb := User{"André-Guy Bruneau", "agbruneau@gmail.com", true, 53}
	fmt.Println(agb)
	fmt.Printf("AGB details are: %+v\n", agb)
	fmt.Printf("Name is %v and email is %v\n", agb.Name, agb.Email)

	agb.GetStatus()
	agb.NewMail()
	fmt.Println(agb)
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}
