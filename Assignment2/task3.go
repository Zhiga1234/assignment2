package main

import (
    "fmt"
)

type Pizza interface {
    GetPrice() int
}

type VeggiePizza struct{}

func (p *VeggiePizza) GetPrice() int {
    return 15
}

type ToppingDecorator struct {
    pizza Pizza
}

func (t *ToppingDecorator) GetPrice() int {
    pizzaPrice := t.pizza.GetPrice()
    return pizzaPrice + t.GetToppingPrice()
}

func (t *ToppingDecorator) GetToppingPrice() int {
    // Implement this method in each specific topping decorator.
    return 0
}

type CheeseTopping struct {
    *ToppingDecorator
}

func (c *CheeseTopping) GetToppingPrice() int {
    return 10
}

type TomatoTopping struct {
    *ToppingDecorator
}

func (t *TomatoTopping) GetToppingPrice() int {
    return 7
}

func main() {
    // Create a VeggiePizza.
    pizza := &VeggiePizza{}

    // Add a cheese topping to the pizza.
    pizzaWithCheese := &CheeseTopping{
        &ToppingDecorator{
            pizza: pizza,
        },
    }

    // Add a tomato topping to the pizza with cheese.
    pizzaWithTomatoAndCheese := &TomatoTopping{
        &ToppingDecorator{
            pizza: pizzaWithCheese,
        },
    }

    // Get the price of the pizza with cheese and tomato toppings.
    price := pizzaWithTomatoAndCheese.GetPrice()

    // Print the price of the pizza with toppings.
    fmt.Printf("Price of VeggiePizza with tomato and cheese toppings is %d\n", price)
}