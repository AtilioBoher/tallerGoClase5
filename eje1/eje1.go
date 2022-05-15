package main

import "fmt"

func multiply(x, y int) int {
	return x * y
}

func min(nums ...int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

type animal struct {
	nombre string
	edad   int
}

type perro struct {
	animal
	dueño string
}

type paloma struct {
	animal
	alturaMax int
}

func genPerro(nombre1, dueño1 string, edad1 int) *perro {
	perro1 := perro{
		animal: animal{
			nombre: nombre1,
			edad:   edad1,
		},
		dueño: dueño1,
	}
	return &perro1
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(multiply(3, 4))
	fmt.Println(min(1, 2, 4, 8, -2, -10, 58, -50))

	perro1 := genPerro("Firulais", "Roberto", 4)
	fmt.Println("nombre del perro:", *&perro1.nombre)
	fmt.Println("edad del perro:", *&perro1.edad, "años")
	fmt.Println("dueño del perro:", *&perro1.dueño)
	fmt.Println("factirial:", fact(5))
}
