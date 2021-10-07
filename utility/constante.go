package utility

import "regexp"

type (
	/*Pc : es un tipo para las paletas de colores.*/
	Pc string
)

var (
	/*URLFor : valida la expresion de cadenas */
	URLFor = regexp.MustCompile(URL)
	/*EmailFor : expresion para validar email*/
	EmailFor = regexp.MustCompile(Email)
	/*TypeContent : contents type para los formatos de correo */
	TypeContent = map[string]string{
		"html": "text/html",
		"text": "text/plain",
	}
	/*SMTPPORT : indica los puertos disponible para el smtp */
	SMTPPORT = map[string]string{
		"gmail1":   "587",
		"hotmail1": "25",
		"hotmail2": "465",
		"yahoo1":   "465",
		"yahoo2":   "587",
	}
	/*SMTPURL :  Host de correos */
	SMTPURL = map[string]string{
		"gmail":   "smtp.gmail.com",
		"hotmail": "smtp.live.com",
		"yahoo":   "smtp.mail.yahoo.com",
	}

	/*Msj : mensajes de error en las librerias*/
	Msj = StMsj{
		Store: map[string]string{
			"GE01": "Mensaje no encontrado.",
			"GE02": "Error al convertir fecha",
			"GE03": "Error al convertir la data a map",
			//libreria encrip
			"EN01": "Error en generar phisher.",
			"EN02": "Error en generar bloque.",
			"EN03": "Error en leer el bloque.",
			"EN04": "Error en abrir el bloque.",
			//libreria convert
			"CO01": "Error en convertir float.",
			"CO02": "Error en convertir int64.",
			"CO03": "Error en convertir int32.",
			//libreria archivo,
			"AR01": "No existe el archivo.",
			"AR02": "Error al crear directorio.",
			"AR03": "Error al crear archivo.",
			"AR04": "Error al crear log.",
			"AR05": "No existe el directorio.",
			"AR06": "Problemas en leer directorio.",
			"AR07": "Error al eliminar un archivo.",
			//libreria database
			"CN01": "Error al convertir a json.",
			"CN02": "Error al convertir a string.",
			"CN03": "Error al convertir a string.",
			"CN04": "No tiene el insert.",
			"CN05": "No tiene los prefijos validos.",
			"CN06": "No se obtuvieron las columnas.",
			"CN07": "No se obtuvieron las filas .",
			"CN08": "Error al leer archivo config json.",
			"CN09": "No existe el arhivo config json.",
			"CN10": "No existe el arhivo config ini.",
			"CN11": "Error al leer archivo config ini.",
			"CN12": "El archivo config no es valido.",
			"CN13": "Tipo de DB no compatible.",
			"CN14": "Error al conectar ala DB.",
			"CN15": "Cantidad de fila es cero.",
			"CN16": "Error al cerrar ala DB.",
			"CN18": "Error en al ejecutar querie.",
			"CN19": "Error en al obtener columna.",
			"CN20": "El db es sqllite necesita el archivo.db",
			"CN21": "La URL no es valida o le falta parametros",
			"CN22": "Cantiadad de acciones menor o igual a ceros",
			"CN23": "Base de datos sqllite vacia",
			//libreria json
			"JS01": "Error al generar un nuevo json.",
			"JS02": "Error al convertir un json.",
			"JS03": "Json null.",
			"JS04": "Error al capturar json.",
			//Libreria de Email
			"ER01": "Error en enviar el mensaje",
			"ER02": "No pasa las validaciones",
			/*libreria server */
			"SR01": "Tipo  de peticion incorrecta.",
			"SR02": "El Path es invalido.",
			"SR03": "No contiene una funcion de peticion.",
			/*libreria database*/
			"DT01": "El numero de fila supera la cantidad total de filas.",
			"DT02": "La columna no existe.",
			"DT03": "No tienen tabla cargada.",
			"DT04": "No tienen columnas cargada.",
			"DT05": "Columna duplicada.",
			"DT06": "No tiene ningun indice asignado",
			"DT07": "Accion invalida datatable",
		},
	}
	/*EXT : extensiones de archivos */
	EXT = map[string]string{
		"JSON": ".json",
		"INI":  ".ini",
		"XML":  ".xml",
		"RUT":  ".rut",
		"SQL":  ".sql",
		"DB":   ".db",
		"CSV":  ".csv",
		"TXT":  ".txt",
		"DBX":  ".dbx",
	}
)

const (
	/*Expresiones numericas para validar formatos de string*/

	/*Email : valida la expresion valida en una cadena de caracteres*/
	Email string = "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	/*URLPath : formato de los path de las url*/
	URLPath string = `((\/|\?|#)[^\s]*)`
	/*URLPort : puertos validos de una url*/
	URLPort string = `(:(\d{1,5}))`
	/*URLIP : estructura de una ip publica*/
	URLIP string = `([1-9]\d?|1\d\d|2[01]\d|22[0-3]|24\d|25[0-5])(\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-5]))`
	/*URLSubdomain : dominios validos de un pagina*/
	URLSubdomain string = `((www\.)|([a-zA-Z0-9]+([-_\.]?[a-zA-Z0-9])*[a-zA-Z0-9]\.[a-zA-Z0-9]+))`
	/*IP : valida el formato de un string*/
	IP string = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
	/*URLUsername : valida los url con usuario*/
	URLUsername string = `(\S+(:\S*)?@)`
	/*URLSchema : valida el esquema de una url*/
	URLSchema string = `((ftp|tcp|udp|wss?|https?):\/\/)`
	/*URL : Formato para validar cadenas de url*/
	URL = `^` + URLSchema + `?` + URLUsername + `?` + `((` + URLIP + `|(\[` + IP + `\])|(([a-zA-Z0-9]([a-zA-Z0-9-_]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|(` + URLSubdomain + `?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))\.?` + URLPort + `?` + URLPath + `?$`
	/*MaxURLRuneCount : maxima cantidad de runas por contar*/
	MaxURLRuneCount = 2083
	/*MinURLRuneCount : minima cantidad de runas por contar*/
	MinURLRuneCount = 3

	/*colores disponibles */

	/*Green : verde */
	Green Pc = "g"
	/*Red : rojo */
	Red Pc = "r"
	/*Blue : azul */
	Blue Pc = "b"
	/*Cyan : celeste */
	Cyan Pc = "c"
	/*White : blanco */
	White Pc = "w"
	/*Black : negro */
	Black Pc = "bl"
	/*Yellow : amarillo*/
	Yellow Pc = "y"
	/*Magenta : magenta */
	Magenta Pc = "m"
	/*HBlack : negro fuerte */
	HBlack Pc = "hbl"
	/*HRed : rojo fuerte */
	HRed Pc = "hr"
	/*HGreen : verde fuerte */
	HGreen Pc = "hg"
	/*HYellow : amarrillo fuerte */
	HYellow Pc = "hy"
	/*HBlue : azul fuerte */
	HBlue Pc = "hb"
	/*HMagenta : magenta fuerte*/
	HMagenta Pc = "hm"
	/*HCyan : celeste fuerte */
	HCyan Pc = "hc"
	/*HWhite : blanco fuerte */
	HWhite Pc = "hw"
	/*FORMFE : Formato de fecha para los archivo YYYYMMDD*/
	FORMFE = "%d%02d%02d"
)
