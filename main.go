package main

import (
	"CollectionOfStruct/handlers"
	"flag"
	"fmt"
)

var (
	file  string
	query string
)

func main() {
	// Парсим аргументы командной строки
	flag.StringVar(&file, "file", "", "Data file")
	flag.StringVar(&query, "query", "", "Database Query")
	flag.Parse()

	// Проверка обязательных аргументов
	if file == "" {
		fmt.Println("Error: you must specify a data file (--file)")
		return
	}

	if query == "" {
		fmt.Println("Error: you must specify a database query (--query)")
		return
	}

	// Вызов функции, которая обрабатывает запрос к базе
	err := handlers.HandleQuery(file, query)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
