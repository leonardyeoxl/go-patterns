package main

import "fmt"

func main() {
    adidasFactory, _ := getSportsAttireFactory("adidas")
    nikeFactory, _ := getSportsAttireFactory("nike")
    nikeShoe := nikeFactory.makeShoe()
    nikeShort := nikeFactory.makeShort()
    adidasShoe := adidasFactory.makeShoe()
    adidasShort := adidasFactory.makeShort()
    printShoeDetails(nikeShoe)
    printShortDetails(nikeShort)
    printShoeDetails(adidasShoe)
    printShortDetails(adidasShort)
}

func printShoeDetails(s iShoe) {
    fmt.Printf("Logo: %s", s.getLogo())
    fmt.Println()
    fmt.Printf("Size: %d", s.getSize())
    fmt.Println()
}

func printShortDetails(s iShort) {
    fmt.Printf("Logo: %s", s.getLogo())
    fmt.Println()
    fmt.Printf("Size: %d", s.getSize())
    fmt.Println()
}