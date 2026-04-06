package main

import (
	"errors"
	"fmt"
	"os"
)

// Создайте интерфейс Writer с методом Write([]byte) (int, error), представляющим запись данных. Затем создайте структуру File и реализуйте этот интерфейс для записи данных в файл. Также создайте структуру Console и реализуйте интерфейс для вывода данных в консоль. Используйте этот интерфейс для записи текста как в файл, так и в консоль.

type Writer interface {
	Write([]byte) (int, error)
}

func WriteData(w Writer, b []byte) (int, error) {
	i, err := w.Write(b)
	return i, err
}

type File struct {
	Name string
}

func (f *File) Write(b []byte) (int, error) {
	if string(b) == "" {
		return 0, errors.New("Нет данных для записи")
	}
	file, err := os.OpenFile(f.Name, os.O_APPEND, 0666)
	if err != nil {
		// file, err = os.Create(f.Name)
		// if err != nil {
		fmt.Println("Невозможно открыть файл:", err)
		os.Exit(1)
		// }
	}
	defer file.Close()

	file.WriteString(string(b))
	// fmt.Println(string(b))

	return len(b), nil

}

type Console struct {
}

func (c *Console) Write(b []byte) (int, error) {
	if string(b) == "" {
		return 0, errors.New("Нет данных для записи")
	}
	fmt.Println(string(b))
	return len(b), nil
}
