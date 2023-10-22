package dbms

import (
	"CollectionOfStruct/structures"
	"errors"
	"fmt"
	"strings"
)

func HandleQuery(file string, query string) error {
	request := strings.Fields(query)

	switch strings.ToLower(request[0]) {
	case "spush":
		if len(request) != 3 {
			return errors.New("invalid request")
		}

		stack := &structures.Stack{}
		stack.Push(request[2])
		return nil

	case "spop":
		if len(request) != 2 {
			return errors.New("invalid request")
		}

		stack := &structures.Stack{}
		pop, err := stack.Pop()
		if err != nil {
			return err
		}

		fmt.Println(pop) // Выводим элемент, который удалили
		return nil

	default:
		return errors.New("invalid request")
	}
}
