package main

import (
	"fmt"
)

// func main() {
// 	var name string = "TANVIR"
// 	var age int = 25
// 	country := "Bangladesh"
// 	isStudent := true

// 	fmt.Println(name, age, country, isStudent)

// 	// Call functions from func.go (same package)
// 	sum := add(10, 20)
// 	fmt.Println("Sum of 10 + 20 =", sum)

// 	mult := multiply(5, 8)
// 	fmt.Println("Product of 5 * 8 =", mult)
// }

// func main(){

// 	var userage int
// 	fmt.Print("Please enter your age: ")

// 	fmt.Scan(&userage)
// 	fmt.Println("User age is:", userage)
// }
type User struct {
	Name string
	Age int
	email string
}

func (u User) GetInfo() {
	fmt.Println("Name:", u.Name)
	fmt.Println("Age:", u.Age)
	fmt.Println("Email:", u.email)
}

func main() {
	user1 := User{
		Name: "TANVIR",
		Age: 25,
		email: "tanvir@example.com",
	}
	user2 := User{
		Name: "JAHID",
		Age: 20,
		email: "jahid@example.com",
	}
	user1.GetInfo()
	user2.GetInfo()
}