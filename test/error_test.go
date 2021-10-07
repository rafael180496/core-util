package test

import (
	"fmt"
	"testing"

	utl "github.com/rafael180496/core-util/utility"
)

/*TestSendError : Envia un error con mensaje */
func TestSendError(t *testing.T) {
	t.Logf("Error:%s", utl.Msj.GetError("GE01").Error())
}

/*TestSendError : Envia un error con mensaje */
func TestSendtTrycatch(t *testing.T) {
	fmt.Println("Comienza")
	utl.Block{
		Try: func() {
			fmt.Println("Entra al Try")
			utl.Throw("Error")
		},
		Catch: func(e utl.Exception) {
			fmt.Printf("Error a capturar %v\n", e)
		},
		Finally: func() {
			fmt.Println("Finalizacion...")
		},
	}.Do()
	fmt.Println("Termina")
}
