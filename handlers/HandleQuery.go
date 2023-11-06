package handlers

import (
	"CollectionOfStruct/structures"
	"errors"
	"strings"
	"sync"
)

func HandleQuery(file string, query string) (string, error) {
	request := strings.Fields(query)
	var mut sync.Mutex

	mut.Lock()
	defer mut.Unlock()

	switch strings.ToLower(request[0]) {
	case "spush":
		if len(request) != 3 {
			return "", errors.New("invalid request")
		}

		stack := &structures.Stack{}
		if err := stack.ReadFromFile(request[1]); err != nil {
			return "", err
		}

		stack.Push(request[2])
		if err := stack.WriteToFile(file, request[1]); err != nil {
			return "", err
		}

		return "", nil

	case "spop":
		if len(request) != 2 {
			return "", errors.New("invalid request")
		}

		stack := &structures.Stack{}
		if err := stack.ReadFromFile(request[1]); err != nil {
			return "", err
		}

		pop, err := stack.Pop()
		if err != nil {
			return "", err
		}

		if pop == "" {
			return "", errors.New("stack is empty")
		}

		if err := stack.WriteToFile(file, request[1]); err != nil {
			return "", err
		}

		return pop, nil

	case "qpush":
		if len(request) != 3 {
			return "", errors.New("invalid request")
		}

		queue := &structures.Queue{}
		if err := queue.ReadFromFile(request[1]); err != nil {
			return "", err
		}

		queue.Enqueue(request[2])
		if err := queue.WriteToFile(file, request[1]); err != nil {
			return "", err
		}

		return "", nil

	case "qpop":
		if len(request) != 2 {
			return "", errors.New("invalid request")
		}

		queue := &structures.Queue{}
		if err := queue.ReadFromFile(request[1]); err != nil {
			return "", err
		}

		dequeue, err := queue.Dequeue()
		if err != nil {
			return "", err
		}

		if dequeue == "" {
			return "", errors.New("dequeue is empty")
		}

		if err := queue.WriteToFile(file, request[1]); err != nil {
			return "", err
		}

		return dequeue, nil

	case "hadd":
		if len(request) != 4 {
			return "", errors.New("invalid request")
		}

		hash := &structures.HashTable{}
		if err := hash.ReadFromFile(request[1]); err != nil {
			return "", err
		}

		err := hash.Insert(request[2], request[3])
		if err != nil {
			return "", err
		}

		err = hash.WriteToFile(file, request[1])
		if err != nil {
			return "", err
		}

		return "", nil

	case "hrem":
		if len(request) != 3 {
			return "", errors.New("invalid request")
		}

		hash := &structures.HashTable{}
		if err := hash.ReadFromFile(request[1]); err != nil {
			return "", err
		}

		remove, err := hash.Remove(request[2])
		if err != nil {
			return "", err
		}

		err = hash.WriteToFile(file, request[1])
		if err != nil {
			return "", err
		}

		return remove, nil

	case "hget":
		if len(request) != 3 {
			return "", errors.New("invalid request")
		}

		hash := &structures.HashTable{}
		if err := hash.ReadFromFile(request[1]); err != nil {
			return "", err
		}

		get, err := hash.Get(request[2])
		if err != nil {
			return "", err
		}

		err = hash.WriteToFile(file, request[1])
		if err != nil {
			return "", err
		}

		return get, nil

	case "sadd":
		if len(request) != 3 {
			return "", errors.New("invalid request")
		}

		set := &structures.Set{}
		if err := set.ReadFromFile(request[1]); err != nil {
			return "", err
		}

		err := set.Add(request[2])
		if err != nil {
			return "", err
		}

		err = set.WriteToFile(file, request[1])
		if err != nil {
			return "", err
		}

		return "", nil

	case "srem":
		if len(request) != 3 {
			return "", errors.New("invalid request")
		}

		set := &structures.Set{}
		if err := set.ReadFromFile(request[1]); err != nil {
			return "", err
		}

		remove, err := set.Remove(request[2])
		if err != nil {
			return "", err
		}

		err = set.WriteToFile(file, request[1])
		if err != nil {
			return "", err
		}

		return remove, nil

	case "sismember":
		if len(request) != 3 {
			return "", errors.New("invalid request")
		}

		set := &structures.Set{}
		if err := set.ReadFromFile(request[1]); err != nil {
			return "", err
		}

		get, err := set.Itmember(request[2])
		if err != nil {
			return "", err
		}

		err = set.WriteToFile(file, request[1])
		if err != nil {
			return "", err
		}

		return map[bool]string{true: "Value found", false: "No such meaning\n"}[get], nil

	default:
		return "", errors.New("invalid request")
	}
}
