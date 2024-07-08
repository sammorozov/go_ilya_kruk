package book

import (
	"fmt"
)

type PhoneBook map[string]PhoneNumber

type PhoneNumber struct {
	Number        string
	LastUpdatedAt int64
}

func (book *PhoneBook) Add(name string, phoneNum PhoneNumber) error {
	if _, exists := (*book)[name]; exists {
		return fmt.Errorf("name %s already exists", name)
	}

	(*book)[name] = phoneNum

	return nil
}

func (book *PhoneBook) Get(name string) (PhoneNumber, error) {
	if numberData, exists := (*book)[name]; exists {
		return numberData, nil
	}

	return PhoneNumber{}, fmt.Errorf("no entry found for %s", name)
}

func (book *PhoneBook) Update(name string, newPhoneNum PhoneNumber) error {
	if _, exists := (*book)[name]; !exists {
		return fmt.Errorf("name %s does not exist", name)
	}

	(*book)[name] = newPhoneNum

	return nil
}

func (book *PhoneBook) Delete(name string) error {
	if _, exists := (*book)[name]; !exists {
		return fmt.Errorf("name %s does not exist", name)
	}

	delete(*book, name)

	return nil
}
