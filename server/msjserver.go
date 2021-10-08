package server

import (
	"fmt"

	echo "github.com/labstack/echo/v4"
)

type (
	/*MapMsjPet : maneja el envio de mensajes con peticiones echo*/
	MapMsjPet map[string]StMsjPet
	/*StMsjPet : envia un mensaje ya con el error pre-cargado para una respuesta */
	StMsjPet struct {
		Code int
		Msj  string
	}
)

/*Send : envia un mensaje de la peticiones pre-cargada */
func (p *MapMsjPet) Send(cod string, data interface{}, e echo.Context) error {
	var (
		datainfo StDataEnv
	)
	maps := *p
	petmsj := maps[cod]
	datainfo.Code = petmsj.Code
	datainfo.Error = petmsj.Msj
	datainfo.Data = data
	return datainfo.ResultJSON(e)
}

/*SendFormat : envia un mensaje de la peticiones pre-cargada con un formato de mensaje */
func (p *MapMsjPet) SendFormat(cod string, data interface{}, e echo.Context, param ...interface{}) error {
	var (
		datainfo StDataEnv
	)
	maps := *p
	petmsj := maps[cod]
	datainfo.Code = petmsj.Code
	datainfo.Error = fmt.Sprintf(petmsj.Msj, param...)
	datainfo.Data = data
	return datainfo.ResultJSON(e)
}
