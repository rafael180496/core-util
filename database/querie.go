package database

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	utl "github.com/rafael180496/core-util/utility"
)

/*QueryNative :  ejecuta la funcion nativa del paquete sql*/
func (p *StConect) QueryNative(sql string, indConect bool, args ...interface{}) (*sql.Rows, error) {
	if !utl.IsNilStr(sql) {
		return nil, utl.StrErr("el Query esta vacio")
	}
	err := p.Con()
	if err != nil {
		return nil, err
	}
	rows, err := p.DBGO.Query(sql, args...)
	if err != nil {
		p.Close()
		return rows, err
	}
	if !indConect {
		p.Close()
	}
	return rows, nil
}

/*QueryOne : Ejecuta un querie en la base de datos y devuelve un map dinamico pero solo envia una fila y no un arreglo
 */
func (p *StConect) QueryOne(query StQuery, indConect bool) (StData, error) {
	result, err := p.queryGeneric(query, 1, indConect, true)
	if err != nil {
		return nil, err
	}
	if len(result) <= 0 {
		return nil, fmt.Errorf("no hay datos en la consulta")
	}

	return result[0], nil
}

/*Query : Ejecuta un querie en la base de datos y
  devuelve un map dinamico para mostrar los datos donde le limitan la cantida
	de registro que debe de devolver
	indConect = true deja la conexion abierta
  Ejemplo:
  map[COD_CLI:50364481 NIS_RAD:5355046 SEC_NIS:1]
*/
func (p *StConect) Query(query StQuery, cantrow int, indConect bool) ([]StData, error) {
	if cantrow <= 0 {
		return nil, fmt.Errorf("row quantity is zero")
	}
	return p.queryGeneric(query, cantrow, indConect, true)
}

/*QueryMap : Ejecuta un querie en la base de datos y
  devuelve un map dinamico para mostrar los datos donde le limitan la cantida
	de registro que debe de devolver
	indConect = true deja la conexion abierta
	indLimit = true limite de fila si esta en false desactiva esta opcion
*/
func (p *StConect) QueryMap(query StQuery, cantrow int, indConect, indLimit bool) ([]StData, error) {
	result, err := p.queryGeneric(query, cantrow, indConect, indLimit)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*QueryJSON : Ejecuta un querie en la base de datos y
  devuelve un json dinamico para mostrar los datos donde le limitan la cantida
	de registro que debe de devolver
	indConect = true deja la conexion abierta
	indLimit = true limite de fila si esta en false desactiva esta opcion
*/
func (p *StConect) QueryJSON(query StQuery, cantrow int, indConect, indLimit bool) ([]byte, error) {
	result, err := p.queryGeneric(query, cantrow, indConect, indLimit)
	if err != nil {
		return nil, err
	}
	JSON, errJSON := json.Marshal(&result)
	if errJSON != nil {
		return nil, errJSON
	}
	return JSON, nil
}

/*QueryStruct : Ejecuta un query en la base de datos y
  captura la data con struct
	EjecutarQueryStruct(&data,sql,true)
	indConect = true deja la conexion abierta
*/
func (p *StConect) QueryStruct(datadest interface{}, query StQuery, indConect bool) error {
	var (
		err     error
		args    []interface{}
		sqltemp string
	)
	err = p.Con()
	if err != nil {
		return err
	}
	sqltemp, args, err = p.NamedIn(query)
	if err != nil {
		return err
	}
	err = p.DBGO.Select(datadest, sqltemp, args...)
	if err != nil {
		p.Close()
		return err
	}
	if !indConect {
		p.Close()
	}
	return nil
}

/*QueryRows : Ejecuta un query en la base de datos y
  devuelve un puntero de *Rows de sqlx
	indConect = true deja la conexion abierta
*/
func (p *StConect) QueryRows(query StQuery, indConect bool) (*sqlx.Rows, error) {
	var (
		err     error
		filas   *sqlx.Rows
		sqltemp string
		args    []interface{}
	)
	err = p.Con()
	if err != nil {
		return filas, err
	}
	sqltemp, args, err = p.NamedIn(query)
	if err != nil {
		return filas, err
	}
	filas, err = p.DBGO.Queryx(sqltemp, args...)
	if err != nil {
		p.Close()
		return filas, err
	}
	if !indConect {
		p.Close()
	}
	return filas, nil
}
