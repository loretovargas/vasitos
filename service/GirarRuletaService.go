package service

import (
	"errors"
	"ruleta/service/port/in"
)

type GirarRuletaService struct {
}

func (service GirarRuletaService) GirarRuleta(command in.GirarRuletaCommand) error {

	
	if command.NumeroSeleccionado.EsMenorOIgualACero(){
		return errors.New("el numero no puede ser menor o igual a 0")
	}

	if command.NumeroSeleccionado.EsMayorASeis() {
		return errors.New("el numero no puede ser mayor que 6")
	}

	
	return nil
}
