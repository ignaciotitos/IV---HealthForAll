package internal

type Moneda struct {
	ValorActual         float64
	Cantidad            float32
	DatosMercadoDivisas struct {
		NivelRentaNacional         float64
		PreciosRelativosBienes     float64
		NivelInversionDomestico    float64
		DiferencialIntereses       float64
		SaldoBalanzaPagos          float64
		RentaMundial               float64
		PreciosRelativos           float64
		CalidadProducciónDoméstica float64
	}
}
