package utility

import (
	"fmt"
	"time"
)

/*Implementacion de try catch*/

type (
	/*StError : Estructura de manejo de error con mensajes personalizados*/
	StError struct {
		When time.Time
		What string
	}
	/*Exception : error capturado*/
	Exception interface{}
	/*Block : bloque de codigo que maneja excepciones*/
	Block struct {
		Try     func()
		Catch   func(Exception)
		Finally func()
	}
)

/*Throw : captura el error y detiene el proceso*/
func Throw(up Exception) {
	panic(up)
}

/*Do : ejecuta el bloque de codigo con exepciones*/
func (tcf Block) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}

/*Error : Muestra el mensaje personalizado del error*/
func (p *StError) Error() string {
	return fmt.Sprintf("%v: %v", time.Now(), p.What)
}

/*MandError : Envia un mensaje con la interfaz de error personalizado.*/
func (p *StError) MandError(mensaje string) error {
	p.What = mensaje
	p.When = time.Now()
	return p
}

/*StrErr : envia un string con errores personalizado*/
func StrErr(str string) error {
	return fmt.Errorf("%s", str)
}
