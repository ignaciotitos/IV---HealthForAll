package internal

import "errors"

type TasaInflacionMensual struct {
	Mes   uint8
	Anio  uint16
	Valor float32
}

func NuevaTasaInflacionMensual(Mes uint8, Anio uint16, Valor float32) (*TasaInflacionMensual, error) {
	if 1 <= Mes && Mes <= 12 && 2012 <= Anio && Anio <= 2022 {
		return &TasaInflacionMensual{Mes, Anio, Valor}, nil
	}

	return nil, errors.New("Ya existe una tasa de inflación con esa fecha")
}
