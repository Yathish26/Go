package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Account struct {
	Name    string
	Balance int
	AccNo   int
}

var accounts []Account
var currentAccount *Account

func main() {
	header()
}

func header() {
	for {
		var option int
		fmt.Println("\nWelcome to the Bank")
		fmt.Println("Please select an option:")
		fmt.Println("1. Create Account")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Check Balance")
		fmt.Println("5. Switch Account")
		fmt.Println("6. Exit")

		fmt.Scanln(&option)
		switch option {
		case 1:
			createAccount()
		case 2:
			deposit()
		case 3:
			withdraw()
		case 4:
			checkBalance()
		case 5:
			switchAccount()
		case 6:
			exit()
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func createAccount() {
	var name string
	fmt.Println("Create Account")
	fmt.Print("Enter Name: ")
	fmt.Scanln(&name)

	// Generate unique account number
	accNo := rand.Intn(1000000) + 1000000
	newAccount := Account{
		Name:    name,
		Balance: 0,
		AccNo:   accNo,
	}
	accounts = append(accounts, newAccount)
	currentAccount = &accounts[len(accounts)-1]

	fmt.Println("Account created successfully!")
	fmt.Printf("Account Holder: %v\n", name)
	fmt.Printf("Account Number: %v\n", accNo)
}

func deposit() {
	if currentAccount == nil {
		fmt.Println("No account selected. Please create or switch to an account.")
		return
	}
	var amount int
	fmt.Println("How much money do you want to deposit?")
	fmt.Scanln(&amount)
	if amount > 0 {
		currentAccount.Balance += amount
		fmt.Printf("%v Rs deposited successfully!\n", amount)
		fmt.Printf("New Balance: %v Rs\n", currentAccount.Balance)
	} else {
		fmt.Println("Invalid deposit amount.")
	}
}

func withdraw() {
	if currentAccount == nil {
		fmt.Println("No account selected. Please create or switch to an account.")
		return
	}
	var amount int
	fmt.Println("How much money do you want to withdraw?")
	fmt.Scanln(&amount)
	if amount > 0 && amount <= currentAccount.Balance {
		currentAccount.Balance -= amount
		fmt.Printf("%v Rs withdrawn successfully!\n", amount)
		fmt.Printf("New Balance: %v Rs\n", currentAccount.Balance)
	} else if amount > currentAccount.Balance {
		fmt.Println("Insufficient balance for this withdrawal.")
	} else {
		fmt.Println("Invalid withdrawal amount.")
	}
}

func checkBalance() {
	if currentAccount == nil {
		fmt.Println("No account selected. Please create or switch to an account.")
		return
	}
	fmt.Println("Check Balance")
	fmt.Printf("Account Number: %v\n", currentAccount.AccNo)
	fmt.Printf("Account Holder: %v\n", currentAccount.Name)
	fmt.Printf("Your Total Balance: %v Rs\n", currentAccount.Balance)
}

func switchAccount() {
	if len(accounts) == 0 {
		fmt.Println("No accounts available. Please create one first.")
		return
	}
	var accNo int
	fmt.Println("Enter your account number to switch:")
	fmt.Scanln(&accNo)
	for _, account := range accounts {
		if account.AccNo == accNo {
			currentAccount = &account
			fmt.Printf("Switched to account: %v (Holder: %v)\n", accNo, account.Name)
			return
		}
	}
	fmt.Println("Account not found. Please try again.")
}

func exit() {
	fmt.Println("Thank you for using our banking services!")
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
