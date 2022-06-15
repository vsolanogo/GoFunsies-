package main

import (
	"fmt"
)

type Calculation interface {
	getCost() float64
}

func getTotal(list []Calculation) (total float64) {
	for _, i := range list {
		total += i.getCost()
	}
	return
}

type Operation interface {
	quantityAbleToBuy() int
}

func quantityAbleToBuy(money float64, prod Product) (quantity int) {
	return int(money/prod.getCost())
}

func areWeAbleToBuyTheseProds(money float64, list []Calculation) (isAbleToBuy bool) {
	return int(money/getTotal(list)) != 0
}

type Product struct {
	name	string
	price	float64
	quantity	float64
}

func (p Product) getCost() float64 {
	return p.price * p.quantity
}

const applePrice float64 = 5.99
const pearPrice float64 = 7
const myBank float64 = 23

func main() {
	fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")

	products := []Calculation {
		Product{"apple", applePrice, 9},
		Product{"pear", pearPrice, 8},
	}
	
	fmt.Println(getTotal(products))

	fmt.Println("2. Скільки груш ми можемо купити?")

	fmt.Println(quantityAbleToBuy(myBank, Product{"pear", pearPrice, 1} ))

	fmt.Println("3. Скільки яблук ми можемо купити?")

	fmt.Println(quantityAbleToBuy(myBank, Product{"apple", applePrice, 1} ))

	fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука?")

	products4 := []Calculation {
		Product{"apple", applePrice, 2},
		Product{"pear", pearPrice, 2},
	}

	fmt.Println(areWeAbleToBuyTheseProds(myBank, products4))
}