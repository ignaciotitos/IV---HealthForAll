package internal

type Cartera struct {
	ListaMonedas      []*Moneda          //Lista con todas las monedas que posee el usuario en su propiedad
	PrediccionMonedas map[string]float32 //Diccionario dónde a cada moneda se le asocia un % (positivo o negativo) de la estimación de modificación del valor de esa moneda en el futuro
}
