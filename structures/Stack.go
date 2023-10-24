package structures

import (
	"CollectionOfStruct/dbms"
	"errors"
	"os"
	"strings"
)

type Stack struct {
	head *Node
}

// Push Добавление значения в стэк.
func (stack *Stack) Push(value string) {
	node := &Node{data: value}

	if stack.head == nil {
		stack.head = node
	} else {
		node.next = stack.head
		stack.head = node
	}
}

// Pop Удаление значения из стэка.
func (stack *Stack) Pop() (string, error) {
	if stack.head == nil {
		return "", errors.New("stack is empty")
	}

	value := stack.head.data
	stack.head = stack.head.next

	return value, nil
}

func (stack *Stack) ReadFromFile(filename string) error {
	content, err := os.ReadFile(filename)

	if err != nil && !os.IsNotExist(err) {
		return err
	} else if os.IsNotExist(err) {
		return nil
	}

	lines := strings.Split(string(content), "\n")
	for i := len(lines) - 1; i >= 0; i-- {
		if lines[i] == "" {
			continue
		}

		line := lines[i]
		stack.Push(line)
	}

	return nil
}

func (stack *Stack) WriteToFile(dataFile string, stackFile string) error {
	file, err := os.Create(stackFile)
	if err != nil {
		return err
	}
	defer file.Close()

	for {
		remove, er := stack.Pop()
		if er != nil {
			break
		} else {
			_, err = file.WriteString(remove + "\n")
			if err != nil {
				return err
			}
		}
	}

	err = dbms.SaveFileToDB(dataFile, stackFile, "-stack")
	if err != nil {
		return err
	}

	return nil
}
