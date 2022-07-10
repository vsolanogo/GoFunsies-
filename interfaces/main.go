package main

import (
	"errors"
	"fmt"
	"reflect"
)

type liveStockOperations interface {
	calculateWeightOfFood() float64
	getName() string
	checkWhetherTypeOfAnimalEqualsItsName() error
	checkWhetherWeightBelowNormal() error
	checkWhetherAnimalIsEdible() error
}

type WithAnimal struct {
	name       string
	weight     float64
	isEdible   bool
	animalType string
}

func (w WithAnimal) getName() string {
	return w.name
}

func (w WithAnimal) getWeight() float64 {
	return w.weight
}

func (w WithAnimal) getType() string {
	return w.animalType
}

func (w WithAnimal) checkWhetherTypeOfAnimalEqualsItsName() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("in printLivestockInfoAndGetTotalFoodConsumtion: %w", err)
		}
	}()

	if w.getType() != w.getName() {
		return fmt.Errorf("%s named %s: failed validation: %s is %s: the type of animal has to be %s", w.getType(), w.getName(), w.getType(), w.getName(), w.getType())
	}
	return nil
}

type Dog struct {
	WithAnimal
}

func (d Dog) getType() string {
	const typeName = "Dog"
	return typeName
}

func (d Dog) checkWhetherAnimalIsEdible() error {
	const isEdible = false

	if d.isEdible == isEdible {
		return nil
	}

	return fmt.Errorf("Dog named %s: failed validation: dog is not edible: dog has to be not edible", d.name)
}

func (d Dog) checkWhetherWeightBelowNormal() error {
	const defaultLowerRate = 3.17

	if d.weight >= defaultLowerRate {
		return nil
	}

	return fmt.Errorf("Dog weight %f: failed validation: dog weight below normal rate: dog normal weight %f", d.weight, defaultLowerRate)
}

func (d Dog) calculateWeightOfFood() float64 {
	const defaultConsumption = 2
	return d.weight * float64(defaultConsumption)
}

type Cat struct {
	WithAnimal
}

func (c Cat) getType() string {
	const typeName = "Cat"
	return typeName
}

func (c Cat) checkWhetherAnimalIsEdible() error {
	const isEdible = false

	if c.isEdible == isEdible {
		return nil
	}

	return errors.New("Cat is not edible")
}

func (c Cat) checkWhetherWeightBelowNormal() error {
	const defaultLowerRate = 3

	if c.weight >= defaultLowerRate {
		return nil
	}

	return fmt.Errorf("Cat weight is %f: failed validation: cat weight below normal rate: cat normal weight %d", c.weight, defaultLowerRate)
}

func (c Cat) calculateWeightOfFood() float64 {
	const defaultConsumption = 7
	return c.weight * float64(defaultConsumption)
}

type Cow struct {
	WithAnimal
}

func (c Cow) getType() string {
	const typeName = "Cow"
	return typeName
}

func (c Cow) checkWhetherAnimalIsEdible() error {
	const isEdible = true

	if c.isEdible == isEdible {
		return nil
	}

	return errors.New("Cow is edible")
}

func (c Cow) calculateWeightOfFood() float64 {
	const defaultConsumption = 25
	return c.weight * float64(defaultConsumption)
}

func (c Cow) checkWhetherWeightBelowNormal() error {
	const defaultLowerRate = 500

	if c.weight >= defaultLowerRate {
		return nil
	}

	return fmt.Errorf("Cow weight is %f: failed validation: cow weight below normal rate: cow normal weight %d", c.weight, defaultLowerRate)
}

func printLivestockInfoAndGetTotalFoodConsumtion(list []liveStockOperations) (_ float64, err error) {
	for _, animal := range list {
		err = animal.checkWhetherTypeOfAnimalEqualsItsName()
		if err != nil {
			return 0, err
		}
	}

	for _, animal := range list {
		err = animal.checkWhetherWeightBelowNormal()
		if err != nil {
			return 0, err
		}
	}

	for _, animal := range list {
		err = animal.checkWhetherAnimalIsEdible()
		if err != nil {
			return 0, err
		}
	}

	var totalFoodNeeded float64
	for _, animal := range list {
		foodNeeded := animal.calculateWeightOfFood()
		totalFoodNeeded = totalFoodNeeded + foodNeeded
		fmt.Println(fmt.Sprintf("%s named %s eats %v kg", reflect.TypeOf(animal).Name(), animal.getName(), animal.calculateWeightOfFood()))
	}

	return totalFoodNeeded, nil
}

func main() {
	liveStock := []liveStockOperations{
		Cow{WithAnimal{name: "cow1", weight: 80.1, isEdible: false, animalType: "Cow"}},
		Dog{WithAnimal{name: "doggy", weight: 7, isEdible: false, animalType: "Dog"}},
		Cat{WithAnimal{name: "cat2", weight: 4, isEdible: false, animalType: "Cat"}},
		Dog{WithAnimal{name: "dog2", weight: 6.66, isEdible: false, animalType: "Dog"}},
		Cow{WithAnimal{name: "cow2", weight: 166, isEdible: true, animalType: "Cow"}},
		Cat{WithAnimal{name: "kitty", weight: 3, isEdible: false, animalType: "Cat"}},
	}

	res, err := printLivestockInfoAndGetTotalFoodConsumtion(liveStock)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Total food needed for all animals: " + fmt.Sprintf("%v", res))
	}
}
