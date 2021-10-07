package utility

import (
	"fmt"
	"strings"
)

type (
	/*Component : crea componentes literal*/
	Component string
)

/*String : convierte en string*/
func (p *Component) String() string {
	vl := *p
	return string(vl)
}

/*Byte : convierte a byte el template*/
func (p *Component) Byte() []byte {
	vl := p.String()
	return StrtoByte(vl)
}

/*Print : reemplaza un key por data*/
func (p *Component) Print(key string, data interface{}) *Component {
	vl := p.String()
	dataStr := ToString(data)
	keyfinal := fmt.Sprintf("{%s}", key)
	vl = strings.Replace(vl, keyfinal, dataStr, -1)
	result := Component(vl)
	return &result
}

/*HTML : genera el  compnente con una base de html*/
func (p *Component) HTML(styles string) *Component {
	vl := p.String()
	tmp := Component(`
	<!DOCTYPE html>
	<html>
	<head>
		<style>
		{styles}
		</style>
	</head>
	<body>
		{body}
	</body>
	</html>
	`)
	return tmp.Print("styles", styles).Print("body", vl)
}

/*Body : carga el cuerto del literal component*/
func (p *Component) Body(vladd string) *Component {
	vl := p.String()
	vl += vladd
	result := Component(vl)
	return &result
}

/*PrintMap : genera un string literal html con datos mapeados*/
func (p *Component) PrintMap(data map[string]interface{}) *Component {
	vl := p.String()
	vl = PrintMap(vl, data)
	result := Component(vl)
	return &result
}

/*Add : agrega componentes a los template*/
func (p *Component) Add(data ...Component) *Component {
	tmp := ""
	for _, item := range data {
		tmp += item.String()
	}
	return p.Body(tmp)
}

/*Map : hace un ciclo de datos en una variable especifica {var}*/
func (p *Component) Map(key string, data ...Component) *Component {
	temp := Component("")
	vl := temp.Add(data...).String()
	return p.Print(key, vl)
}
