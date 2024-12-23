package main

import "fmt"

func main() {
	const ebulicaoK = 373.0

	tempK := ebulicaoK
	tempC := tempK - 273.0

	fmt.Printf("A ebulição da água em K = %g, e a ebulição da água em °C = %g.", tempK, tempC)

}
