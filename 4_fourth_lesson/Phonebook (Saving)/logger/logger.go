package logger

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"phonebook/book"
	"strings"
	"time"
)

func main() {
	phoneBook := make(book.PhoneBook)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome!")
	fmt.Println("available commands: add, get, delete, update, list, exit")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		command := parts[0]

		switch command {
		case "add":
			kv := strings.SplitN(parts[1], "=", 2)
			if len(kv) != 2 {
				fmt.Println("Invalid format. Use: add name=number")
				continue
			}
			name, number := kv[0], kv[1]
			phoneBook[name] = book.PhoneNumber{
				Number:        number,
				LastUpdatedAt: time.Now().Unix(),
			}
			fmt.Printf("Added/Updated: %s -> %s\n", name, number)

		case "get":
			name := parts[1]
			number, exist := phoneBook[name]
			if exist {
				fmt.Printf("Number for %s is %s\n", name, number)
			} else {
				fmt.Printf("No entry found for %s\n", name)
			}

		case "delete":
			name := parts[1]
			_, exists := phoneBook[name]
			if exists {
				delete(phoneBook, name)
				fmt.Printf("Deleted entry for %s\n", name)
			} else {
				fmt.Printf("No entry to delete for %s\n", name)
			}

		case "update":
			kv := strings.SplitN(parts[1], "=", 2)
			if len(kv) != 2 {
				fmt.Println("Invalid format. Use: update name=number")
				continue
			}

			name, newNumber := kv[0], kv[1]
			_, exists := phoneBook[name]
			if exists {
				phoneBook[name] = book.PhoneNumber{
					Number:        newNumber,
					LastUpdatedAt: time.Now().Unix(),
				}
				fmt.Printf("Updated %s -> %s\n", name, newNumber)
			} else {
				fmt.Println("No such entry to update")
			}

		case "list":
			if len(phoneBook) == 0 {
				fmt.Println("Phonebook is empty.")
			} else {
				for name, number := range phoneBook {
					fmt.Printf("%s -> %s\n", name, number.Number)
				}
			}

		case "exit":
			fmt.Println("Exiting phonebook...")
			return

		default:
			fmt.Printf("unsupported \n")
		}

	}
}

var (
	warnLogger = log.New(os.Stdout, "WARN: ", log.LstdFlags)
	infoLogger = log.New(os.Stdout, "INFO: ", log.LstdFlags)
)

func Info(message string) {
	infoLogger.Println(message)
}

func Warn(err error, messages ...string) {
	if err == nil {
		return
	}

	message := "An error"

	if len(messages) > 0 {
		message = fmt.Sprintf("%s: %s", message, strings.Join(messages, " "))

	}

	warnLogger.Printf("%s: %v", message, err)
}
