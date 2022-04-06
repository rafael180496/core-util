package test

import (
	"fmt"
	"testing"

	utl "github.com/rafael180496/core-util/utility"
)

/*TestExistFile : Valida si un archivo existe*/
func TestExistFile(t *testing.T) {
	pathorig := "config/pr1/prueba.txt"
	rest := utl.FileExist(pathorig, false)
	fmt.Printf("Result File: %v", rest)
	t.Logf("Result File: %v", rest)
}

/*TestExistDir : Valida si un directorio existe*/
func TestExistDir(t *testing.T) {
	pathorig := "config/pr1"
	rest := utl.FileExist(pathorig, true)
	fmt.Printf("Result Dir: %v", rest)
	t.Logf("Result Dir: %v", rest)
}

/*TestCpfile : Copia un archivo orig a un archivo dest*/
func TestCpfile(t *testing.T) {
	pathorig := "config/pr1/prueba.txt"
	utl.CpFile(pathorig, "config/pr2/")
}

/*TestCpDir : Copia un directorio orig a un archivo dest*/
func TestCpDir(t *testing.T) {
	pathorig := "config/pr1/"
	utl.CpDir(pathorig, "config/pr2/")
}

/*TestRmFile : test para eliminar un archivo*/
func TestRmFile(t *testing.T) {
	pathorig := "config/pr2/prueba.txt"
	utl.RmFile(pathorig)
}

/*TestRmDir : Elimna un directorio completo*/
func TestRmDir(t *testing.T) {
	pathorig := "config/pr2/"
	utl.RmDir(pathorig)
}

/*TestListDir : lista la informacion de un directorio*/
func TestListDir(t *testing.T) {
	pathorig := "config/pr1"
	info, _ := utl.ListDir(pathorig)
	utl.PrintPc("%v", info[0].Name())
}
