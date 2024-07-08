package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"phonebook/book"
	"phonebook/logger"
	"strings"
	"time"
)

const phoneBookFile = "phonebook.json"

func main() {
	phoneBook, err := loadPhoneBook()
	if err != nil {
		log.Fatalf("Failed to load phonebook: %v\n", err)
	}

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
		var args []string
		if len(parts) > 1 {
			args = parts[1:]
		} else {
			args = []string{}
		}

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
			if err := savePhoneBook(phoneBook); err != nil {
				logger.Warn(err, "Failed to save phonebook")
			}
			return

		default:
			logger.Warn(errors.New("unsupported command"))
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
		return errors.New("invalid format, use: add name=number")
	}

	name, number := kv[0], kv[1]
	err := phoneBook.Add(name, book.PhoneNumber{
		Number:        number,
		LastUpdatedAt: time.Now().Unix(),
	})
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("added an entry: %s -> %s", name, number))

	return nil
}

func doGet(args []string, phoneBook book.PhoneBook) error {
	if len(args) < 1 {
		return errors.New("missing params, use: get name")
	}
	name := args[0]
	numberData, err := phoneBook.Get(name)
	if err != nil {
		return err
	}

	unixUpAt := time.Unix(numberData.LastUpdatedAt, 0)

	logger.Info(
		fmt.Sprintf(
			"Number for %s is %s, updated at %s",
			name,
			numberData.Number,
			unixUpAt.Format("2006-01-02 15:04:05"),
		),
	)

	return nil
}

func doList(_ []string, phoneBook book.PhoneBook) error {
	if len(phoneBook) == 0 {
		return errors.New("phonebook is empty")
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
		return errors.New("missing params, use: update name=newNumber")
	}

	kv := strings.SplitN(args[0], "=", 2)

	if len(kv) != 2 {
		return errors.New("invalid format, use: update name=newNumber")
	}

	name, newNumber := kv[0], kv[1]

	err := phoneBook.Update(name, book.PhoneNumber{
		Number:        newNumber,
		LastUpdatedAt: time.Now().Unix(),
	})
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("updated %s -> %s", name, newNumber))

	return nil
}

func doDelete(args []string, phoneBook book.PhoneBook) error {
	if len(args) < 1 {
		return errors.New("missing params, use: delete name")
	}

	name := args[0]

	err := phoneBook.Delete(name)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("deleted %s", name))

	return nil
}

func savePhoneBook(phoneBook book.PhoneBook) error {
	data, err := json.MarshalIndent(phoneBook, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(phoneBookFile, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func loadPhoneBook() (book.PhoneBook, error) {
	phoneBook := make(book.PhoneBook)

	data, err := ioutil.ReadFile(phoneBookFile)
	if err != nil {
		if os.IsNotExist(err) {
			return phoneBook, nil
		}
		return nil, err
	}

	err = json.Unmarshal(data, &phoneBook)
	if err != nil {
		return nil, err
	}

	return phoneBook, nil
}
