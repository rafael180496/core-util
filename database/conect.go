package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	utl "github.com/rafael180496/core-util/utility"

	/*Conexion a mysql*/
	_ "github.com/go-sql-driver/mysql"
	/*Conexion a postgrest*/
	_ "github.com/lib/pq"
	/*Conexion a sql server*/
	_ "github.com/denisenkom/go-mssqldb"
	/*Conexion a sqllite*/
	_ "github.com/mattn/go-sqlite3"
)

type (
	/*StCadConect : Estructura para generar la cadena de  conexiones de base de datos */
	StCadConect struct {
		File    string `json:"filedb"  ini:"filedb"`
		User    string `json:"userName" ini:"userName"`
		Pass    string `json:"pass"   ini:"pass"`
		Name    string `json:"name"  ini:"name"`
		TP      string `json:"tp"    ini:"tp"`
		Host    string `json:"host"    ini:"host"`
		Port    int    `json:"port"  ini:"port"`
		Sslmode string `json:"sslmode" ini:"sslmode"`
	}
	/*StConect : Estructura que contiene la conexion a x TP de base de datos.*/
	StConect struct {
		Conexion     StCadConect
		urlNative    string
		DBGO         *sqlx.DB
		DBTx         *sql.Tx
		DBStmt       *sql.Stmt
		backupScript string
		Queries      map[string]string
	}
)

/*SetBackupScript : setea un scrip backup para la creacion de base de datos en modelos go*/
func (p *StConect) SetBackupScript(sql string) {
	p.backupScript = sql
}

/*ExecBackup : ejecuta el querie backup */
func (p *StConect) ExecBackup() error {
	if len(p.backupScript) <= 0 {
		return fmt.Errorf("number of shares less than or equal to zeros")
	}
	err := p.Con()
	if err != nil {
		return err
	}
	tx := p.DBGO.MustBegin()
	_, err = tx.Exec(p.backupScript)
	if err != nil {
		p.Close()
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		p.Close()
		tx.Rollback()
		return err
	}
	return nil
}

/*SendSQL : envia un sql con los argumentos */
func (p *StConect) SendSQL(code string, args map[string]interface{}) StQuery {
	return StQuery{
		Querie: p.Queries[code],
		Args:   args,
	}
}

/*Close : cierra las conexiones de base de datos intanciadas*/
func (p *StConect) Close() error {
	if p.DBGO == nil {
		return nil
	}
	err := p.DBGO.Close()
	if err != nil {
		return err
	}
	return nil
}

/*NamedIn : procesa los argumentos y sql para agarrar la clausula IN */
func (p *StConect) NamedIn(query StQuery) (string, []interface{}, error) {
	var (
		sqltemp string
		args    []interface{}
		err     error
	)
	sqltemp, args, err = sqlx.Named(query.Querie, query.Args)
	if err != nil {
		return "", nil, err
	}
	sqltemp, args, err = sqlx.In(sqltemp, args...)
	if err != nil {
		return "", nil, err
	}
	sqltemp = p.DBGO.Rebind(sqltemp)

	return sqltemp, args, err
}

/*Trim : Elimina los espacio en cualquier campo string */
func (p *StCadConect) Trim() {
	p.File = utl.Trim(p.File)
	p.User = utl.Trim(p.User)
	p.Pass = utl.Trim(p.Pass)
	p.Name = utl.Trim(p.Name)
	p.TP = utl.Trim(p.TP)
	p.Host = utl.Trim(p.Host)
	p.Sslmode = utl.Trim(p.Sslmode)
	if p.TP == Post && p.Sslmode == "" {
		p.Sslmode = Ssmodes[0]
	}

}

/*ConfigURL : captura una conexion nativa de drive para base de datos*/
func (p *StConect) ConfigURL(url string) {
	p.urlNative = url
}

/*ConfigJSON : Lee las configuraciones de conexion mediante un .json

Ejemplo:

{

	"User":"prueba",
	"Pass":"prueba",
	"Name":"prueba",
	"TP":"POST",
	"host":"Localhost",
	"Port":3000,
	"sslmode":"",
	"filedb":""

}

*/
func (p *StConect) ConfigJSON(PathJSON string) error {
	var (
		err     error
		cad     StCadConect
		ptrArch *os.File
	)
	if !utl.FileExt(PathJSON, "JSON") {
		return fmt.Errorf("the config json file does not exist")
	}
	PathJSON, err = utl.TrimFile(PathJSON)
	if err != nil {
		return err
	}
	ptrArch, err = os.Open(PathJSON)
	if err != nil {
		return err
	}
	defer ptrArch.Close()
	decJSON := json.NewDecoder(ptrArch)
	err = decJSON.Decode(&cad)
	if err != nil {
		return err
	}
	if !cad.ValidCad() {
		return fmt.Errorf("the config file is invalid")
	}
	p.Conexion = cad
	return nil
}

/*ConfigDBX : Lee las configuraciones de conexion mediante un archivo encriptado .dbx este se debe enviar la Pass*/
func (p *StConect) ConfigDBX(path, pass string) error {
	if !utl.FileExt(path, "DBX") {
		return utl.StrErr("No existe el archivo .dbx")
	}
	dataraw, err := utl.ReadFileStr(path)
	if err != nil {
		return err
	}
	cad, err := DecripConect(utl.StrtoByte(dataraw), pass)
	if err != nil {
		return err
	}
	p.Conexion = cad
	return nil
}

/*ConfigINI : Lee las configuraciones de conexion mediante un .ini

Ejemplo:

[database]

User = prueba

Pass = prueba

Name  = prueba

TP = POST

Port = 5433

host = Localhost

sslmode = opcional

filedb = opcional sqllite

*/
func (p *StConect) ConfigINI(PathINI string) error {
	if !utl.FileExt(PathINI, "INI") {
		return fmt.Errorf("the config ini file does not exist")
	}
	cad, err := readIni(PathINI)
	if err != nil {
		return err
	}
	p.Conexion = cad
	return nil
}

/*ConfigENV : lee las configuracion de la base de datos mediante variables de entorno
Ejemplo:
ENV User = prueba
ENV Pass = prueba
ENV Name  = prueba
ENV TP = POST
ENV Port = 5433
ENV HOST = Localhost
ENV SSLMODE = opcional
ENV  FILEDB = opcional sqllite
*/
func (p *StConect) ConfigENV() error {
	var (
		cad StCadConect
	)
	cad.Pass = os.Getenv("PASS")
	cad.User = os.Getenv("USERNAME")
	cad.Name = os.Getenv("NAME")
	cad.TP = os.Getenv("TP")
	cad.Port = utl.ToInt(os.Getenv("PORT"))
	cad.Host = os.Getenv("HOST")
	cad.Sslmode = os.Getenv("SSLMODE")
	cad.File = os.Getenv("FILEDB")
	if !cad.ValidCad() {
		return fmt.Errorf("the config file is invalid")
	}
	p.Conexion = cad
	return nil
}

/*ResetCnx : Limpia la cadena de conexion*/
func (p *StConect) ResetCnx() {
	p.Conexion = StCadConect{}
}

/*ToString : Muestra la estructura  StCadConect*/
func (p *StCadConect) ToString() string {
	return fmt.Sprintf(FORMATTOSTRCONECT, p.Pass, p.Host, p.Name, p.Port, p.Sslmode, p.TP, p.User, p.File)
}

/*ValidCad : valida la cadena de conexion capturada */
func (p *StCadConect) ValidCad() bool {
	p.Trim()
	if !validTp(p.TP) {
		return false
	}
	if p.TP != SQLLite && (!utl.IsNilArrayStr(p.Pass, p.User, p.Name, p.TP, p.Host) || p.Port <= 0) {
		return false
	}
	if p.TP == SQLLite && !utl.IsNilStr(p.File) {
		return false
	}
	return true
}

/*Con : Crear una conexion ala base de datos configurada en la cadena.*/
func (p *StConect) Con() error {
	var (
		err, errping error
	)
	conexion := p.Conexion
	prefijo, cadena := strURL(p.Conexion.TP, conexion)
	cadena = utl.ReturnIf(!utl.IsNilStr(p.urlNative), cadena, p.urlNative).(string)
	if cadena == "" {
		return fmt.Errorf("unsupported DB type")
	}
	if p.DBGO != nil {
		errping = p.DBGO.Ping()
	}
	if errping != nil || p.DBGO == nil {
		if p.Conexion.TP == SQLLite && p.createDB() != nil {
			return fmt.Errorf("the db is sqllite you need the file.d")
		}
		p.DBGO, err = sqlx.Connect(prefijo, cadena)
		if err != nil {
			return err
		}
	}
	return nil
}

/*Insert : Inserta a cualquier tabla donde esta conectado devuelve true si fue guardado o false si no guardo nada.*/
func (p *StConect) Insert(Data []StQuery) error {
	return p.ExecValid(Data, INSERT)
}

/*UpdateOrDelete : actualiza e elimina a cualquier tabla donde esta conectado devuelve la cantidad de filas afectadas.*/
func (p *StConect) UpdateOrDelete(Data []StQuery) (int64, error) {
	err := p.ExecValid(Data, DELETE)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

/*ExecDatatable : ejecuta a nivel de base de datos una accione datable esta puede ser INSERT,DELETE,UPDATE*/
func (p *StConect) ExecDatatable(data DataTable, acc string, indConect bool) error {
	queries, err := data.GenSQL(acc)
	if err != nil {
		return err
	}
	err = p.Exec(queries, indConect)
	if err != nil {
		return err
	}
	return nil
}

/*Exec :Ejecuta una accion de base de datos nativa con rollback*/
func (p *StConect) Exec(Data []StQuery, indConect bool) error {
	return p.execAux(Data, "", false, indConect)
}

/*ExecOne :Ejecuta un StQuery navito haciendo rollback con un error*/
func (p *StConect) ExecOne(Data StQuery, indConect bool) error {
	err := p.Con()
	if err != nil {
		return err
	}
	//Bloque de ejecucion
	tx := p.DBGO.MustBegin()
	_, err = tx.NamedExec(Data.Querie, Data.Args)
	if err != nil {
		p.Close()
		tx.Rollback()
		return err
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

/*ExecValid :Ejecuta una accion de base de datos nativa con rollback y validacion de insert e delete o que TP de accion es */
func (p *StConect) ExecValid(Data []StQuery, tipacc string) error {
	return p.execAux(Data, tipacc, true, false)
}

/*ExecNative :  ejecuta la funcion nativa del paquete sql*/
func (p *StConect) ExecNative(sql string, indConect bool, args ...interface{}) (sql.Result, error) {
	if !utl.IsNilStr(sql) {
		return nil, utl.StrErr("El querie esta vacio")
	}
	err := p.Con()
	if err != nil {
		return nil, err
	}
	tx := p.DBGO.MustBegin()
	rel, err := tx.Exec(sql, args...)
	if err != nil {
		p.Close()
		tx.Rollback()
		return rel, err
	}
	err = tx.Commit()
	if err != nil {
		p.Close()
		tx.Rollback()
		return nil, err
	}
	if !indConect {
		p.Close()
	}
	return rel, nil
}

/*Test : Valida si se puede conectar ala base de datos antes de un  uso.*/
func (p *StConect) Test() bool {
	prueba := new(StQuery)
	switch p.Conexion.TP {
	case Post, Mysql, Sqlser, SQLLite:
		prueba.Querie = `SELECT 1`
	}
	dato, err := p.QueryMap(*prueba, 1, false, true)
	if err != nil || len(dato) <= 0 {
		return false
	}
	return true
}

/*ValidTable : valida si la tabla a buscar existe*/
func (p *StConect) ValidTable(table string) bool {
	prueba := StQuery{
		Querie: TESTTABLE[p.Conexion.TP],
		Args: map[string]interface{}{
			"TABLENAME": table,
		},
	}
	dato, err := p.QueryMap(prueba, 1, false, true)
	if err != nil || len(dato) <= 0 {
		return false
	}
	num, erraux := dato[0].ToInt("REG")
	if num <= 0 || erraux != nil {
		return false
	}
	return true
}
