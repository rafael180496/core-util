package utility

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

type (
	/*Date : formato Time especial con parseo los xml y json de manera que el formato puede variar*/
	Date struct {
		time.Time
	}
	/*HrTime : valida las horas en un time*/
	HrTime struct {
		hr, min int
	}
)

/*NewTicker : Crea un nuevo ticker por medio de formatos
hr -> horas
min -> minutos
*/
func NewTicker(ft, tp string) (*time.Ticker, error) {
	var err error
	timeTicker := new(time.Ticker)
	Block{
		Try: func() {
			dataint := ToInt(ft)
			switch tp {
			case "hr":
				timeTicker = time.NewTicker(time.Duration(dataint) * time.Hour)
			case "min":
				timeTicker = time.NewTicker(time.Duration(dataint) * time.Minute)
			default:
				err = fmt.Errorf("tipo no es valido por favor solo hr,min")
			}
		},
		Catch: func(e Exception) {
			err = fmt.Errorf("error en conversion del formato")
		},
	}.Do()
	return timeTicker, err
}

/*NewHrTimeStr : crea un time con formato de hora 24hr ejemplo 24:00*/
func NewHrTimeStr(fr string) (HrTime, error) {
	var (
		err    error
		hrtime HrTime
	)
	Block{
		Try: func() {
			hr := ToInt(SubString(fr, 0, 2))
			min := ToInt(SubString(fr, 3, 5))
			hrtime, err = NewHrTime(hr, min)
		},
		Catch: func(e Exception) {
			err = fmt.Errorf("error en conversion del formato")
		},
	}.Do()
	return hrtime, err
}

/*Valid : valida la estructura*/
func (p *HrTime) Valid() error {
	if p.hr > 24 || p.hr < 0 {
		return StrErr("la hora es invalida colocar en formato de 24HR")
	}
	if p.min > 59 || p.min < 0 {
		return StrErr("los Minutos es invalida colocar en formato de 59Min")
	}
	return nil
}

/*SetHr : setea un tipo time */
func (p *HrTime) SetHr(hr, min int) error {
	val := HrTime{
		hr:  hr,
		min: min,
	}
	if err := val.Valid(); err != nil {
		return err
	}
	*p = val
	return nil
}

/*GetHr : extrae un tipo time */
func (p *HrTime) GetHr() (hr, min int) {
	return p.hr, p.min
}

/*SetHrTime : setea un tipo time */
func (p *HrTime) SetHrTime(hr, min int) error {
	val := HrTime{
		hr:  hr,
		min: min,
	}
	if err := val.Valid(); err != nil {
		return err
	}
	*p = val
	return nil
}

/*NewHrTime : crea un time */
func NewHrTime(hr, min int) (HrTime, error) {
	val := HrTime{
		hr:  hr,
		min: min,
	}
	if err := val.Valid(); err != nil {
		return HrTime{}, err
	}
	return val, nil
}

/*NewHrTimeDate : crea un time  con un tipo date*/
func NewHrTimeDate(data Date) (HrTime, error) {
	val := HrTime{
		hr:  data.Hour(),
		min: data.Minute(),
	}
	if err := val.Valid(); err != nil {
		return HrTime{}, err
	}
	return val, nil
}

/*EqualNow : Valida si un hrtime es identico a un date de la hora actual*/
func (p *HrTime) EqualNow() bool {
	date := time.Now()
	return ReturnIf(date.Hour() == p.hr && date.Minute() == p.min, true, false).(bool)
}

/*EqualDate : Valida si un hrtime es identico a un date*/
func (p *HrTime) EqualDate(date Date) bool {
	return ReturnIf(date.Hour() == p.hr && date.Minute() == p.min, true, false).(bool)
}

/*HrTime : crea una instancia de hrtime*/
func (p *Date) HrTime() HrTime {
	return HrTime{
		hr:  p.Hour(),
		min: p.Minute(),
	}
}

/*NewDateNow : crear un tipo date con la fecha actual*/
func NewDateNow() Date {
	date := Date{
		Time: time.Now(),
	}
	return date
}

/*ToString : lo convierte  un date en string */
func (p *Date) ToString() string {
	return TimetoStr(p.Time)
}

/*Equals : Valida dos date a -> b

1 -> mayor a

0 -> igual b

-1 -> menor a
*/
func (p *Date) Equals(data Date) int {
	switch {
	case p.Year() == data.Year() && p.Month() == data.Month() && p.Day() == data.Day():
		return 0
	case p.Year() > data.Year(), p.Year() == data.Year() && p.Month() > data.Month(), p.Year() == data.Year() && p.Month() == data.Month() && p.Day() > data.Day():
		return 1
	default:
		return -1
	}
}

/*UnmarshalJSON : fomateo especial a json para los tipo Time en golang*/
func (p *Date) UnmarshalJSON(input []byte) error {
	newTime, err := StringToDate(strings.Trim(string(input), `"`))
	if err != nil {
		return err
	}
	p.Time = newTime
	return nil
}

/*UnmarshalXML : fomateo especial a xml para los tipo Time en golang*/
func (p *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parse, err := StringToDate(v)
	if err != nil {
		return err
	}
	*p = Date{parse}
	return nil
}
