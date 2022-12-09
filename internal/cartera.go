package internal

type Cartera struct {
	ListaMonedas        []*Moneda //Lista con todas las monedas que posee el usuariocd  en su propiedad
	ValorTotal          float64   //Cálculo del valor actual de su cartera en USD
	ValorFuturo         float64   //Estimación del valor de su cartera en un futuro, medido en USD
	PorcentajeConfianza float64   //Porcentaje que indica la confianza en la predicción
}
