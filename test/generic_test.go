package test

import (
	"fmt"
	"reflect"
	"testing"

	utl "github.com/rafael180496/core-util/utility"
)

/*SubString : recorta una cadena */
func TestSubString(t *testing.T) {
	fmt.Printf("%s", utl.SubString("18:20", 0, 2))
	fmt.Printf("%s", utl.SubString("18:20", 3, 5))
}

/*TestReturnIf : retorna con un condicional */
func TestReturnIf(t *testing.T) {
	t.Logf("%s", utl.ReturnIf(5 > 4, "It's true", "It's false"))
}

/*TestTrim : Quita los espacio de un texto */
func TestTrim(t *testing.T) {
	text := "Hola Mundo TDA"
	t.Logf("text:[%s]", text)
	t.Logf("trim:[%s]", utl.Trim(text))
}

/*TestFloat64 : Redondea un valor float a x decimales*/
func TestFloat64(t *testing.T) {
	valor := 12.34661
	t.Logf("valor:[%f]", valor)
	valor = utl.RoundFloat64(valor, 2)
	t.Logf("valor:[%f]", valor)
}

/*Test_Merge_string : combina unas cadena de caracteres*/
func Test_Merge_string(t *testing.T) {
	a := ""
	b := "foo"
	expected := "foo"
	ret := utl.Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

/*Test_Merge_Array : combina los array de un mismo tipo de datos*/
func Test_Merge_Array(t *testing.T) {
	a := []int{10, 11, 12}
	b := []int{}
	expected := []int{10, 11, 12}
	ret := utl.Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}
