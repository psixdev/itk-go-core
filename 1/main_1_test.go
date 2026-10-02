package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var numDecimal int = 7            // Десятичная система
var numOctal int = 0123           // Восьмеричная система
var numHexadecimal int = 0x3B     // Шестнадцатиричная система
var numFloat float64 = 9.75       // Тип float64
var name string = "Harry"         // Тип string
var isActive bool = false         // Тип bool
var complexNum complex64 = 5 - 3i // Тип complex64

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}

type printTestCase struct {
	v   any
	out string
}

func TestPrintType(t *testing.T) {
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()

	tests := []struct {
		name     string
		variable any
		expected string
	}{
		{name: "decimal", variable: numDecimal, expected: "int\n"},
		{name: "octal", variable: numOctal, expected: "int\n"},
		{name: "hexadecimal", variable: numHexadecimal, expected: "int\n"},
		{name: "float", variable: numFloat, expected: "float64\n"},
		{name: "string", variable: name, expected: "string\n"},
		{name: "bool", variable: isActive, expected: "bool\n"},
		{name: "complex", variable: complexNum, expected: "complex64\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, w, _ := os.Pipe()
			os.Stdout = w

			PrintType(tt.variable)

			w.Close()

			var buf bytes.Buffer
			io.Copy(&buf, r)

			output := buf.String()

			fmt.Println(output)

			assert.Equal(t, tt.expected, output)
		})
	}
}

func TestJoinVariables(t *testing.T) {
	t.Run("default case", func(t *testing.T) {
		strVars := JoinVariables(
			numDecimal,
			numOctal,
			numHexadecimal,
			numFloat,
			name,
			isActive,
			complexNum,
		)

		assert.Equal(t, "71233b9.75Harryfalse(5-3i)", strVars)
	})

	t.Run("another case", func(t *testing.T) {
		strVars := JoinVariables(
			0,
			01,
			0x2,
			3.,
			"",
			true,
			0+0i,
		)

		assert.Equal(t, "0123true(0+0i)", strVars)
	})
}

func TestStrToRunes(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected []rune
	}{
		{name: "empty string", str: "", expected: []rune{}},
		{name: "non-empty string", str: "foobar", expected: []rune{'f', 'o', 'o', 'b', 'a', 'r'}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runes := StrToRunes(tt.str)

			assert.Equal(t, tt.expected, runes)
		})
	}
}

func TestHashRunes(t *testing.T) {
	tests := []struct {
		name     string
		runes    []rune
		expected string
	}{
		{
			name:  "empty runes",
			runes: []rune{},
			// go-2024
			expected: "66802df107aace17871a5b610ff9eb11706e13477bb24e93966ca80671c0fac6",
		},
		{
			name:  "even num of elements",
			runes: []rune{'f', 'o', 'o', 'b', 'a', 'r'},
			// foogo-2024bar
			expected: "111ec58e17ffc94abf07b07fa27f004d1c9fff253e9ea8957edc115e52b94d01",
		},
		{
			name:  "odd num of elements",
			runes: []rune{'f', 'o', 'b', 'a', 'r'},
			// fogo-2024-bar
			expected: "b90956554e291f50b248e46932c8083aa23bde9f7e833365a8b214ffad30c509",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash := HashRunes(tt.runes)

			assert.Equal(t, tt.expected, hash)
		})
	}
}
