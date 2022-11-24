package internal

type Prediccion struct{
	Moneda string
	Datos struct{
		TipoCambio string
		CrecimientoPIB float64
		DeudaExterior float64
		BalanzaCuentaCorriente float64
		CreditoSectorPublico float64
		CreditoSectorPrivado float64
		NivelPrecios float64
		NivelReservas float64
		DiferencialTipoInteres float64
		IndiceBolsa float64
	}

}