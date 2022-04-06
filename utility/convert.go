package utility

import (
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"
)

/*Float64XML : tipo que repara el error de los xml en los millones*/
type Float64XML float64

func jsonStringToObject(s string, v interface{}) error {
	data := []byte(s)
	return json.Unmarshal(data, v)
}

/*ToMapInterface : convierte cualquier data en map [string]interface{}*/
func ToMapInterface(i interface{}) (map[string]interface{}, error) {
	var m = map[string]interface{}{}

	switch v := i.(type) {
	case map[interface{}]interface{}:
		for k, val := range v {
			m[ToString(k)] = val
		}
		return m, nil
	case map[string]interface{}:
		return v, nil
	case string:
		err := jsonStringToObject(v, &m)
		return m, err
	default:
		return m, fmt.Errorf("error converting data to map")
	}
}

/*MarshalXMLAttr : tramforma el xml*/
func (f Float64XML) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	s := fmt.Sprintf("%.2f", f)
	return xml.Attr{Name: name, Value: s}, nil
}

/*TimetoStr : Convierte un time al string con el formato YYYYMMDD que esta en constante.*/
func TimetoStr(dia time.Time) string {
	var (
		str strings.Builder
	)
	fmt.Fprintf(&str, FORMFE, dia.Year(), dia.Month(), dia.Day())
	return str.String()
}

/*StrtoByte : Convierte un string a array byte */
func StrtoByte(str string) []byte {
	return []byte(str)
}

/*BytetoStrHex : Convierte un   array byte a string */
func BytetoStrHex(str []byte) string {
	return hex.EncodeToString(str)
}

/*SliceStrToStr : convierte un array de strings a un string con separador */
func SliceStrToStr(strs []string, separador string) string {
	str := ""
	for i := 0; i < len(strs); i++ {
		if i != (len(strs) - 1) {
			str += strs[i] + separador
		} else {
			str += strs[i]
		}
	}
	return str
}

/*StrtoByteHex : Convierte un string a array byte */
func StrtoByteHex(str string) []byte {
	strByte, _ := hex.DecodeString(str)
	return strByte
}

/*BytetoStr : Convierte un   array byte a string */
func BytetoStr(str []byte) string {
	return string(str)
}

/*RoundFloat64 : redondea un  */
func RoundFloat64(valor float64, dec int) float64 {
	dec = ReturnIf(dec <= 0, 1, dec).(int)
	format := fmt.Sprintf(".%df", dec)
	format = "%" + format
	valstr := fmt.Sprintf(format, valor)
	val, err := StrToFloat64(valstr)
	if err != nil {
		return 0.00
	}
	return val

}

/*StrToFloat64 : Convierte un valor string a float64.*/
func StrToFloat64(val string) (float64, error) {
	Float64, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, err
	}
	return Float64, nil
}

/*Float64ToStr : Convierte un valor float64 a string.*/
func Float64ToStr(val float64) string {
	return strconv.FormatFloat(val, 'f', -1, 64)
}

/*StrToInt64 : Convierte un string a int64 */
func StrToInt64(val string) (int64, error) {
	int64, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, err
	}
	return int64, nil
}

/*Int64ToStr : Convierte un int64 a string.*/
func Int64ToStr(val int64) string {
	return strconv.FormatInt(val, 10)
}

/*StrToInt32 : Convierte un string a int32 */
func StrToInt32(val string) (int32, error) {
	int64, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, err
	}
	return int32(int64), nil
}

/*Int32ToStr : Convierte un int32 a string.*/
func Int32ToStr(val int32) string {
	return strconv.FormatInt(int64(val), 10)
}

/*IntToStr : Convierte un int a string.*/
func IntToStr(val int) string {
	return strconv.FormatInt(int64(val), 10)
}

/*ToDateStr : envia en formato largo un datetime YYYY-MM-DDTHH:MM:SS*/
func ToDateStr(t time.Time) string {
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}

/*ToDateStrSingle : envia en formato largo un datetime YYYY-MM-DD*/
func ToDateStrSingle(t time.Time) string {
	return fmt.Sprintf("%d-%02d-%02d",
		t.Year(), t.Month(), t.Day())
}

/*StringToDate : convierte un string a un time */
func StringToDate(s string) (time.Time, error) {
	return parsedate(s, []string{
		time.RFC3339,
		"2006-01-02T15:04:05",
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC850,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		"2006-01-02 15:04:05.999999999 -0700 MST",
		"2006-01-02",
		"02 Jan 2006",
		"2006-01-02T15:04:05-0700",
		"2006-01-02 15:04:05 -07:00",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05Z07:00",
		"2006-01-02 15:04:05Z0700",
		"2006-01-02 15:04:05",
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	})
}

func parsedate(s string, dates []string) (d time.Time, e error) {
	for _, dateType := range dates {
		if d, e = time.Parse(dateType, s); e == nil {
			return
		}
	}
	return d, fmt.Errorf("Error converting date")
}
