package main

func main() {
	packaging := &packaging{}

	// set next for assembly section
	assembly := &assembly{}
	assembly.setNext(packaging)

	material := &material{}
	material.setNext(assembly)

	task := &task{name: "truck_toy"}
	material.execute(task)
}