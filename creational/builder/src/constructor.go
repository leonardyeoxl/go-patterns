package main

type constructor struct {
    builder iBuilder
}

func newConstructor(b iBuilder) *constructor {
    return &constructor{
        builder: b,
    }
}

func (c *constructor) setBuilder(b iBuilder) {
    c.builder = b
}

func (c *constructor) buildCar() car {
    c.builder.setDoorType()
    c.builder.setWindowType()
    return c.builder.getCar()
}