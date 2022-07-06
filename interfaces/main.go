package main

import (
	"fmt"
	"reflect"
)

type liveStockOperation interface {
	calculateWeightOfFood() float64
	getName() string
}

type WithName struct {
	name string
}

func (w WithName) getName() string {
	return w.name
}

type WithWeight struct {
	weight float64
}

func (w WithWeight) getWeight() float64 {
	return w.weight
}

type Dog struct {
	WithName
	WithWeight
}

func (d Dog) calculateWeightOfFood() float64 {
	defaultConsumption := 2
	return d.weight * float64(defaultConsumption)
}

type Cat struct {
	WithName
	WithWeight
}

func (c Cat) calculateWeightOfFood() float64 {
	defaultConsumption := 7
	return c.weight * float64(defaultConsumption)
}

type Cow struct {
	WithName
	WithWeight
}

func (c Cow) calculateWeightOfFood() float64 {
	defaultConsumption := 25
	return c.weight * float64(defaultConsumption)
}

func livestockInfo(list []liveStockOperation) float64 {
	var totalFoodNeeded float64
	for _, animal := range list {
		foodNeeded := animal.calculateWeightOfFood()
		totalFoodNeeded = totalFoodNeeded + foodNeeded
		fmt.Println(fmt.Sprintf("%s named %s eats %v kg", reflect.TypeOf(animal).Name(), animal.getName(), animal.calculateWeightOfFood()))
	}

	return totalFoodNeeded
}

func main() {
	liveStock := []liveStockOperation{
		Cow{WithName{name: "cow1"}, WithWeight{weight: 80.1}},
		Dog{WithName{name: "doggy"}, WithWeight{weight: 7}},
		Cat{WithName{name: "cat2"}, WithWeight{weight: 4}},
		Dog{WithName{name: "dog2"}, WithWeight{weight: 6.66}},
		Cow{WithName{name: "cow2"}, WithWeight{weight: 166}},
		Cat{WithName{name: "kitty"}, WithWeight{weight: 3}},
	}

	res := livestockInfo(liveStock)

	fmt.Println("Food needed for all animals: " + fmt.Sprintf("%v", res))
}
