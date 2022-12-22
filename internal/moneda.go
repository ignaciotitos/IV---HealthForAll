package internal

type Moneda struct {
	ValorActual   float64
	Cantidad      float32
	IndicePrecios map[int]float32
}
