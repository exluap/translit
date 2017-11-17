package translit_test

import (
	"fmt"
	"github.com/exluap/translit"
)

func ExampleRuTranslit() {
	tests := []string{
		"Проверочная СТРОКА для транслитерации",
		"ЧАЩА",
		"ЧаЩа",
		"Чаща",
		"чаЩА",
	}
	for _, text := range tests {
		fmt.Println(translit.Ru(text))
	}
	// Output:
	// Proverochnaja STROKA dlja transliteracii
	// CHASCHA
	// ChaScha
	// Chascha
	// chaSCHA
}
