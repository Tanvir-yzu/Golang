package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Expense struct {
	title  string
	amount float64
}

type Tracker struct {
	expenses []Expense
}

func (t *Tracker) AddExpense(title string, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than 0")
	}
	newExpense := Expense{
		title:  title,
		amount: amount,
	}
	t.expenses = append(t.expenses, newExpense)
	return nil
}

func (t *Tracker) CalculateTotal() (float64, error) {
	total := float64(0)
	if len(t.expenses) == 0 {
		return 0, errors.New("no expenses added")
	}
	for _, expense := range t.expenses {
		total += expense.amount
	}
	return total, nil
}

func (t *Tracker) ShowAllExpenses() {
	if len(t.expenses) == 0 {
		fmt.Println("No expenses added.")
		return
	}
	fmt.Println("All expenses:")
	for i, expense := range t.expenses {
		fmt.Printf("%d. %s: %.2f\n", i+1, expense.title, expense.amount)
	}
}

// Helper function to read input line
func readInput(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func main() {
	tracker := Tracker{}
	var choice int
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n1. Add expense")
		fmt.Println("2. Calculate total expenses")
		fmt.Println("3. Show all expenses")
		fmt.Println("4. Exit")

		fmt.Print("Enter your choice: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		var err error
		choice, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error: Please enter a valid number.")
			continue
		}

		switch choice {
		case 1:
			title := readInput("Enter expense title: ")
			if title == "" {
				fmt.Println("Error: Title cannot be empty.")
				continue
			}

			amountStr := readInput("Enter expense amount: ")
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				fmt.Println("Error: Please enter a valid number for amount.")
				continue
			}

			err = tracker.AddExpense(title, amount)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println("Expense added successfully.")

		case 2:
			total, err := tracker.CalculateTotal()
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Printf("Total expenses: %.2f\n", total)

		case 3:
			tracker.ShowAllExpenses()

		case 4:
			fmt.Println("Exiting the program.")
			return

		default:
			fmt.Println("Invalid choice. Please select a valid option (1-4).")
		}
	}
}
