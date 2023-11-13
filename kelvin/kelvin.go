package main

import "fmt"

const ebulicaoK float64 = 373

func main() {

	tempK := ebulicaoK
	tempC := (tempK - 273)

	fmt.Printf("A temperatura de ebulição da água em Celsius é de %g° e em Kelvin é %g K.", tempC, tempK)
}
