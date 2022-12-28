package internal

type VariacionValorMoneda struct {
	TasasInflacionMesesAnteriores []*TasaInflacionMensual
}

func (this VariacionValorMoneda) añadirTasaAListaDeTasas(tasa *TasaInflacionMensual) {
	if fechaExistenteEnLista(this.TasasInflacionMesesAnteriores, tasa) == false {
		this.TasasInflacionMesesAnteriores = append(this.TasasInflacionMesesAnteriores, tasa)
	}

}

func fechaExistenteEnLista(lista []*TasaInflacionMensual, tasa *TasaInflacionMensual) bool {
	for _, tasaLista := range lista {
		if (tasaLista.Mes == tasa.Mes) && (tasaLista.Anio == tasa.Anio) {
			return true
		}
	}
	return false
}
