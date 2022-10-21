package database

type (
	/*TpCore : Enums de tipos de datos para los sql e validacion*/
	TpCore string
)

const (
	/*Tp tipos tp*/

	/*INTP : tipo core entero*/
	INTP TpCore = "integer"
	/*STTP : tipo core texto*/
	STTP TpCore = "string"
	/*FLTP : tipo core numerico*/
	FLTP TpCore = "number"
	/*BLTP : tipo core condicional*/
	BLTP TpCore = "bool"
	/*DTTP : tipo core date*/
	DTTP TpCore = "date"
	/*JSONTP : tipo core json*/
	JSONTP TpCore = "json"

	/*SQLLite : conexion tipo sqllite
	https://github.com/mattn/go-sqlite3
	*/
	SQLLite = "SQLLITE"
	/*Ora : conexion tipo oracle
	https://github.com/sijms/go-ora
	*/
	Ora = "ORA"
	/*Post : conexion tipo postgres
	https://github.com/lib/pq
	*/
	Post = "POST"
	/*Mysql : conexion tipo mysql
	https://github.com/go-sql-driver/mysql
	*/
	Mysql = "MYSQL"
	/*Sqlser : conexion tipo sql server
	https://github.com/denisenkom/go-mssqldb
	*/
	Sqlser = "SQLSER"
	/*INSERT : prefijo de insert */
	INSERT = "INSERT"
	/*UPDATE : prefijo de UPDATE */
	UPDATE = "UPDATE"
	/*DELETE : prefijo de DELETE */
	DELETE = "DELETE"
	/*SELECT : prefijo de select*/
	SELECT = "SELECT"
	/*FROM : prefijo de tablas */
	FROM = "FROM"
)

var (
	/*TESTTABLE : quieries para probar si una tabla existe en la base de datos*/
	TESTTABLE = map[string]string{
		Ora: `
		SELECT CASE WHEN COUNT(*) > 0
			THEN 1 ELSE 0 END REG
		FROM ALL_TABLES
		WHERE  TABLE_NAME = :TABLENAME
		`,
		Post: `
		SELECT CASE WHEN EXISTS (
			SELECT FROM PG_TABLES
			WHERE    TABLENAME  = :TABLENAME
			) THEN 1 ELSE 0 END  REG
		`,
		Mysql: `
		SELECT CASE  WHEN EXISTS(
			SELECT * FROM INFORMATION_SCHEMA.TABLES
			WHERE TABLE_NAME = :TABLENAME
			) > 0 THEN 1 ELSE 0 END REG
		`,
		Sqlser: `
		SELECT CASE  WHEN EXISTS(
			SELECT * FROM INFORMATION_SCHEMA.TABLES
			WHERE TABLE_NAME = :TABLENAME
			) THEN 1 ELSE 0 END REG
		`,
		SQLLite: `
		SELECT CASE  WHEN EXISTS(
			SELECT *
            FROM SQLITE_MASTER
            WHERE TYPE = 'TABLE' AND NAME = :TABLENAME
			) THEN 1 ELSE 0 END REG
		`,
	}
	/*CADCONN : contiene el formato de las cadenas de conexion*/
	CADCONN = map[string]string{
		/*oracle://user:pass@server/service_name*/
		Ora:     "oracle://%s:%s@%s:%d/%s",
		Post:    "postgres://%s:%s@%s:%d/%s?sslmode=%s",
		Mysql:   "%s:%s@tcp(%s:%d)/%s",
		Sqlser:  "server=%s;user id=%s;password=%s;port=%d;database=%s;",
		SQLLite: "%s",
	}
	/*PrefijosDB : contiene los string de conexion al momento de ejecutar la funcion open*/
	PrefijosDB = map[string]string{
		Ora:     "oracle",
		Post:    "postgres",
		Mysql:   "mysql",
		Sqlser:  "mssql",
		SQLLite: "sqlite3",
	}
	/*FORMATTOSTRCONECT : formato to string para la conexion de base de datos*/
	FORMATTOSTRCONECT = "[%s|%s|%s|%d|%s|%s|%s|%s]"
	/*Ssmodes : hace referencia si tienen conexion ssl
	0* disable - No SSL
	1* require - Always SSL (skip verification)
	2* verify-ca - Always SSL (verify that the certificate presented by the
	  server was signed by a trusted CA)
	3* verify-full - Always SSL (verify that the certification presented by
	  the server was signed by a trusted CA and the server host name
	  matches the one in the certificate)

	*/
	Ssmodes = []string{"disable", "require", "verify-ca", "verify-full"}
)
