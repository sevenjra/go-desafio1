package main

import "fmt"

//declaração da variável CONST do ponto de ebulição da água em F
const ebulicaoK = 373

//função principal
func main() {

	tempK := ebulicaoK     
	tempC := (tempK - 273) 
	fmt.Printf("A temperatura de ebulição da água em K = %v , temperatura de ebulição da água em °C =%v .", tempK, tempC)

}