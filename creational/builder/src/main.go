package main

import (
    "fmt"
)

func main() {
    normalCarBuilder := getBuilder("normal")
    sportsCarBuilder := getBuilder("sports")

    constructor := newConstructor(normalCarBuilder)
    normalCar := constructor.buildCar()

    fmt.Printf("Normal Car Door Type: %s\n", normalCar.doorType)
    fmt.Printf("Normal Car Window Type: %s\n", normalCar.windowType)

    constructor.setBuilder(sportsCarBuilder)
    sportsCar := constructor.buildCar()

    fmt.Printf("\nSports Car Door Type: %s\n", sportsCar.doorType)
    fmt.Printf("Sports Car Window Type: %s\n", sportsCar.windowType)
}