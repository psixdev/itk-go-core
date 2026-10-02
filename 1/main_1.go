package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
)

func PrintType(variable any) {
	// fmt.Printf("%T\n", variable)
	fmt.Println(reflect.TypeOf(variable))
}

func JoinVariables(
	numDecimal int,
	numOctal int,
	numHexadecimal int,
	pi float64,
	name string,
	isActive bool,
	complexNum complex64,
) string {
	return strings.Join(
		[]string{
			strconv.FormatInt(int64(numDecimal), 10),
			strconv.FormatInt(int64(numOctal), 8),
			strconv.FormatInt(int64(numHexadecimal), 16),
			strconv.FormatFloat(pi, 'f', -1, 64),
			name,
			strconv.FormatBool(isActive),
			fmt.Sprintf("%v", complexNum),
		},
		"",
	)
}

func StrToRunes(str string) []rune {
	return []rune(str)
}

func HashRunes(runes []rune) string {
	h := sha256.New()

	buf := make([]byte, utf8.UTFMax)
	mid := len(runes) / 2

	for i := 0; i < mid; i++ {
		n := utf8.EncodeRune(buf, runes[i])
		h.Write(buf[:n])
	}

	h.Write([]byte("go-2024"))

	for i := mid; i < len(runes); i++ {
		n := utf8.EncodeRune(buf, runes[i])
		h.Write(buf[:n])
	}

	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	PrintType(numDecimal)
	PrintType(numOctal)
	PrintType(numHexadecimal)
	PrintType(pi)
	PrintType(name)
	PrintType(isActive)
	PrintType(complexNum)

	strVars := JoinVariables(
		numDecimal,
		numOctal,
		numHexadecimal,
		pi,
		name,
		isActive,
		complexNum,
	)
	fmt.Println("Строковое представление: " + strVars)

	runes := StrToRunes(strVars)
	fmt.Printf("Массив рун: %q\n", runes)

	hash := HashRunes((runes))
	fmt.Printf("Хеш: %s\n", hash)
}
