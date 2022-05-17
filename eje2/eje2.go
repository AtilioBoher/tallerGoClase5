package main

import (
	"fmt"
	"sort" // https://pkg.go.dev/sort
)

const faltasMax int = 3

type curso struct {
	id          int
	nombre      string
	aula        string // aula asignada
	cantMaxEstu int    // cantidad máxima de estudiantes
	profe       profe
	alumn       []alumn
}

type profe struct {
	id     int
	nombre string
	antig  int //antigüedad
}

type alumn struct {
	id     int
	nombre string
	notas  []float64
	faltas int
}

func (a alumn) prom() float64 {
	var sum float64
	sum = 0
	for _, n := range a.notas {
		sum += n
	}
	return sum / float64(len(a.notas))
}

func (c curso) debajoDe6() []alumn {
	var alumnos []alumn
	for _, n := range c.alumn {
		if n.prom() < 6 {
			alumnos = append(alumnos, n)
		}
	}
	return alumnos
}

func (c curso) libres() []alumn {
	var alumnos []alumn
	for _, n := range c.alumn {
		if n.faltas > faltasMax {
			alumnos = append(alumnos, n)
		}
	}
	return alumnos
}

func (c curso) tresMejorPromedio() []alumn {
	alumnos := c.alumn
	sort.Sort(byMean(alumnos))
	return alumnos[0:2]
}

type byMean []alumn

func (a byMean) Len() int           { return len(a) }
func (a byMean) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byMean) Less(i, j int) bool { return a[i].prom() > a[j].prom() }

func (c curso) classroomUsage() (int, float64) {
	var porc float64
	var numOfStud int
	numOfStud = len(c.alumn)
	porc = float64(numOfStud) / float64(c.cantMaxEstu) * 100
	return numOfStud, porc
}

func main() {

	curso1 := curso{
		id:          100,
		nombre:      "curso re copado",
		aula:        "el aula más copadad",
		cantMaxEstu: 30,
		profe: profe{
			id:     20,
			nombre: "nombre del profe",
			antig:  5,
		},
		alumn: []alumn{
			{
				id:     1,
				nombre: "rolando",
				notas:  []float64{7, 10, 8, 9, 6},
				faltas: 1,
			},
			{
				id:     2,
				nombre: "jaimito",
				notas:  []float64{1, 2, 3, 4, 5},
				faltas: 10,
			},
			{
				id:     3,
				nombre: "lucas",
				notas:  []float64{1, 8.8, 3, 4, 5},
				faltas: 3,
			},
			{
				id:     4,
				nombre: "fulano",
				notas:  []float64{8, 7, 7.4, 4, 5},
				faltas: 4,
			},
		},
	}

	fmt.Println(curso1.alumn[1].prom())
	fmt.Println(curso1.debajoDe6())
	fmt.Println(curso1.libres())
	fmt.Println(curso1.tresMejorPromedio())
	fmt.Println(curso1.classroomUsage())
}
