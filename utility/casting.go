package utility

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

/*MapKeys : Conviert un map a interface o datos dinamicos.*/
func MapKeys(data *map[interface{}]interface{}) []interface{} {
	keys := make([]interface{}, len(*data))
	i := 0
	for key := range *data {
		keys[i] = key
		i++
	}
	return keys
}

/*MapStrKeys : Convierte un map a arreglo de string*/
func MapStrKeys(data *map[string]interface{}) []string {
	keys := make([]string, len(*data))
	i := 0
	for key := range *data {
		keys[i] = key
		i++
	}
	return keys

}

/*JsonStructMapData : convierte un struct con tag json a map interface*/
func JsonStructMapData(data interface{}) (map[string]interface{}, error) {
	jsonIn, err := NewJSON(data)
	if err != nil {
		return nil, err
	}
	mapData, err := JSONtoObj(jsonIn)
	if err != nil {
		return nil, err
	}
	if len(mapData) <= 0 {
		return nil, fmt.Errorf("empty map")
	}
	return mapData[0], nil
}

/*StructToMap : Convierte un struct a map[string] */
func StructToMap(i interface{}) (valores url.Values) {
	valores = url.Values{}
	atributo := reflect.ValueOf(i).Elem()
	tipo := atributo.Type()
	for i := 0; i < atributo.NumField(); i++ {
		f := atributo.Field(i)
		var v string
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()
		}
		valores.Set(tipo.Field(i).Name, v)
	}
	return
}

/*AsignarPtr : asigna la data despues de referencias en punteros muchas veces*/
func AsignarPtr(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

/*ToTime : Convierte un interface a time*/
func ToTime(data interface{}) (tim time.Time, err error) {
	data = AsignarPtr(data)
	switch v := data.(type) {
	case time.Time:
		return v, nil
	case string:
		return StringToDate(v)
	case int:
		return time.Unix(int64(v), 0), nil
	case int64:
		return time.Unix(v, 0), nil
	case int32:
		return time.Unix(int64(v), 0), nil
	case uint:
		return time.Unix(int64(v), 0), nil
	case uint64:
		return time.Unix(int64(v), 0), nil
	case uint32:
		return time.Unix(int64(v), 0), nil
	default:
		return time.Time{}, fmt.Errorf("Error converting date")
	}
}

/*ToBoolean : convierte un interface a bool*/
func ToBoolean(data interface{}) bool {
	data = AsignarPtr(data)
	switch b := data.(type) {
	case bool:
		return b
	case nil:
		return false
	case int:
		if data.(int) != 0 {
			return true
		}
		return false
	case string:
		resp, _ := strconv.ParseBool(data.(string))
		return resp
	default:
		return false
	}
}

/*ToFloat64 : convierte un interface a float64*/
func ToFloat64(i interface{}) float64 {
	i = AsignarPtr(i)
	switch s := i.(type) {
	case float64:
		return s
	case float32:
		return float64(s)
	case int:
		return float64(s)
	case int64:
		return float64(s)
	case int32:
		return float64(s)
	case int16:
		return float64(s)
	case int8:
		return float64(s)
	case uint:
		return float64(s)
	case uint64:
		return float64(s)
	case uint32:
		return float64(s)
	case uint16:
		return float64(s)
	case uint8:
		return float64(s)
	case string:
		v, err := strconv.ParseFloat(s, 64)
		return ReturnIf(err == nil, v, 0).(float64)
	case bool:
		return ReturnIf(s, 1, 0).(float64)
	default:
		return 0
	}
}

/*ToFloat : convierte un interface a float32*/
func ToFloat(i interface{}) float32 {
	i = AsignarPtr(i)
	switch s := i.(type) {
	case float64:
		return float32(s)
	case float32:
		return s
	case int:
		return float32(s)
	case int64:
		return float32(s)
	case int32:
		return float32(s)
	case int16:
		return float32(s)
	case int8:
		return float32(s)
	case uint:
		return float32(s)
	case uint64:
		return float32(s)
	case uint32:
		return float32(s)
	case uint16:
		return float32(s)
	case uint8:
		return float32(s)
	case string:
		v, err := strconv.ParseFloat(s, 64)
		return ReturnIf(err == nil, v, 0).(float32)
	case bool:
		return ReturnIf(s, 1, 0).(float32)
	default:
		return 0
	}
}

/*ToInt64 :  convierte un interface a int64*/
func ToInt64(i interface{}) int64 {
	i = AsignarPtr(i)
	switch s := i.(type) {
	case int:
		return int64(s)
	case int64:
		return s
	case int32:
		return int64(s)
	case int16:
		return int64(s)
	case int8:
		return int64(s)
	case uint:
		return int64(s)
	case uint64:
		return int64(s)
	case uint32:
		return int64(s)
	case uint16:
		return int64(s)
	case uint8:
		return int64(s)
	case float64:
		return int64(s)
	case float32:
		return int64(s)
	case string:
		v, err := strconv.ParseFloat(s, 64)
		return ReturnIf(err == nil, v, 0).(int64)
	case bool:
		return ReturnIf(s, 1, 0).(int64)
	case nil:
		return 0
	default:
		return 0
	}
}

/*ToInt32 :  convierte un interface a int32*/
func ToInt32(i interface{}) int32 {
	i = AsignarPtr(i)
	switch s := i.(type) {
	case int:
		return int32(s)
	case int64:
		return int32(s)
	case int32:
		return s
	case int16:
		return int32(s)
	case int8:
		return int32(s)
	case uint:
		return int32(s)
	case uint64:
		return int32(s)
	case uint32:
		return int32(s)
	case uint16:
		return int32(s)
	case uint8:
		return int32(s)
	case float64:
		return int32(s)
	case float32:
		return int32(s)
	case string:
		v, err := strconv.ParseFloat(s, 64)
		return ReturnIf(err == nil, v, 0).(int32)
	case bool:
		return ReturnIf(s, 1, 0).(int32)
	case nil:
		return 0
	default:
		return 0
	}
}

/*ToInt16 :  convierte un interface a int16*/
func ToInt16(i interface{}) int16 {
	i = AsignarPtr(i)

	switch s := i.(type) {
	case int:
		return int16(s)
	case int64:
		return int16(s)
	case int32:
		return int16(s)
	case int16:
		return s
	case int8:
		return int16(s)
	case uint:
		return int16(s)
	case uint64:
		return int16(s)
	case uint32:
		return int16(s)
	case uint16:
		return int16(s)
	case uint8:
		return int16(s)
	case float64:
		return int16(s)
	case float32:
		return int16(s)
	case string:
		v, err := strconv.ParseFloat(s, 64)
		return ReturnIf(err == nil, v, 0).(int16)
	case bool:
		return ReturnIf(s, 1, 0).(int16)
	case nil:
		return 0
	default:
		return 0
	}
}

/*ToInt8 :  convierte un interface a int8*/
func ToInt8(i interface{}) int8 {
	i = AsignarPtr(i)

	switch s := i.(type) {
	case int:
		return int8(s)
	case int64:
		return int8(s)
	case int32:
		return int8(s)
	case int16:
		return int8(s)
	case int8:
		return s
	case uint:
		return int8(s)
	case uint64:
		return int8(s)
	case uint32:
		return int8(s)
	case uint16:
		return int8(s)
	case uint8:
		return int8(s)
	case float64:
		return int8(s)
	case float32:
		return int8(s)
	case string:
		v, err := strconv.ParseFloat(s, 64)
		return ReturnIf(err == nil, v, 0).(int8)
	case bool:
		return ReturnIf(s, 1, 0).(int8)
	case nil:
		return 0
	default:
		return 0
	}
}

/*ToInt :  convierte un interface a int*/
func ToInt(i interface{}) int {
	i = AsignarPtr(i)

	switch s := i.(type) {
	case int:
		return s
	case int64:
		return int(s)
	case int32:
		return int(s)
	case int16:
		return int(s)
	case int8:
		return int(s)
	case uint:
		return int(s)
	case uint64:
		return int(s)
	case uint32:
		return int(s)
	case uint16:
		return int(s)
	case uint8:
		return int(s)
	case float64:
		return int(s)
	case float32:
		return int(s)
	case string:
		v := ToFloat64(s)
		return int(v)
	case bool:
		return ReturnIf(s, 1, 0).(int)
	case nil:
		return 0
	default:
		return 0
	}
}
func asignarPtrString(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	var errorType = reflect.TypeOf((*error)(nil)).Elem()
	var fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	v := reflect.ValueOf(a)
	for !v.Type().Implements(fmtStringerType) && !v.Type().Implements(errorType) && v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

/*ToString : convierte una data interface a un string cualquiera*/
func ToString(i interface{}) string {
	i = asignarPtrString(i)
	switch s := i.(type) {
	case string:
		return s
	case bool:
		return strconv.FormatBool(s)
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32)
	case int:
		return strconv.Itoa(s)
	case int64:
		return strconv.FormatInt(s, 10)
	case int32:
		return strconv.Itoa(int(s))
	case int16:
		return strconv.FormatInt(int64(s), 10)
	case int8:
		return strconv.FormatInt(int64(s), 10)
	case uint:
		return strconv.FormatUint(uint64(s), 10)
	case uint64:
		return strconv.FormatUint(uint64(s), 10)
	case uint32:
		return strconv.FormatUint(uint64(s), 10)
	case uint16:
		return strconv.FormatUint(uint64(s), 10)
	case uint8:
		return strconv.FormatUint(uint64(s), 10)
	case []byte:
		return string(s)
	case template.HTML:
		return string(s)
	case template.URL:
		return string(s)
	case template.JS:
		return string(s)
	case template.CSS:
		return string(s)
	case template.HTMLAttr:
		return string(s)
	case nil:
		return ""
	case fmt.Stringer:
		return s.String()
	case error:
		return s.Error()
	default:
		return ""
	}
}

/*EncodeBase64 :  convierte a base64 un string */
func EncodeBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

/*DecodeBase64 : convierte un string base64 a string normal*/
func DecodeBase64(s string) string {
	sDec, _ := base64.StdEncoding.DecodeString(s)
	return string(sDec)
}
