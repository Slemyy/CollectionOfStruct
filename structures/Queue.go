package structures

import (
	"CollectionOfStruct/dbms"
	"errors"
	"os"
	"strings"
)

type Queue struct {
	head *Node
	tail *Node
}

// Enqueue Добавление значения в очередь.
func (queue *Queue) Enqueue(value string) {
	node := &Node{data: value}
	if queue.head == nil {
		queue.head = node
		queue.tail = node
	} else {
		queue.tail.next = node
		queue.tail = node
	}
}

// Dequeue Удаление значения из очереди.
func (queue *Queue) Dequeue() (string, error) {
	if queue.head == nil {
		return "", errors.New("queue is empty")
	} else {
		value := queue.head.data
		queue.head = queue.head.next
		if queue.head == nil {
			queue.tail = nil
		}

		return value, nil
	}
}

func (queue *Queue) ReadFromFile(filename string) error {
	content, err := os.ReadFile(filename)

	if err != nil && !os.IsNotExist(err) {
		return err
	} else if os.IsNotExist(err) {
		return nil
	}

	lines := strings.Split(string(content), "\n")
	for i := 0; i <= len(lines)-1; i++ {
		line := lines[i]
		if line != "" {
			queue.Enqueue(line)
		}
	}

	return nil
}

func (queue *Queue) WriteToFile(dataFile string, queueFile string) error {
	file, err := os.Create(queueFile)
	if err != nil {
		return err
	}
	defer file.Close()

	for {
		remove, er := queue.Dequeue()
		if er != nil {
			break
		} else {
			_, err = file.WriteString(remove + "\n")
			if err != nil {
				return err
			}
		}
	}

	err = dbms.SaveFileToDB(dataFile, queueFile, "-queue")
	if err != nil {
		return err
	}

	return nil
}
