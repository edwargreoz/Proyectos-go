package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//Enteros con signo
	var a int8 = 127                  //-128 a 127
	var b int16 = 32767               // -32 768 a 32 767
	var c int32 = 2147483647          // -2^31 a 2^31-1
	var d int64 = 9223372036854775807 // -2^63 a 2^63-1

	//Enteros sin signo
	var ua uint8 = 255                   // 0 a 255
	var ub uint16 = 65535                // 0 a 65,535
	var uc uint32 = 4294967295           // 0 a 2^32-1
	var ud uint64 = 18446744073709551615 // 0 a 2^64-1

	//Tipos dependientes de la arquitectura
	var e int = 42  // 32 o 64 bits segun la arquitectura
	var f uint = 42 //32 0 64 bits segun la arquitectura

	//tipos especiales

	var g byte = 25 //Alias para uint8

	var h rune = 'A'            //Alias para int32, Representa un punto de codigo UNICODE
	var i uintptr = 0x123445678 // tamaño suficiente para punteros

	fmt.Printf("Tamaños en bytes:\n")
	fmt.Printf("int8: %d, int16: %d, int32: %d, int64: %d\n",
		unsafe.Sizeof(a), unsafe.Sizeof(b), unsafe.Sizeof(c),
		unsafe.Sizeof(d))
	fmt.Printf("uint8: %d, uint16: %d, uint32: %d, uint64: %d\n",
		unsafe.Sizeof(ua), unsafe.Sizeof(ub), unsafe.Sizeof(uc),
		unsafe.Sizeof(ud))
	fmt.Printf("int: %d, uint: %d, uintptr: %d\n",
		unsafe.Sizeof(e), unsafe.Sizeof(f), unsafe.Sizeof(i))
	fmt.Printf("init32: %d, unit8: %d\n", unsafe.Sizeof(h), unsafe.Sizeof(g))
	// Operaciones aritméticas
	suma := a + int8(10)
	resta := b - int16(100)
	producto := c * int32(2)
	fmt.Printf("Operaciones: suma=%d, resta=%d, producto=%d\n",
		suma, resta, producto)
	// Overflow example (comentado para evitar panic)
	// var overflow int8 = 128 // ERROR: constant 128 overflows int8
	fmt.Scanln()
}
