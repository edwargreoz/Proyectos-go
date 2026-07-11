package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("----Estructuras if / else-----")

	//if basico

	edad := 30

	if edad >= 18 {
		fmt.Println("Es mayor de edad")
	}

	//if-else
	temperatura := 22

	if temperatura > 25 {
		fmt.Println("Hace calor y la temperatura es de ", temperatura, "°")
	} else {
		fmt.Println("La temperatura es agradable y es de ", temperatura, "°")
	}

	//if-else-if (cadena)
	puntuacion := 85

	if puntuacion >= 90 {
		fmt.Println("Excelente mi crack")
	} else if puntuacion >= 75 {
		fmt.Println("Ta bien mi crack")
	} else if puntuacion >= 60 {
		fmt.Println("No se rick tienes que mejorar")
	} else {
		fmt.Println("No pasa nada contigo compare")
	}

	//if con inicializacion (patron muy comun en Go)
	if hora := time.Now().Hour(); hora < 12 {
		fmt.Println("Buenos dias")

	} else if hora < 18 {
		fmt.Println("Buenas tardes")
	} else {
		fmt.Println("Buenas noches")
	}

	//Verificacion de errores (patron idiomático)
	if numero, err := strconv.Atoi("123"); err != nil {
		fmt.Printf("Error de conversion: %v\n", err)
	} else {
		fmt.Printf("Numero convertido: %d\n", numero)
	}

	//MULTIPLE CONDICIONALES
	usuario := "admin"
	password := "secret123"

	if usuario == "admin" && password == "secret123" {
		fmt.Println("Acceso concedido")
	} else {
		fmt.Println("Acceso denegado")
	}

	//Condicionales complejas
	estado := "activo"
	ultimoacceso := time.Now().Add(-24 * time.Hour)

	if estado == "activo" && time.Since(ultimoacceso) < 30*24*time.Hour {
		fmt.Println("Usuario activo pero inactivo por tiempo")
	} else {
		fmt.Println("Usuario inactivo")
	}

	//caso practico
	demostrarCasoPracticosI()
	fmt.Scanln()
}

func demostrarCasoPracticosI() {
	fmt.Println("\n--- Caso practico con if ---")
	//Validacion de entrada

	email := "usuario@dominio.com"

	if len(email) == 0 {
		fmt.Println("Email vacio")
	} else if !strings.Contains(email, "@") {
		fmt.Println("Email invalido: falta @")
	} else if !strings.Contains(email, ".") {
		fmt.Println("Email invalido: falta dominio")
	} else {
		fmt.Println("Email valido")
	}

	//2. Categorizacion de rangos
	velocidad := 75 // km/h
	limite := 60

	if velocidad <= limite {
		fmt.Println("Velocidad normal")
	} else if velocidad <= limite+10 {
		fmt.Println("Ligero exceso de velocidad")
	} else if velocidad <= limite+20 {
		fmt.Println("Exceso moderado - multa")
	} else {
		fmt.Println("Exceso grave - suspension")
	}

	// 3. logica de negocio con multiples factores

	edad := 35
	experiencia := 3 //años
	certificaciones := 2

	if edad >= 21 && experiencia >= 2 && certificaciones >= 1 {
		fmt.Println("Candidato calificado para posicion senior")
	} else if edad >= 18 && (experiencia >= 1 || certificaciones >= 1) {
		fmt.Println("Candidadto calificado para posicion junior")
	} else if edad >= 18 {
		fmt.Println("Candidato para posicion de entrenamiento")
	} else {
		fmt.Println("No cumple con requisitos minimos")
	}

	//4. Manejo de casos especiales
	valor := 0.0

	if valor > 0 {
		fmt.Printf("Valor positivo: %.2f\n", valor)
	} else if valor < 0 {
		fmt.Printf("valor negativo: %.2f\n", valor)
	} else {
		// Caso especial: exactamente 0
		fmt.Println("Valor es exactamente cero")
	}
	// 5. Verificacion de recursos

	memoryUsage := 50 //porcentaje
	cpuUsage := 70.2
	diskUsage := 45.0

	alertLevel := "normal"

	if memoryUsage > 90 || cpuUsage > 90 || diskUsage > 90 {
		alertLevel = "critrico"
	} else if memoryUsage > 80 || cpuUsage > 80 || diskUsage > 85 {
		alertLevel = "warning"
	}

	switch alertLevel {
	case "critico":
		fmt.Println("ALERTA CRITICA: Recursos del sistema agotados")
	case "warning":
		fmt.Println("ADVERTENCIA: Alto uso de recursos")
	default:
		fmt.Println("Recursos del sistema Normales")
	}
	fmt.Scanln()
}
