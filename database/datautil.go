package database

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	utl "github.com/rafael180496/core-util/utility"
	"gopkg.in/ini.v1"
)

/*strURL : Arma la cadena de conexion dependiendo del tipo*/
func strURL(tp string, conexion StCadConect) (string, string) {
	switch tp {
	case Post:
		/*"postgres://user:password@localhost:port/bdnamme?sslmode=verify-full"*/
		return PrefijosDB[tp], fmt.Sprintf(CADCONN[tp], conexion.User, conexion.Pass, conexion.Host, conexion.Port, conexion.Name, conexion.Sslmode)
	case Mysql:
		/*sql.Open("mssql", "user:password@tcp(localhost:5555)/dbname?tls=skip-verify&autocommit=true") */
		return PrefijosDB[tp], fmt.Sprintf(CADCONN[tp], conexion.User, conexion.Pass, conexion.Host, conexion.Port, conexion.Name)
	case Sqlser:
		return PrefijosDB[tp], fmt.Sprintf(CADCONN[tp], conexion.Host, conexion.User, conexion.Pass, conexion.Port, conexion.Name)
	case SQLLite:
		return PrefijosDB[tp], fmt.Sprintf(CADCONN[tp], conexion.File)
	default:
		return "", ""
	}
}

/*urlConect : manda la conexion de base datos correspondiente por tipo*/
func (p *StConect) urlConect() (string, string) {
	if utl.Trim(p.urlNative) != "" {
		return PrefijosDB[p.Conexion.TP], p.urlNative
	}
	return strURL(p.Conexion.TP, p.Conexion)
}

/*ValidPrefix : Valida si el prefijo de base de datos existe */
func ValidPrefix(tp string) bool {
	_, exist := PrefijosDB[tp]
	return exist
}

/*validTipDB : valida el tipo de sql que insertara.*/
func validTipDB(sqlOrig string, tp string) error {
	var (
		valid bool
	)
	sqlOrig = strings.ToUpper(sqlOrig)

	switch tp {
	case INSERT:
		valid = strings.Contains(sqlOrig, INSERT)
		if !valid {
			return fmt.Errorf("it does not have the insert")
		}
		return nil
	case UPDATE, DELETE:
		valid = strings.Contains(sqlOrig, UPDATE)
		if !valid {
			valid = strings.Contains(sqlOrig, DELETE)
			if !valid {
				return fmt.Errorf("it does not have valid prefixes")
			}
		}
		return nil
	default:
		return fmt.Errorf("it does not have valid prefixes")
	}

}

/*sendData : captura los datos de la tabla*/
func sendData(valores []interface{}, columnas []string) StData {
	data := make(StData)
	for i, col := range valores {
		if col == nil {
			continue
		}
		switch col := col.(type) {

		case []byte:
			data[columnas[i]] = string(col)
		default:
			data[columnas[i]] = col
		}
	}
	return data
}

/*scanData : escanea las fila regresando un tipo generico */
func scanData(rows *sqlx.Rows, maxRows int, indLimit bool) ([]StData, error) {
	var (
		result    []StData
		columnas  []string
		err       error
		countRows = 0
	)
	maxRows = utl.ReturnIf(maxRows <= 0, 1, maxRows).(int)
	columnas, err = rows.Columns()
	if err != nil {
		return result, fmt.Errorf("columns were not obtained")
	}
	ptrData := make([]interface{}, len(columnas))
	valores := make([]interface{}, len(columnas))
	for i := range valores {
		ptrData[i] = &valores[i]
	}
	for rows.Next() {
		if indLimit {
			if countRows > maxRows {
				break
			}
			countRows++
		}
		err = rows.Scan(ptrData...)
		if err != nil {
			return result, err
		}
		data := sendData(valores, columnas)
		result = append(result, data)
	}
	return result, nil
}

/*validTp : valida los tipos de conexion disponible*/
func validTp(tp string) bool {
	switch tp {
	case SQLLite, Post, Mysql, Sqlser:
		return true
	default:
		return false
	}
}
func readIni(source interface{}) (StCadConect, error) {
	var cnx StCadConect
	cfg, err := ini.Load(source)
	if err != nil {
		return cnx, err
	}
	err = cfg.Section("database").MapTo(&cnx)
	if err != nil {
		return cnx, err
	}
	if !cnx.ValidCad() {
		return cnx, fmt.Errorf("the config file is invalid")
	}
	return cnx, nil
}

/*funcion para crear base de datos sqllite*/
func (p *StConect) createDB() error {
	if utl.FileExt(p.Conexion.File, "DB") {
		return nil
	}
	_, err := utl.FileNew(p.Conexion.File)
	if err != nil {
		return err
	}
	return nil
}

/*queryGeneric : ejecuta sql dinamicos regresando un map*/
func (p *StConect) queryGeneric(query StQuery, cantrow int, indConect, indLimit bool) ([]StData, error) {
	var (
		err     error
		filas   *sqlx.Rows
		result  []StData
		args    []interface{}
		sqltemp string
	)
	err = p.Con()
	if err != nil {
		return result, err
	}
	sqltemp, args, err = p.NamedIn(query)
	if err != nil {
		p.Close()
		return result, err
	}
	filas, err = p.DBGO.Queryx(sqltemp, args...)
	if err != nil {
		p.Close()
		return result, err
	}
	result, err = scanData(filas, cantrow, indLimit)
	if err != nil {
		p.Close()
		filas.Close()
		return result, err
	}
	if !indConect {
		p.Close()
	}
	filas.Close()
	return result, nil
}

/*execAux : Ejecuta una accion de base de datos  auxiliar*/
func (p *StConect) execAux(Data []StQuery, tipACC string, indvalid, indConect bool) error {
	if len(Data) <= 0 {
		return fmt.Errorf("number of shares less than or equal to zeros")
	}
	err := p.Con()
	if err != nil {
		return err
	}
	//Bloque de ejecucion
	tx := p.DBGO.MustBegin()
	for _, dat := range Data {
		if indvalid {
			err = validTipDB(dat.Querie, tipACC)
			if err != nil {
				p.Close()
				tx.Rollback()
				return err
			}
		}
		_, err = tx.NamedExec(dat.Querie, dat.Args)
		if err != nil {
			p.Close()
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		p.Close()
		tx.Rollback()
		return err
	}
	if !indConect {
		p.Close()
	}
	return nil
}
