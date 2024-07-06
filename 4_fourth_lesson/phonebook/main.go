package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"phonebook/book"
	"phonebook/logger"
	"strings"
	"time"
)

func main() {
	phoneBook := make(book.PhoneBook)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Phonebook")
	fmt.Println("Available commands: add, get, delete, update, list, exit")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		command := parts[0]
		args := parts[1:]

		switch command {
		case "add":
			handleCommand(doAdd, args, phoneBook)

		case "get":
			handleCommand(doGet, args, phoneBook)

		case "delete":
			handleCommand(doDelete, args, phoneBook)

		case "update":
			handleCommand(doUpdate, args, phoneBook)

		case "list":
			handleCommand(doList, args, phoneBook)

		case "exit":
			logger.Info("Exiting phonebook...")
			return

		default:
			logger.Warn(errors.New("unsupported \n"))
		}
	}
}

func handleCommand(cmd func([]string, book.PhoneBook) error, args []string, phoneBook book.PhoneBook) {
	if err := cmd(args, phoneBook); err != nil {
		logger.Warn(err, "cmd failed")
	}
}

func doAdd(args []string, phoneBook book.PhoneBook) error {
	if len(args) < 1 {
		return errors.New("missing params")
	}

	kv := strings.SplitN(args[0], "=", 2)
	if len(kv) != 2 {
		return errors.New("invalid format")
	}

	name, number := kv[0], kv[1]
	err := phoneBook.Add(name, number)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("added an entry: %s -> %s\n", name, number))

	return nil
}

func doGet(args []string, phoneBook book.PhoneBook) error {
	if len(args) < 1 {
		return errors.New("missing params")
	}
	name := args[0]
	numberData, err := phoneBook.Get(name)
	if err != nil {
		return err
	}

	unixUpAt := time.Unix(numberData.LastUpdatedAt, 0)

	logger.Info(
		fmt.Sprintf(
			"Num for %s is %s updated at %s\n",
			name,
			numberData.Number,
			unixUpAt.Format("2006-01-02 15:04:05"),
		),
	)

	return nil
}

func doList(_ []string, phoneBook book.PhoneBook) error {
	if len(phoneBook) == 0 {
		return errors.New("empty")
	} else {
		results := ""

		for name, number := range phoneBook {
			results += fmt.Sprintf("%s -> %s\n", name, number.Number)
		}

		logger.Info(results)
	}
	return nil
}

func doUpdate(args []string, phoneBook book.PhoneBook) error {
	if len(args) < 1 {
		return errors.New("misss")
	}

	kv := strings.SplitN(args[0], "=", 2)

	if len(kv) != 2 {
		return errors.New("invalid")
	}

	name, newNumber := kv[0], kv[1]

	err := phoneBook.Update(name, newNumber)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("updated"))

	return nil
}

func doDelete(args []string, phoneBook book.PhoneBook) error {
	if len(args) < 1 {
		return errors.New("miss")
	}

	name := args[0]

	err := phoneBook.Delete(name)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("deleted"))

	return nil
}
