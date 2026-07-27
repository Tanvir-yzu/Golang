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

// oop example
// type User struct {
// 	Name string
// 	Age int
// 	email string
// }

// func (u User) GetInfo() {
// 	fmt.Println("Name:", u.Name)
// 	fmt.Println("Age:", u.Age)
// 	fmt.Println("Email:", u.email)
// }

// func main() {
// 	user1 := User{
// 		Name: "TANVIR",
// 		Age: 25,
// 		email: "tanvir@example.com",
// 	}
// 	user2 := User{
// 		Name: "JAHID",
// 		Age: 20,
// 		email: "jahid@example.com",
// 	}
// 	user1.GetInfo()
// 	user2.GetInfo()
// }

// একটি মিনি ব্যাংক অ্যাকাউন্ট সিস্টেম তৈরি করি।
// type BankAccount struct {
// 	AccountNumber string
// 	Balance float64
// }

// func (b *BankAccount) Deposit(amount float64) {
// 	b.Balance += amount
// 	fmt.Println("Deposit of", amount, "was successful.")
// }

// func (b *BankAccount) Withdraw(amount float64) {
// 	if amount > b.Balance {
// 		fmt.Println("Insufficient balance. Withdrawal failed.")
// 		return
// 	}
// 	b.Balance -= amount
// 	fmt.Println("Withdrawal of", amount, "was successful.")
// }

// func (b *BankAccount) GetBalance() {
// 	fmt.Println("Current balance:", b.Balance)
// }

// func main(){

// 	rahim := BankAccount{
// 		AccountNumber: "1234567890",
// 		Balance: 1000.0,
// 	}
// 	rahim.Deposit(500.0)
// 	rahim.Withdraw(200.0)
// 	rahim.GetBalance()
// }



type BankAccount struct {
	AccountNumber string
	UserName string
	Password string
	Balance float64
	IsLoggedIn bool
}
func (b *BankAccount) Login(inputUser string, inputPassword string) {
	if inputUser == b.UserName && inputPassword == b.Password {
		b.IsLoggedIn = true
		fmt.Println("Login successful.")
	} else {
		fmt.Println("Login failed.")
	}
}

func (b *BankAccount) Deposit(amount float64) {
	if !b.IsLoggedIn {
		fmt.Println("Please login first.")
		return
	}
	b.Balance += amount
	fmt.Println("Deposit of", amount, "was successful.")
}

func (b *BankAccount) Withdraw(amount float64) {
	if !b.IsLoggedIn {
		fmt.Println("Please login first.")
		return
	}
	if amount > b.Balance {
		fmt.Println("Insufficient balance. Withdrawal failed.")
		return
	}
	b.Balance -= amount
	fmt.Println("Withdrawal of", amount, "was successful.")
}

func (b *BankAccount) GetBalance() {
	fmt.Println("Current balance:", b.Balance)
}

func main(){

	rahim := BankAccount{
		AccountNumber: "1234567890",
		UserName: "rahim",
		Password: "123456",
		Balance: 1000.0,
	}
	rahim.Login("rahim", "123456")
	rahim.Deposit(500.0)
	rahim.Withdraw(200.0)
	rahim.GetBalance()
}