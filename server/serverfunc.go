package server

import (
	"fmt"
	"os"
	"strings"

	utl "github.com/rafael180496/core-util/utility"
	"gopkg.in/ini.v1"

	echo "github.com/labstack/echo/v4"
)

/*InfPetKey : captura la informacion de una peticion y un token valido*/
func InfPetKey(c echo.Context) (StInfoPet, string) {
	info := StInfoPet{}
	key := ExtKey(c)
	info.Method = c.Request().Method
	info.URL = c.Path()
	return info, key
}

/*InfPet : captura la informacion de una peticion*/
func InfPet(c echo.Context) StInfoPet {
	info := StInfoPet{}
	info.Method = c.Request().Method
	info.URL = c.Path()
	return info
}

/*ResultJSON : envia un json de los resultado del api rest */
func (p *StDataEnv) ResultJSON(e echo.Context) error {
	return e.JSON(p.Code, p)
}

/*ResultCodeJSON : envia un json de los resultado del api rest pero asignando el codigo */
func (p *StDataEnv) ResultCodeJSON(code int, e echo.Context) error {
	p.Code = code
	return e.JSON(p.Code, p)
}

/*ResultFullJSON : envia un json de los resultado del api rest pero asignando el codigo la data y el mensaje */
func (p *StDataEnv) ResultFullJSON(code int, msjerr string, data interface{}, e echo.Context) error {
	p.Code = code
	p.Error = msjerr
	p.Data = data
	return e.JSON(p.Code, p)
}

/*JSONSend : envia un json generic  */
func JSONSend(data StDataEnv, e echo.Context) error {
	return e.JSON(data.Code, data)
}

/*FindInfoReq : Envia la informacion de un request despues de una peticion */
func FindInfoReq(c echo.Context) (StInfoReq, error) {
	var (
		info StInfoReq
	)
	ua := New(c.Request().UserAgent())
	info.UserAgent = c.Request().UserAgent()
	info.HostOrig = c.Scheme() + "://" + c.Request().Host
	info.IPRemote = c.RealIP()
	info.Browser = fmt.Sprintf("%s %s", ua.Browser.Name, ua.Browser.Version)
	info.SystemOI = ua.OS
	return info, nil
}

/*Valid : Valida si la peticon es valida*/
func (p *HTTPPet) Valid() error {
	if !p.Tip.Valid() {
		return utl.Msj.GetError("SR01")
	}
	if utl.Trim(p.Path) == "" || utl.IsSpace(p.Path) {
		return utl.Msj.GetError("SR02")
	}
	if p.Pet == nil {
		return utl.Msj.GetError("SR03")
	}
	return nil
}

/*Valid : Valida si un tipo de perticion es valida.*/
func (p *HTTPTip) Valid() bool {
	switch *p {
	case POST, GET, PUT, DELETE, PATCH, ANY:
		return true
	default:
		return false
	}
}

/*AsigServer : asigna peticiones a una instancia ECHO */
func AsigServer(e *echo.Echo, cs []Controller) error {
	for _, c := range cs {
		err := asigpet(e, c.Pets, "")
		if err != nil {
			return err
		}
	}
	return nil
}

/*AsigServerGroup : crea sub asignaciones o grupos de path */
func AsigServerGroup(e *echo.Echo, cs []Controller, group string) error {
	for _, c := range cs {
		err := asigpet(e, c.Pets, group)
		if err != nil {
			return err
		}
	}
	return nil
}

/*asigpet : asigna las peticiones y las valida */
func asigpet(e *echo.Echo, ps []HTTPPet, group string) error {
	if utl.Trim(group) != "" {
		group = "/" + utl.Trim(group)
	}
	for _, p := range ps {
		err := p.Valid()
		if err != nil {
			return err
		}
		err = findpet(e, p, group)
		if err != nil {
			return err
		}
	}
	return nil
}

/*findpet : busca la peticion para asignar*/
func findpet(e *echo.Echo, p HTTPPet, group string) error {
	path := utl.ReturnIf(group != "", group+p.Path, p.Path).(string)
	switch p.Tip {
	case POST:
		e.POST(path, p.Pet)
		return nil
	case PUT:
		e.PUT(path, p.Pet)
		return nil
	case GET:
		e.GET(path, p.Pet)
		return nil
	case DELETE:
		e.DELETE(path, p.Pet)
		return nil
	case PATCH:
		e.PATCH(path, p.Pet)
		return nil
	case ANY:
		e.Any(path, p.Pet)
		return nil
	default:
		return utl.Msj.GetError("SR01")
	}
}

/*ExtKey : extrae el key de la sesion en el api */
func ExtKey(c echo.Context) string {
	return strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", -1)
}

/*LoadIni : leer el archivo de configuracion del servicio esta debe ser [server]*/
func (p *ConfigServer) LoadIni(path string) error {
	Config := *p
	cfg, erraux := ini.Load(path)
	if erraux != nil {
		return erraux
	}
	err := cfg.Section("server").MapTo(&Config)
	if err != nil {
		return err
	}
	*p = Config
	return p.Valid()
}

/*loadEnv : carga algunas variables por medio de variables de entorno*/
func (p *ConfigServer) loadEnv() error {
	var err error
	utl.Block{
		Try: func() {
			p.Puerto = utl.ToInt(os.Getenv("PORT"))
		},
		Catch: func(e utl.Exception) {
			err = fmt.Errorf("error al capturar una variable en paquete server")
		},
	}.Do()
	return err
}

/*Valid : valida la estructa y la configura para el servicio*/
func (p *ConfigServer) Valid() error {
	if p.Env {
		p.loadEnv()
	}
	p.Ipser = utl.Trim(p.Ipser)
	if p.Ipser == "" {
		p.Ipser = utl.GetLocalIPV4()
	}
	if p.Local {
		p.Ipser = "localhost"
	}
	if p.Puerto <= 0 {
		return utl.StrErr("Puerto no valido.")
	}
	if !utl.InStr(p.Protocol, HTTP, HTTPS) {
		p.Protocol = HTTP
	}
	if p.Protocol == HTTPS {
		if !utl.FileExist(p.DirSSL, true) {
			return utl.StrErr("Directorio SSL no existe.")
		}
		pathcert := utl.PlecaAdd(p.DirSSL) + p.CertFile
		if !utl.FileExist(pathcert, false) {
			return utl.StrErr("Archivo del certificado no existe.")
		}
		pathkey := utl.PlecaAdd(p.DirSSL) + p.KeyFile
		if !utl.FileExist(pathkey, false) {
			return utl.StrErr("Archivo de la llave no existe.")
		}
	}
	return nil
}

/*Host : envia el host del servicio*/
func (p *ConfigServer) Host() string {
	if p.Local {
		return fmt.Sprintf("%s://%s:%d", p.Protocol, p.Ipser, p.Puerto)
	}
	return fmt.Sprintf("%s://%s", p.Protocol, p.Ipser)
}

/*PathKey : recupera el path de la llave*/
func (p *ConfigServer) PathKey() string {
	return utl.PlecaAdd(p.DirSSL) + p.KeyFile
}

/*PathCert : recupera el path del certificado*/
func (p *ConfigServer) PathCert() string {
	return utl.PlecaAdd(p.DirSSL) + p.CertFile
}

/*StarServer : inicia el servicio echo .*/
func (p *ConfigServer) StarServer(e *echo.Echo) {
	if p.Protocol == HTTPS {
		e.Logger.Fatal(e.StartTLS(":"+utl.IntToStr(p.Puerto), p.PathCert(), p.PathKey()))
	} else {
		e.Logger.Fatal(e.Start(":" + utl.IntToStr(p.Puerto)))
	}
}
