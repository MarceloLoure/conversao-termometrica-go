package main

import "fmt"

const ebulicaoK int = 373

func main() {

	var ebulicaoC int = ebulicaoK - 273

	fmt.Printf("A temperatura de ebulição em Kelvin é de : %vK e em Celcius é de : %vC", ebulicaoK, ebulicaoC)

}
