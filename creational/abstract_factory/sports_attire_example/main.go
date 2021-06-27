package main

import "fmt"

func main() {
    brandAFactory, _ := getSportsAttireFactory("brandA")
    brandBFactory, _ := getSportsAttireFactory("brandB")
    brandBShoe := brandBFactory.makeShoe()
    brandBShort := brandBFactory.makeShort()
    brandAShoe := brandAFactory.makeShoe()
    brandAShort := brandAFactory.makeShort()
    printShoeDetails(brandBShoe)
    printShortDetails(brandBShort)
    printShoeDetails(brandAShoe)
    printShortDetails(brandAShort)
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