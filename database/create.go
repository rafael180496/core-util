package database

import (
	"bytes"
	"fmt"

	utl "github.com/rafael180496/core-util/utility"
	"gopkg.in/ini.v1"
)

/*DecripConect : desencripta una conexion de base de datos .ini con una encriptacion AES256 creada del mismo
paquete utility*/
func DecripConect(data []byte, pass string) (StCadConect, error) {
	var cnx StCadConect
	dataNew, err := utl.DesencripAES(pass, utl.BytetoStr(data))
	if err != nil {
		return cnx, err
	}
	cnx, err = readIni([]byte(dataNew))
	if err != nil {
		return cnx, err
	}
	return cnx, nil
}

/*CreateDBConect : Crea una conexion de base de datos valida y la genera como un .db con una clave aes*/
func CreateDBConect(cnx StCadConect, pass string) ([]byte, error) {
	pass = utl.Trim(pass)
	if !utl.IsNilStr(pass) {
		return nil, utl.StrErr("La clave esta vacia por favor introducir una clave")
	}
	if !cnx.ValidCad() {
		return nil, utl.StrErr("La conexion no pasa las validaciones.")
	}
	cfg := ini.Empty()
	sec, err := cfg.NewSection("database")
	if err != nil {
		return nil, err
	}
	err = sec.ReflectFrom(&cnx)
	if err != nil {
		return nil, err
	}
	cfg.DeleteSection("DEFAULT")
	var buf bytes.Buffer
	cfg.WriteTo(&buf)
	data := buf.String()
	dataencrip, err := utl.EncripAES(pass, data)
	if err != nil {
		return nil, err
	}
	return utl.StrtoByte(dataencrip), nil
}

/*CreateDbFile : crea un archivo de configuracion valida para base de datos encriptado*/
func CreateDbFile(cnx StCadConect, pass, dir, name string) error {
	if !utl.FileExist(dir, true) {
		return fmt.Errorf("el directorio destino no existe")
	}
	dir = utl.PlecaAdd(dir)
	f, err := utl.FileNew(dir + name + utl.EXT["DBX"])
	if err != nil {
		return err
	}
	data, err := CreateDBConect(cnx, pass)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}
