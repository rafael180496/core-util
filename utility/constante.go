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
