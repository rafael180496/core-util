package utility

import "time"

/*FNulo : Genera una fecha nula 29991231*/
func FNulo() time.Time {
	return time.Date(2999, 12, 31, 0, 0, 0, 0, time.Now().Location())
}

/*FActual : devuelve la fecha de hoy redondeada. */
func FActual() time.Time {
	toRound := time.Now()
	return time.Date(toRound.Year(), toRound.Month(), toRound.Day(), 0, 0, 0, 0, toRound.Location())
}

/*DateEquals :
valida  las fechas
b fecha base
v fecha a validar */
func DateEquals(b time.Time, v time.Time) bool {
	switch {
	case b.Year() > v.Year():
		return false
	case b.Month() > v.Month() && b.Year() == v.Year():
		return false
	case b.Day() > v.Day() && b.Month() == v.Month():
		return false
	}
	return true
}

/*
DateIdent : Valida si las fechas son identicas.
b fecha base
v fecha a validar
*/
func DateIdent(b time.Time, v time.Time) bool {
	if b.Year() == v.Year() && b.Month() == v.Month() && b.Day() == v.Day() {
		return true
	}
	return false
}

/*
TimeEqual : valida los tiempos
b fecha base
v fecha a validar
*/
func TimeEqual(b time.Time, v time.Time) bool {
	switch {
	case b.Hour() > v.Hour():
		return false
	case b.Minute() > v.Minute() && b.Hour() == v.Hour():
		return false
	}
	return true
}
