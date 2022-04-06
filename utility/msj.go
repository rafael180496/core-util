package utility

type (
	/*StMsj : estructura para manejo de errores y mensajes personalizados */
	StMsj struct {
		Store map[string]string
	}
)

/*AddMsj : agrega un mensaje con su codigo unico */
func (p *StMsj) AddMsj(code, msj string) {
	p.Store[code] = msj
}

/*AddStore : agrega su store completo */
func (p *StMsj) AddStore(store map[string]string) {
	p.Store = store
}

/*GetString : envia un mensaje por medio del codigo */
func (p *StMsj) GetString(code string) string {
	return p.Store[code]
}

/*GetError : envia un error medio el codigo */
func (p *StMsj) GetError(code string) error {
	return new(StError).MandError(p.Store[code])
}
