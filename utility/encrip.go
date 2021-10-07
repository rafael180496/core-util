package utility

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"time"
)

/*GeneredHashSha256 : Genera un hash con encriptacion sha256 */
func GeneredHashSha256(key string) string {
	hash := sha256.New()
	io.WriteString(hash, key)
	text := hex.EncodeToString(hash.Sum(nil))
	return text
}

/*GeneredHashMd5 : Genera un hash con encriptacion md5 */
func GeneredHashMd5(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

/*EncripAES : Encripta en aes 256 un texto donde la key es la llave noce una llave secundaria y el texto a ecriptar */
func EncripAES(key string, text string) (string, error) {
	textByte := StrtoByte(text)
	bloque, err := aes.NewCipher([]byte(GeneredHashMd5(key)))
	if err != nil {
		fmt.Println(err.Error())
		return "", Msj.GetError("EN01")
	}
	gcm, err2 := cipher.NewGCM(bloque)
	if err2 != nil {
		return "", Msj.GetError("EN02")
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err3 := io.ReadFull(rand.Reader, nonce); err3 != nil {
		return "", Msj.GetError("EN03")
	}
	textoEncrip := gcm.Seal(nonce, nonce, textByte, nil)

	return BytetoStrHex(textoEncrip), nil
}

/*DesencripAES : Desencripta en aes 256 un texto donde la key es la llave noce una llave secundaria y el texto a ecriptar */
func DesencripAES(key string, text string) (string, error) {
	textByte := StrtoByteHex(text)
	keyHash := []byte(GeneredHashMd5(key))
	bloque, err := aes.NewCipher(keyHash)
	if err != nil {
		return "", Msj.GetError("EN01")
	}
	gcm, err2 := cipher.NewGCM(bloque)
	if err2 != nil {
		return "", Msj.GetError("EN02")
	}
	nonceSize := gcm.NonceSize()
	if nonceSize > len(textByte) {
		return "", Msj.GetError("EN04")
	}
	nonce, ciphertext := textByte[:nonceSize], textByte[nonceSize:]
	textoDesencrip, err3 := gcm.Open(nil, nonce, ciphertext, nil)
	if err3 != nil {
		return "", Msj.GetError("EN04")
	}

	return BytetoStr(textoDesencrip), nil
}

/*GenToken : Genera un token dependiendo de un string.*/
func GenToken(str string) string {
	return GeneredHashSha256(StrRand(len(str), false) + string(rune(time.Now().Second())))
}

/*GeneredUUID : Genera un codigo uuid unico */
func GeneredUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
