package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	lab2 "github.com/1Laggy1/Lab2"
)

func main() {
	// Оголошення флагів
	expressionFlag := flag.String("e", "", "вираз")
	fileFlag := flag.String("f", "", "файл із виразом")
	outputFlag := flag.String("o", "", "файл результату")

	// Парсинг флагів командного рядка
	flag.Parse()

	// Розпізнавання типу вводу
	input, err := lab2.ParseInput(*expressionFlag, *fileFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Помилка визначення вводу: %v\n", err)
		os.Exit(1)
	}

	// Використання стандартного виводу, якщо файл для виведення не вказано
	var output io.Writer = os.Stdout
	if *outputFlag != "" {
		outputFile, err := os.Create(*outputFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Помилка створення файлу для виведення: %v\n", err)
			os.Exit(1)
		}
		defer outputFile.Close()
		output = outputFile
	}

	// Створення екземпляру ComputeHandler та обробка виразу
	handler := lab2.NewComputeHandler(input, output)
	err = handler.Compute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Помилка обробки виразу: %v\n", err)
		os.Exit(1)
	}
}
