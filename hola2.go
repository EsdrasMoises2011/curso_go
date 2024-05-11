package main

import (
	"fmt"
	"time"
)

func main() {
	numeros := []int{1, 2, 3, 4}
	numeros_volteados := make([]int, 0)
	fmt.Println(len(numeros))

	for x := 0; x < len(numeros); x++ {
		numeros_volteados = append(numeros_volteados, numeros[len(numeros)-x-1])
	}

	fmt.Println(numeros_volteados)
}

func xd() {
	var cantidad int

	fmt.Scanln(&cantidad)

	slice_numeros := make([]int, cantidad)

	for i := 0; i < cantidad; i++ {
		fmt.Println("Ingresa los datos:")
		var numero int
		fmt.Scanln(&numero)
		slice_numeros[i] = numero
	}

	fmt.Println(slice_numeros)
}

func leer_numero(mensaje string) (int, error) {
	fmt.Println(mensaje)
	var numero int
	_, err := fmt.Scanln(&numero)
	if err != nil {
		return 0, err
	}
	return numero, nil
}

func fecha_hora_minutos() string {
	tiempo := time.Now()

	dia := tiempo.Day()
	mes := tiempo.Month()
	anio := tiempo.Year()

	hora := tiempo.Hour()
	minuto := tiempo.Minute()
	segundo := tiempo.Second()

	return fmt.Sprintf("%d/%d/%d %d:%d:%d", dia, mes, anio, hora, minuto, segundo)
}
