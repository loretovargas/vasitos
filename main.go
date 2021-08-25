package main

import (
	"log"
	"vasitos/domain"
)

func main() {
	v := domain.Vasito{}
	v.TieneBolita = true
	var mensaje string
	if v.TieneBolita == true {
		mensaje = "el vaso tiene bolita"

	}
	if v.TieneBolita == false {
		mensaje = "el vaso no tiene bolita"
	}
	log.Print(mensaje)

}
