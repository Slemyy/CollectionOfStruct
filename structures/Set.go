package structures

import (
	"CollectionOfStruct/dbms"
	"errors"
	"os"
	"strings"
)

type Set struct {
	table [arraySize]string
}

func (set *Set) Clear() error {
	for i := 0; i < arraySize; i++ {
		if set.table[i] != "" {
			set.table[i] = ""
		}
	}
	return nil
}

// Add Добавление значения во множество.
func (set *Set) Add(value string) error {
	index := hashFunction(value)

	if set.table[index] == "" {
		set.table[index] = value
		return nil
	}

	if set.table[index] == value {
		return errors.New("such a key already exists")
	}

	for i := (index + 1) % arraySize; i != index; i = (i + 1) % arraySize {
		if i == index {
			return errors.New("table is full")
		}

		if set.table[i] == "" {
			set.table[i] = value
			return nil
		}

		if set.table[i] == value {
			return errors.New("such a key already exists")
		}
	}

	return errors.New("failed to add element")
}

// Remove Rem Удаление значения во множестве.
func (set *Set) Remove(value string) (string, error) {
	index := hashFunction(value)

	if set.table[index] == "" {
		return "", errors.New("no such meaning")
	}

	if set.table[index] == value {
		remove := set.table[index]
		set.table[index] = ""
		return remove, nil
	}

	for i := (index + 1) % arraySize; i != index; i = (i + 1) % arraySize {
		if set.table[i] == value && set.table[i] != "" {
			remove := set.table[i]
			set.table[i] = ""
			return remove, nil
		}
	}

	return "", errors.New("failed to delete element")
}

func (set *Set) Itmember(value string) (bool, error) {
	index := hashFunction(value)

	if set.table[index] == "" {
		return false, errors.New("no such meaning")
	}

	if set.table[index] == value {
		return true, nil
	}

	for i := (index + 1) % arraySize; i != index; i = (i + 1) % arraySize {
		if set.table[index] == "" {
			return false, errors.New("no such meaning")
		}

		if set.table[i] == value && set.table[i] != "" {
			return true, nil
		}
	}

	return false, errors.New("element not found")
}

func (set *Set) ReadFromFile(filename string) error {
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
			err := set.Add(line)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (set *Set) WriteToFile(dataFile string, setFile string) error {
	file, err := os.Create(setFile)
	if err != nil {
		return err
	}
	defer file.Close()

	for i := 0; i < arraySize; i++ {
		value := set.table[i]
		if value != "" {
			_, err = file.WriteString(value + "\n")
			if err != nil {
				return err
			}
		}
	}

	set.Clear()

	err = dbms.SaveFileToDB(dataFile, setFile, "-set")
	if err != nil {
		return err
	}

	return nil
}
