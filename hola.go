package main

import (
	"fmt"
	"time"
)

// Escribir un programa que pida al usuario un número entero positivo y muestre por pantalla todos los números impares desde 1 hasta ese número separados por comas.
// Escribir un programa que pida al usuario un número entero positivo y muestre por pantalla la cuenta atrás desde ese número hasta cero separados por comas.
// Escriba un programa que pregunte cuántos números se van a introducir, pida esos números, y muestre un mensaje cada vez que un número no sea mayor que el primero.

func eu() {

	var edad int

	fmt.Println("Ingresa tu edad")

	fmt.Scan(&edad)

	total := 100 - edad

	fmt.Printf("Tienes: %d años", edad)

	fmt.Printf("Te faltan: %d años para cumplir 100", total)

}

func oof() {

	var largo int

	var ancho int

	fmt.Println("Bienvenido, ingresa el el largo del cuadrado:")

	fmt.Scan(&largo)

	fmt.Println("ok, ahora ingresa el ancho del cuadrado:")

	fmt.Scan(&ancho)

	var total = largo * ancho

	fmt.Printf("El area del cuadrado es: %d", total)

}
func temp() {

	tiempo := time.Now()

	fmt.Println(tiempo)

}

func moises() {
	var nombre string

	fmt.Println("Dime tu nombre")

	fmt.Scan(&nombre)

	time.Sleep(2 * time.Second)

	fmt.Println("Procesando todo...")

	time.Sleep(1 * time.Second)

	fmt.Printf("Tu nombre es: %s ,\n cerrando el programa.. \n", nombre)

	time.Sleep(4 * time.Second)

	fmt.Println("cerrado")
}

func face() {

	var nombre string

	var cont string

	fmt.Printf("Bienvenido! , \n Ingresa tu nombre \n")

	fmt.Println("--------------------------------------")

	fmt.Scan(&nombre)

	time.Sleep(3 * time.Second)

	fmt.Println("Ok, ahora ingresa tu contraseña:")

	fmt.Println("----------------------------")

	fmt.Scan(&cont)

	time.Sleep(4 * time.Second)

	fmt.Printf("Bienvenido %s", nombre)

}

func xdd() {

	var anio = 2024

	var edad int

	fmt.Println("Buenas tardes, ingrese su edad:")

	fmt.Scan(&edad)

	var edad_max = edad + 10

	for edad < edad_max {

		edad++

		anio++

		fmt.Printf("En el %d tendrias %d \n", anio, edad)
	}

}

func gg() {

	var largo int

	fmt.Println("Bienvenido, ingresa el el largo del cuadrado:")

	fmt.Scan(&largo)

	stars := "*"
	for i := 0; i < largo; {
		fmt.Println(stars)
		stars = stars + "*"
		i++
	}

}

func for_piramide() {
	var cantd_filas int

	fmt.Println("Ingresa las filas")

	fmt.Scan(&cantd_filas)

	numero_imp := 1
	fila_act := 0
	for fila_act <= cantd_filas {
		for range fila_act {
			fmt.Printf("%d ", numero_imp)
			numero_imp++
		}
		fmt.Println("")
		fila_act++
	}

}

func fibbo() {
	var cantd_digito int

	fmt.Println("Ingresa la cantidad de digitos")

	fmt.Scan(&cantd_digito)

	n1 := 0
	n2 := 1
	digito_actual := 1
	for digito_actual <= cantd_digito {
		fmt.Printf("%d ", n1)
		aux := n2
		n2 = n1 + n2
		n1 = aux
		digito_actual++
	}

}

func doblrefor() {
	contador := 0
	for contador < 3 {
		fmt.Println("Principio")
		for i := 0; i < 5; i++ {
			fmt.Println("	Principio 2do")
			fmt.Println("U")
			fmt.Println("	Fin 2do")
		}
		fmt.Println()
		contador++
		fmt.Println("Fin")
	}
}

func impar() {

	var Nnatural int

	fmt.Println("Ingresa un numero natural")

	fmt.Scan(&Nnatural)

	for i := 1; i <= Nnatural; i = i + 2 {
		fmt.Printf(" %d,", i)
	}

}

func resta() {

	var npositivo int

	fmt.Println("Ingresa un numero natural")
	fmt.Scan(&npositivo)

	for i := npositivo; i >= 0; {

		fmt.Printf("%d,", i)

		i--

	}

}
