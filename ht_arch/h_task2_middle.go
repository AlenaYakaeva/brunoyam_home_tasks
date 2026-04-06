package main

import "fmt"

type Vehicle interface {
	Start()
	Stop()
}

func StartRide(v Vehicle) {
	v.Start()
}

func StopRide(v Vehicle) {
	v.Stop()
}

type Car struct {
	Brand  string
	Model  string
	Color  string
	Number string
}

func (c *Car) Start() {
	fmt.Printf("Автомобиль %s запущен\n", c.Number)
}

func (c *Car) Stop() {
	fmt.Printf("Автомобиль %s заглушен\n", c.Number)
}

type Moto struct {
	Brand  string
	Model  string
	Color  string
	Number string
}

func (c *Moto) Start() {
	fmt.Printf("Мотоцикл %s запущен\n", c.Number)
}

func (c *Moto) Stop() {
	fmt.Printf("Мотоцикл %s заглушен\n", c.Number)
}
