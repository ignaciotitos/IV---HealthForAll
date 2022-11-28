package internal

type Moneda struct {
	Nombre        string
	Valor         float64
	Estimacion    float64
	Recomendacion *Recomendacion
	Predicciones  []*Prediccion
}
