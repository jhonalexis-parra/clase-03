package main

import (
	"errors"
	"fmt"
)

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func main() {
	animalPerro, err := animal("perro")
	if err != nil {
		fmt.Println(err.Error())
	}

	animalGato, err := animal("gato")
	if err != nil {
		fmt.Println(err.Error())
	}

	animalHamnster, err := animal("hamster")
	if err != nil {
		fmt.Println(err.Error())
	}

	animalTarantula, err := animal("tarantula")
	if err != nil {
		fmt.Println(err.Error())
	}

	var amount float32
	amount += animalPerro(10)
	amount += animalGato(10)
	amount += animalHamnster(10)
	amount += animalTarantula(10)
	fmt.Println(amount)

	_, err = animal("canario")
	if err != nil {
		fmt.Println(err.Error())
	}

}

func animal(animal string) (func(int) float32, error) {
	switch animal {
	case perro:
		return animalPerro, nil
	case gato:
		return animalGato, nil
	case hamster:
		return animalHamnster, nil
	case tarantula:
		return animalTarantula, nil
	default:
		return nil, errors.New("No se encuentra animal")
	}
}

func animalPerro(cantidad int) float32 {
	return float32(10 * 1000 * cantidad)
}

func animalGato(cantidad int) float32 {
	return float32(5 * 1000 * cantidad)
}

func animalHamnster(cantidad int) float32 {
	return float32(250 * cantidad)
}

func animalTarantula(cantidad int) float32 {
	return float32(150 * cantidad)
}
