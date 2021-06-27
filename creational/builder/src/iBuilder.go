package main

type iBuilder interface {
    setWindowType()
    setDoorType()
    getCar() car
}

func getBuilder(builderType string) iBuilder {
    if builderType == "normal" {
        return &normalCarBuilder{}
    }
    if builderType == "sports" {
        return &sportsCarBuilder{}
    }
    return nil
}