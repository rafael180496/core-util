package test

import (
	"testing"

	utl "github.com/rafael180496/core-util/utility"
)

/*TestMsjBlue : Envie un string en color celeste*/
func TestMsjBlue(t *testing.T) {
	t.Logf("Texto:%s", utl.MsjBlue("Hola"))
}

/*TestMsjRed : Envie un string en color rojo*/
func TestMsjRed(t *testing.T) {
	t.Logf("Texto:%s", utl.MsjRed("Hola"))
}

/*TestMsjGreen : Envie un string en color verde*/
func TestMsjGreen(t *testing.T) {
	t.Logf("Texto:%s\n", utl.MsjGreen("Hola"))
	utl.PrintPc(utl.Green, "Texto:Hola\n")
}

/*TestMsjGreen : Envie un string en color verde*/
func TestIP4(t *testing.T) {
	t.Logf("IPLOCAL:%s\n", utl.GetLocalIPV4())
}

/*TestMsjPc : prueba todos los texto disponible */
func TestMsjPc(t *testing.T) {
	t.Logf("Texto:%s", utl.MsjPc(utl.Green, "%s", utl.Green))
	t.Logf("Texto:%s", utl.MsjPc(utl.Red, "%s", utl.Red))
	t.Logf("Texto:%s", utl.MsjPc(utl.Blue, "%s", utl.Blue))
	t.Logf("Texto:%s", utl.MsjPc(utl.Cyan, "%s", utl.Cyan))
	t.Logf("Texto:%s", utl.MsjPc(utl.White, "%s", utl.White))
	t.Logf("Texto:%s", utl.MsjPc(utl.Black, "%s", utl.Black))
	t.Logf("Texto:%s", utl.MsjPc(utl.Yellow, "%s", utl.Yellow))
	t.Logf("Texto:%s", utl.MsjPc(utl.Magenta, "%s", utl.Magenta))
	t.Logf("Texto:%s", utl.MsjPc(utl.HBlack, "%s", utl.HBlack))
	t.Logf("Texto:%s", utl.MsjPc(utl.HRed, "%s", utl.HRed))
	t.Logf("Texto:%s", utl.MsjPc(utl.HGreen, "%s", utl.HGreen))
	t.Logf("Texto:%s", utl.MsjPc(utl.HYellow, "%s", utl.HYellow))
	t.Logf("Texto:%s", utl.MsjPc(utl.HBlue, "%s", utl.HBlue))
	t.Logf("Texto:%s", utl.MsjPc(utl.HMagenta, "%s", utl.HMagenta))
	t.Logf("Texto:%s", utl.MsjPc(utl.HCyan, "%s", utl.HCyan))
}
