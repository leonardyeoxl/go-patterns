package main

import (
	"fmt"
)

type iSportsAttireFactory interface {
	makeShoe() iShoe
	makeShort() iShort
}

func getSportsAttireFactory(brand string) (iSportsAttireFactory, error) {
	if brand == "brandA" {
		return &brandA{}, nil
	}
	if brand == "brandB" {
		return &brandB{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}