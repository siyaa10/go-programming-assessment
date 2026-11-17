package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type BankAccount struct {
	Name    string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) {

	if amount < 0 {
		fmt.Println("Invalid deposit amount")
		return
	}
	b.Balance = b.Balance + amount
}

func (b *BankAccount) Withdraw(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("withdrawal amount cannot be negative")
	}
	if amount > b.Balance {
		return fmt.Errorf("insufficient balance")
	}

	b.Balance = b.Balance - amount
	return nil
}

func main() {

	for {
		fmt.Println("\n================================")
		fmt.Println("     GO PROGRAMMING ASSESSMENT")
		fmt.Println("================================")
		fmt.Println("1. Bank Account")
		fmt.Println("2. Notification System")
		fmt.Println("3. Age Validation")
		fmt.Println("4. File Handling")
		fmt.Println("5. Goroutines")
		fmt.Println("6. Safe Counter")
		fmt.Println("7. Discount Calculator")
		fmt.Println("0. Exit")

		var choice int

		fmt.Print("\nEnter your choice: ")
		_, err := fmt.Scan(&choice)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			fmt.Scanln()
			continue
		}

		switch choice {

		case 1:
			runBankAccountDemo()

		case 2:
			runNotifierDemo()

		case 3:
			runValidationDemo()

		case 4:
			runFileDemo()

		case 5:
			runGoroutineDemo()

		case 6:
			runSafeCounterDemo()

		case 7:
			runDiscountDemo()

		case 0:
			fmt.Println("Thank you for using the application!")
			return

		default:
			fmt.Println("Invalid choice! Please select 0 to 7.")
		}
	}
}

// Topic 1: Bank Account demonstration
func runBankAccountDemo() {

	fmt.Println("\n===== BANK ACCOUNT =====")

	account := BankAccount{
		Name:    "Siya",
		Balance: 1000,
	}

	fmt.Println("Account Holder:", account.Name)
	fmt.Println("Initial Balance:", account.Balance)

	var depositAmount float64

	fmt.Print("Enter amount to deposit: ")
	fmt.Scan(&depositAmount)

	account.Deposit(depositAmount)

	fmt.Println("Balance after deposit:", account.Balance)

	var withdrawAmount float64

	fmt.Print("\nEnter amount to withdraw: ")
	fmt.Scan(&withdrawAmount)

	err := account.Withdraw(withdrawAmount)

	if err != nil {
		fmt.Println("Withdrawal failed:", err)
	} else {
		fmt.Println("Withdrawal successful!")
	}

	fmt.Println("Final Balance:", account.Balance)
}

// Topic 2: Notification demonstration
func runNotifierDemo() {
	fmt.Println("\n===== NOTIFICATION SYSTEM =====")

	fmt.Println("1. Email")
	fmt.Println("2. SMS")

	var choice int
	fmt.Print("Choose notification type: ")
	_, err := fmt.Scan(&choice)

	if err != nil {
		fmt.Println("Invalid input!")
		return
	}

	var notifier Notifier

	switch choice {
	case 1:
		notifier = EmailNotifier{}
	case 2:
		notifier = SMSNotifier{}
	default:
		fmt.Println("Invalid choice!")
		return
	}

	var message string
	fmt.Print("Enter your message: ")
	fmt.Scanln()
	message, _ = bufio.NewReader(os.Stdin).ReadString('\n')

	message = strings.TrimSpace(message)

	err = notifier.SendNotification(message)
	if err != nil {
		fmt.Println("Notification failed:", err)
	}
}
