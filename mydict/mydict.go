package mydict

import (
	"errors"
	"fmt"
)

type Dictionary map[string]string

var errNotFound = errors.New("not found")
var errCantUpdate = errors.New("Can't Update")
var errAlreadyExist = errors.New("This word is already existed")

//Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exist := d[word]
	if exist {
		return value, nil
	}

	return "", errNotFound
}

// Add word
func (d Dictionary) Add(word, value string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = value
	case nil:
		return errAlreadyExist
	}

	return nil
}

// Update word
func (d Dictionary) Update(word, value string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = value
	case errNotFound:
		return errCantUpdate
	}

	return nil
}

// Delete word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	if err == errNotFound {
		return errNotFound
	} else if err == nil {
		delete(d, word)
		fmt.Println("데이터를 삭제했습니다")
	}

	return nil
}
