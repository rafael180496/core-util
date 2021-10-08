package server

import (
	"net/http"
	"regexp"

	echo "github.com/labstack/echo/v4"
)

type (
	/*Request : estructara para tener los datos del servicio*/
	Request struct {
		Method      HTTPTip
		BaseURL     string
		Headers     map[string]string
		QueryParams map[string]string
		Body        []byte
	}
	/*Response : estructura para cargar la respuesta del servidor*/
	Response struct {
		StatusCode int
		Body       string
		Headers    map[string][]string
	}
	/*RestError : Obtiene una respuesta incorrecta al lado del servidor*/
	RestError struct {
		Response *Response
	}
	/*Client : cliente http*/
	Client struct {
		HTTPClient *http.Client
	}
	/*section : contiene la informacion general del servicio*/
	section struct {
		name    string
		version string
		comment []string
	}
	/*Browser : detecta el navegador que esta ocupando*/
	Browser struct {
		Engine        string
		EngineVersion string
		Name          string
		Version       string
	}
	/*UserAgent : Estructura para capturar el useragent de la peticion*/
	UserAgent struct {
		UA           string
		Mozilla      string
		Platform     string
		OS           string
		Localization string
		Browser      Browser
		Bot          bool
		Mobile       bool
		Undecided    bool
	}
	/*ConfigServer : configuraciones del servicio generales */
	ConfigServer struct {
		/*Ipser : sirve para cuando este en modo produccion coloque la ip asignada al server*/
		Ipser string `ini:"ipser"`
		/*Debug : modo debug para el proyecto*/
		Debug bool `ini:"debug"`
		/*Local : habilita el modo local del servicio*/
		Local bool `ini:"local"`
		/*Puerto : Coloca el puerto del servicio*/
		Puerto int `ini:"puerto"`
		/*Protocol : protocolo del servicio*/
		Protocol string `ini:"protocol"`
		/*CertFile : certificado ssl*/
		CertFile string `ini:"certfile"`
		/*KeyFile : llave del certificado ssl*/
		KeyFile string `ini:"keyfile"`
		/*DirSSL : carpeta donde esta el certificado*/
		DirSSL string `ini:"dirssl"`
		/*Env :  carga todas las variables en variables de entorno*/
		Env bool `ini:"env"`
	}
	/*StInfoPet : envia la informacion general de la peticion */
	StInfoPet struct {
		URL    string `json:"url"`
		Method string `json:"method"`
	}
	/*StInfoReq :  Obtiene  datos generales del request*/
	StInfoReq struct {
		HostOrig  string `json:"host"`
		IPRemote  string `json:"ip"`
		Browser   string `json:"browser"`
		SystemOI  string `json:"oi"`
		UserAgent string `json:"useragent"`
	}
	/*StDataEnv : Envio de datos generico*/
	StDataEnv struct {
		Code  int         `json:"code"`
		Error string      `json:"msj"`
		Data  interface{} `json:"data"`
	}
	/*HTTPTip : especifica el tipo de peticion a elaborar*/
	HTTPTip string
	/*Controller : estructura para crear Controladores */
	Controller struct {
		Pets []HTTPPet
	}

	/*HTTPPet Guarda una peticion HTTP basada en ECHO */
	HTTPPet struct {
		Path string
		Pet  func(echo.Context) error
		Tip  HTTPTip
	}
)

var (
	/*DefaultClient : cliente http con parametros default*/
	DefaultClient     = &Client{HTTPClient: &http.Client{}}
	ie11Regexp        = regexp.MustCompile("^rv:(.+)$")
	botRegex          = regexp.MustCompile("(?i)(bot|crawler|sp(i|y)der|search|worm|fetch|nutch)")
	botFromSiteRegexp = regexp.MustCompile(`http[s]?://.+\.\w+`)
	sectionBrow       = map[string]string{
		"CriOS":    "Chrome",
		"Chrome":   "Chrome",
		"Chromium": "Chromium",
		"FxiOS":    "Firefox",
		"Safari":   "Safari",
	}
	osWindow = map[string]string{
		"5.0":  "Windows 2000",
		"5.01": "Windows 2000, Service Pack 1 (SP1)",
		"5.1":  "Windows XP",
		"5.2":  "Windows XP x64 Edition",
		"6.0":  "Windows Vista",
		"6.1":  "Windows 7",
		"6.2":  "Windows 8",
		"6.3":  "Windows 8.1",
		"10.0": "Windows 10",
	}
)

const (
	/*Peticiones http*/

	/*POST : peticion http POST*/
	POST HTTPTip = "POST"
	/*GET : peticion http GET*/
	GET HTTPTip = "GET"
	/*PUT : peticion http PUT*/
	PUT HTTPTip = "PUT"
	/*DELETE : peticion http DELETE*/
	DELETE HTTPTip = "DELETE"
	/*Tipos de protocolos : http https */
	PATCH HTTPTip = "PATCH"
	/*Tipos de protocolos : http https */
	ANY HTTPTip = "ANY"

	/*HTTP : protocolo http*/
	HTTP = "http"
	/*HTTPS : protocolo https*/
	HTTPS = "https"
)
