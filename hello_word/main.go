package main

import (
	"errors"
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



// type BankAccount struct {
// 	AccountNumber string
// 	UserName string
// 	Password string
// 	Balance float64
// 	IsLoggedIn bool
// }
// func (b *BankAccount) Login(inputUser string, inputPassword string) {
// 	if inputUser == b.UserName && inputPassword == b.Password {
// 		b.IsLoggedIn = true
// 		fmt.Println("Login successful.")
// 	} else {
// 		fmt.Println("Login failed.")
// 	}
// }
// func (b *BankAccount) Logout() {
// 	b.IsLoggedIn = false
// 	fmt.Println("Logout successful.")
// }

// func (b *BankAccount) Deposit(amount float64) {
// 	if !b.IsLoggedIn {
// 		fmt.Println("Please login first.")
// 		return
// 	}
// 	b.Balance += amount
// 	fmt.Println("Deposit of", amount, "was successful.")
// }

// func (b *BankAccount) Withdraw(amount float64) {
// 	if !b.IsLoggedIn {
// 		fmt.Println("Please login first.")
// 		return
// 	}
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
// 		UserName: "rahim",
// 		Password: "123456",
// 		Balance: 1000.0,
// 	}
// 	//rahim.Login("rahim", "123456")
// 	rahim.Deposit(500.0)
// 	// rahim.Withdraw(200.0)
// 	//rahim.Logout()
// 	//rahim.Withdraw(200.0)
// 	//rahim.GetBalance()
// 	//rahim.Logout()
// }

// func main() {

// 	var sliceUsers []string 
// 	sliceUsers = append(sliceUsers, "rahim")
// 	sliceUsers = append(sliceUsers, "jahid")
// 	fmt.Println(sliceUsers[0])

// 	for i, v := range sliceUsers {
// 		fmt.Println(i, v)
// 	}

	
// 	func main() {

// 	var prices []float64 
// 	prices = append(prices, 100.0)
// 	prices = append(prices, 200.0)
// 	prices = append(prices, 300.0)
// 	total := 0.0

// 	for _, v := range prices {
// 		total += v
// 	}
// 	fmt.Println("Total price:", total)
// }

// func main() {
// 	// ১. Map তৈরি করা (make ফাংশন ব্যবহার করে)
// 	UserAge := make(map[string]int)
// 	UserAge["rahim"] = 25
// 	UserAge["jahid"] = 20
// 	fmt.Println("rahim age", UserAge["rahim"])
// 	fmt.Println("jahid age", UserAge["jahid"])

// 	age, ok := UserAge["rahim"]
// 	if ok {
// 		fmt.Println("rahim age:", age)
// 	} else {
// 		fmt.Println("rahim age not found")
// 	}

// 	delete(UserAge, "rahim")
// 	fmt.Println(" After deleting rahim age:", UserAge)
// }

// func main(){
// 	bookPrices := make(map[string]float64)
// 	bookPrices["Book1"] = 100.0
// 	bookPrices["Book2"] = 200.0
// 	bookPrices["Book3"] = 300.0
// 	book, ok := bookPrices["Book4"]
// 	if ok {
// 		fmt.Println("Book4 price:", book)
// 	} else {
// 		fmt.Println("Book4 price not found")
// 	}
// 	fmt.Println("book4",bookPrices["Book4"])
// }

// func divide(a int, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("division by zero")
// 	}
// 	return a / b, nil
// }

// func main() {
// 	result, err := divide(10, 2)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Result:", result)
// 	result2, err := divide(10, 0)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Result2:", result2)
// }

func checkAge(age int) error {
	if age < 18 {
		return errors.New("age cannot be under 18")
	}
	return nil
}

func main() {
	age := 25
	err := checkAge(age)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Age is valid.")
}