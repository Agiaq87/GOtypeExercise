package main

import (
	"fmt"
	"math"
	"strings"
	"unsafe"
)

func main() {
	printSize("uint8", unsafe.Sizeof(uint8(0)))
	printSize("uint16", unsafe.Sizeof(uint16(0)))
	printSize("uint32", unsafe.Sizeof(uint32(0)))
	printSize("uint64", unsafe.Sizeof(uint64(0)))

	printSize("int8", unsafe.Sizeof(int8(0)))
	printSize("int16", unsafe.Sizeof(int16(0)))
	printSize("int32", unsafe.Sizeof(int32(0)))
	printSize("int64", unsafe.Sizeof(int64(0)))

	printSize("float32", unsafe.Sizeof(float32(0)))
	printSize("float64", unsafe.Sizeof(float64(0)))

	printSize("complex64", unsafe.Sizeof(complex64(0)))
}

func printSize(name string, value uintptr) {
	fmt.Printf("Il tipo %s occupa %d bit, ovvero %d bytes\n", name, value*8, value)

	if strings.Contains(name, "u") {
		var result = uint64(math.Pow(2, float64(value*8)))
		fmt.Printf("Il range di valori è: [0, %d]\n", result-1)
	} else {
		value *= 8
		var result = uint64(math.Pow(2, float64(value-1)))
		fmt.Printf("Il range di valori è: [-%d, %d]\n", result, result-1)
	}

	println()
}
