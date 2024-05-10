package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(fecha_hora_minutos())

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
