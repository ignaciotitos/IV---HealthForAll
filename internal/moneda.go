package internal

type Moneda struct {
	Nombre              string
	Cantidad            float64 //Cantidad de unidades que posee de dicha moneda
	ValorActual         float64 //Valor actual de dicha moneda en USD
	ValorFuturo         float64 //Valor que se estima que tendrá dicha moneda medido en USD
	PorcentajeConfianza float64 //Porcentaje que indica la confianza de la predicción

}
