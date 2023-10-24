package handlers

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
		if err := stack.ReadFromFile(request[1]); err != nil {
			return err
		}

		stack.Push(request[2])
		if err := stack.WriteToFile(file, request[1]); err != nil {
			return err
		}

		return nil

	case "spop":
		if len(request) != 2 {
			return errors.New("invalid request")
		}

		stack := &structures.Stack{}
		if err := stack.ReadFromFile(request[1]); err != nil {
			return err
		}

		pop, err := stack.Pop()
		if err != nil {
			return err
		}

		if pop == "" {
			return errors.New("stack is empty")
		}

		fmt.Println(pop) // Выводим элемент, который удалили

		if err := stack.WriteToFile(file, request[1]); err != nil {
			return err
		}

		return nil

	case "qpush":
		if len(request) != 3 {
			return errors.New("invalid request")
		}

		queue := &structures.Queue{}
		if err := queue.ReadFromFile(request[1]); err != nil {
			return err
		}

		queue.Enqueue(request[2])
		if err := queue.WriteToFile(file, request[1]); err != nil {
			return err
		}

		return nil

	case "qpop":
		if len(request) != 2 {
			return errors.New("invalid request")
		}

		queue := &structures.Queue{}
		if err := queue.ReadFromFile(request[1]); err != nil {
			return err
		}

		dequeue, err := queue.Dequeue()
		if err != nil {
			return err
		}

		if dequeue == "" {
			return errors.New("dequeue is empty")
		}

		fmt.Println(dequeue) // Выводим элемент, который удалили
		if err := queue.WriteToFile(file, request[1]); err != nil {
			return err
		}

		return nil

	default:
		return errors.New("invalid request")
	}
}
