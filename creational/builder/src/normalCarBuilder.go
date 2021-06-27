package main

type normalCarBuilder struct {
    windowType string
    doorType   string
}

func newNormalCarBuilder() *normalCarBuilder {
    return &normalCarBuilder{}
}

func (b *normalCarBuilder) setWindowType() {
    b.windowType = "Normal Window"
}

func (b *normalCarBuilder) setDoorType() {
    b.doorType = "Normal Door"
}

func (b *normalCarBuilder) getCar() car {
    return car{
        doorType:   b.doorType,
        windowType: b.windowType,
    }
}