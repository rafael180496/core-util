package test

import (
	"testing"

	utl "github.com/rafael180496/core-util/utility"
)

/*TestAes : prueba las funciones de encriptamiento aes */
func TestAes(t *testing.T) {
	key := "abc123"
	text := "hola mundo"
	t.Logf("Key:%s,Text:%s", key, text)
	bloque, err := utl.EncripAES(key, text)
	if err != nil {
		t.Errorf("Error:%s", err.Error())
	}
	t.Logf("BloqueEncrip:%s", bloque)
	bloque, err = utl.DesencripAES(key, bloque)
	if err != nil {
		t.Errorf("Error:%s", err.Error())
	}
	t.Logf("BloqueDesencrip:%s", bloque)
}

/*TestMd5 : Genera un hash Md5 */
func TestMd5(t *testing.T) {
	key := "abc123"
	t.Logf("key:%s", key)
	t.Logf("bloquemd5:%s", utl.GeneredHashMd5(key))
}

/*Test256 : Genera un hash sha256 */
func Test256(t *testing.T) {
	key := "abc123"
	t.Logf("key:%s", key)
	t.Logf("bloque256:%s", utl.GeneredHashSha256(key))
}

/*TestToken : Genera un token con hash sha 256 unico */
func TestToken(t *testing.T) {
	key := "abc123"
	t.Logf("key:%s", key)
	t.Logf("bloquetoken:%s", utl.GenToken(key))
}

/*TestUUID : Genera una clave unica*/
func TestUUID(t *testing.T) {
	text := utl.GeneredUUID()
	t.Logf("text:[%s]", text)
}
