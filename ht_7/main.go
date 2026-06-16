package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Cache struct {
	mutex sync.RWMutex
	data  map[int]string
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[int]string, 5),
	}
}

func (c *Cache) Set(i int, str string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[i] = str

	fmt.Printf("###Запись###\tДля %d установлено значение: %s \n\n", i, str)
}

func (c *Cache) Get(i int) (string, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	str, ok := c.data[i]
	if !ok {
		return "", errors.New("Кэшированные данные отсутствуют.")
	} else {
		return str, nil
	}
}

func main() {
	c := NewCache()
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		for range 5 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				i := rand.Intn(5)
				str := "Кэш для запроса: " + strconv.Itoa(i*rand.Intn(3))
				c.Set(i, str)
				time.Sleep(time.Millisecond * 300)
			}()
		}
	}()

	go func() {
		defer wg.Done()
		for range 3 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := range 5 {
					str, err := c.Get(i)
					if err != nil {
						fmt.Printf("Для %d\n Ошибка: %v\n\n", i, err)
					} else {
						fmt.Printf("Для %d\n Значение: %v\n\n", i, str)
					}
				}
				time.Sleep(time.Millisecond * 500)
			}()
		}
	}()

	wg.Wait()
	fmt.Printf("%v\n", c.data)
}
