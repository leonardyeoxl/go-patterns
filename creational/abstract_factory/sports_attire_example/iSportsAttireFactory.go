package main

import (
	"fmt"
)

type iSportsAttireFactory interface {
	makeShoe() iShoe
	makeShort() iShort
}

func getSportsAttireFactory(brand string) (iSportsAttireFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}
	if brand == "nike" {
		return &nike{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}