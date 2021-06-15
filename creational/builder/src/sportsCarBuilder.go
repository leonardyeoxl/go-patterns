package main

type sportsCarBuilder struct {
    windowType string
    doorType   string
}

func newSportsCarBuilder() *sportsCarBuilder {
    return &sportsCarBuilder{}
}

func (b *sportsCarBuilder) setWindowType() {
    b.windowType = "Sports Window"
}

func (b *sportsCarBuilder) setDoorType() {
    b.doorType = "Sports Door"
}

func (b *sportsCarBuilder) getCar() car {
    return car{
        doorType:   b.doorType,
        windowType: b.windowType,
    }
}