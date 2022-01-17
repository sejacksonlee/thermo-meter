package thermometer

func ConvertToF(c float64) float64 {
	f := (c * 1.8) + 32
	return f
}

func ConvertToC(f float64) float64 {
	c := (f - 32) * 0.56
	return c
}
