package structures

import "errors"

type Stack struct {
	head *Node
}

// Push Добавление значения в стэк
func (stack *Stack) Push(value string) {
	node := &Node{data: value}

	if stack.head == nil {
		stack.head = node
	} else {
		node.next = stack.head
		stack.head = node
	}
}

// Pop Удаление значения из стэка
func (stack *Stack) Pop() (string, error) {
	if stack.head == nil {
		return "", errors.New("stack is empty")
	}

	value := stack.head.data
	stack.head = stack.head.next

	return value, nil
}

func (stack *Stack) ReadFromFile(filename string) {

}

func (stack *Stack) WriteToFile(filename string) {

}
