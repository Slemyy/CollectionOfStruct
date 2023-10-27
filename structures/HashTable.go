package structures

import (
	"CollectionOfStruct/dbms"
	"errors"
	"os"
	"strings"
)

type HashTable struct {
	table [arraySize]*HashNode
}

func hashFunction(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash += int(key[i])
	}

	return hash % arraySize
}

func (hashMap *HashTable) Clear() error {
	for i := 0; i < arraySize; i++ {
		if hashMap.table[i] != nil {
			hashMap.table[i] = nil
		}
	}
	return nil
}

// Insert Добавление значения в Хэш-таблицу.
func (hashMap *HashTable) Insert(key string, value string) error {
	keyValue := &HashNode{key: key, value: value}
	index := hashFunction(keyValue.key)

	if hashMap.table[index] == nil {
		hashMap.table[index] = keyValue
		return nil
	}

	if hashMap.table[index].key == key {
		return errors.New("such a key already exists")
	}

	for i := (index + 1) % arraySize; i != index; i = (i + 1) % arraySize {
		if i == index {
			return errors.New("table is full")
		}

		if hashMap.table[i] == nil {
			hashMap.table[i] = keyValue
			return nil
		}

		if hashMap.table[i].key == key {
			return errors.New("such a key already exists")
		}
	}

	return errors.New("failed to add element")
}

// Remove Удаление значения в Хэш-таблицы.
func (hashMap *HashTable) Remove(key string) (string, error) {
	index := hashFunction(key)

	if hashMap.table[index] == nil {
		return "", errors.New("no such meaning")
	}

	// Если найден ключ в текущем индексе, удаляем его
	if hashMap.table[index].key == key {
		remove := hashMap.table[index].value
		hashMap.table[index] = nil
		return remove, nil
	}

	// В случае коллизии, ищем ключ в следующих индексах
	for i := (index + 1) % arraySize; i != index; i = (i + 1) % arraySize {
		if hashMap.table[i] != nil && hashMap.table[i].key == key {
			remove := hashMap.table[i].value
			hashMap.table[i] = nil
			return remove, nil
		}
	}

	return "", errors.New("no such meaning")
}

// Get Получение элемента хэш-таблицы по индексу.
func (hashMap *HashTable) Get(key string) (string, error) {
	index := hashFunction(key)

	if hashMap.table[index] == nil {
		return "", errors.New("element not found")
	} else if hashMap.table[index].key == key {
		return hashMap.table[index].value, nil
	} else {
		for i := (index + 1) % arraySize; i != index; i = (i + 1) % arraySize {
			if hashMap.table[i] == nil {
				return "", errors.New("element not found")
			}

			if hashMap.table[i].key == key {
				return hashMap.table[i].value, nil
			}
		}
	}

	return "", errors.New("element not found")
}

func (hashMap *HashTable) ReadFromFile(filename string) error {
	content, err := os.ReadFile(filename)

	if err != nil && !os.IsNotExist(err) {
		return err
	} else if os.IsNotExist(err) {
		return nil
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) >= 2 {
			key := parts[0]
			value := strings.Join(parts[1:], " ")
			err := hashMap.Insert(key, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (hashMap *HashTable) WriteToFile(dataFile string, hashFile string) error {
	file, err := os.Create(hashFile)
	if err != nil {
		return err
	}
	defer file.Close()

	for i := 0; i < arraySize; i++ {
		if hashMap.table[i] != nil {
			_, err = file.WriteString(hashMap.table[i].key + " " + hashMap.table[i].value + "\n")
			if err != nil {
				return err
			}
		}
	}

	hashMap.Clear()

	err = dbms.SaveFileToDB(dataFile, hashFile, "-hash")
	if err != nil {
		return err
	}

	return nil
}
