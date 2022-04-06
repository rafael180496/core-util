package database

import (
	"encoding/json"
	"strings"
	"time"

	utl "github.com/rafael180496/core-util/utility"
)

type (

	/*StSQLData : Estructura para manipulacion de columnas JSON esto funciona para capturar una sola respuesta a nivel de solucion
	ejemplo
	select getdata() as data --> donde data deve ser un json o []byte
	*/
	StSQLData struct {
		Data []byte `db:"data"`
	}

	/*StQuery : Estructura para ejecutar query de base de datos. */
	StQuery struct {
		Querie string `json:"querie"`
		Args   map[string]interface{}
	}
	/*StData : Estructura que extrae los datos de una consulta de base de datos tramformandola en map*/
	StData map[string]interface{}
)

/*ToJSON : Convierte un  StSQLData a utl.JSON*/
func (p *StSQLData) ToJSON() utl.JSON {
	var data utl.JSON = p.Data
	return data
}

/*ToMap : Convierte un StSQLData a map interface*/
func (p *StSQLData) ToMap() ([]map[string]interface{}, error) {
	return utl.JSONtoObj(p.ToJSON())
}

/*Unmarshal : captura una structura*/
func (p *StSQLData) Unmarshal(v interface{}) error {
	return json.Unmarshal(p.Data, v)
}

/*NewStData : crea un Stdata de un map*/
func NewStData(mp map[string]interface{}) StData {
	var data StData = mp
	return data
}

/*Filter : Excluye key de un map interface*/
func (p *StData) Filter(keys ...string) StData {
	datanew := make(StData)
	clone := *p
	for k, vl := range clone {
		if !utl.InStr(k, keys...) {
			datanew[k] = vl
		}
	}
	return datanew
}

/*ValidColum : valida si un campo existe*/
func (p *StData) ValidColum(col string) bool {
	clone := *p
	_, ok := clone[col]
	return ok
}

/*FindTp : busca el tipo de datos para las tablas con constante prederminadas
INTP : tipo core entero
STTP : tipo core texto
FLTP : tipo core numerico
BLTP : tipo core condicional
DTTP : tipo core date
*/
func FindTp(v interface{}, tp TpCore) interface{} {
	switch tp {
	case INTP:
		return utl.ToInt(v)
	case STTP:
		return utl.ToString(v)
	case DTTP:
		vl := utl.ToString(v)
		date, err := utl.StringToDate(vl)
		return utl.ReturnIf(err != nil, utl.FNulo(), date)
	case FLTP:
		return utl.ToFloat(v)
	case BLTP:
		return utl.ToBoolean(v)
	case JSONTP:
		data, _ := utl.NewJSON(v)
		return data
	default:
		return nil
	}
}

/*UpperKey : coloca en mayusculas todas las keys*/
func (p *StData) UpperKey() StData {
	datanew := make(StData)
	data := *p
	for k := range data {
		kNew := strings.ToUpper(k)
		datanew[kNew] = data[k]
	}
	return datanew
}

/*KeyColum : envia las columnas que contiene la data*/
func (p *StData) KeyColum() []string {
	var colums []string
	clone := *p
	for k := range clone {
		colums = append(colums, k)
	}
	return colums
}

/*ToJSON : Convierte la estructura  StData en JSON para envios externos a rest api.*/
func (p *StData) ToJSON() ([]byte, error) {
	jsonData, err := json.Marshal(&p)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/*ToString : Convierte el valor del map interface{} a string.*/
func (p *StData) ToString(columna string) (string, error) {
	var valor interface{} = (*p)[columna]
	return utl.ToString(valor), nil
}

/*ToInt : Convierte el valor del map interface{} a int.*/
func (p *StData) ToInt(columna string) (int, error) {
	var valor interface{} = (*p)[columna]
	return utl.ToInt(valor), nil
}

/*ToInt32 : Convierte el valor del map interface{} a int32.*/
func (p *StData) ToInt32(columna string) int32 {
	var valor interface{} = (*p)[columna]
	return utl.ToInt32(valor)
}

/*ToBool : Convierte el valor del map interface{} a bool.*/
func (p *StData) ToBool(columna string) bool {
	var valor interface{} = (*p)[columna]
	return utl.ToBoolean(valor)
}

/*ToInt64 : Convierte el valor del map interface{} a int64.*/
func (p *StData) ToInt64(columna string) (int64, error) {
	var valor interface{} = (*p)[columna]
	return utl.ToInt64(valor), nil
}

/*ToFloat : Convierte el valor del map interface{} a float.*/
func (p *StData) ToFloat(columna string) (float32, error) {
	var valor interface{} = (*p)[columna]
	return utl.ToFloat(valor), nil
}

/*ToFloat64 : Convierte el valor del map interface{} a float.*/
func (p *StData) ToFloat64(columna string) (float64, error) {
	var valor interface{} = (*p)[columna]
	return utl.ToFloat64(valor), nil
}

/*ToDate : Convierte el valor del map interface{} a time.*/
func (p *StData) ToDate(columna string) (time.Time, error) {
	var valor interface{} = (*p)[columna]
	vl := utl.ToString(valor)
	return utl.StringToDate(vl)
}
