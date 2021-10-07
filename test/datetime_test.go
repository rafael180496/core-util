package test

import (
	"testing"
	"time"

	"github.com/rafael180496/core-util/utility"
)

/*TestFNulo : Envie una fecha nula.*/
func TestFActual(t *testing.T) {
	t.Logf("Factual:%s", utility.FActual())
}

/*TestFNulo : Envie una fecha nula.*/
func TestFNulo(t *testing.T) {
	t.Logf("NULO:%s", utility.FNulo())
}

/*TestTimetoString : Convierte un date en string */
func TestTimetoString(t *testing.T) {
	tiempo := time.Now()
	t.Logf("Tiempo:%s", tiempo)
	t.Logf("TiempoStr:%s", utility.TimetoStr(tiempo))
}
