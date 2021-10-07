package utility

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type (
	/*ConfigWorker : configuracion del worker maestro para integracion de demonio */
	ConfigWorker struct {
		Workers   FormatWorker `json:"workers"`
		ConfigDir StArchMa     `json:"files"`
		Pathlog   string       `json:"pathlog"`
		Debug     bool         `json:"debug"`
	}
	/*FormatWorker : contiene los formatos de los worker para poder obtenerlos de forma de un json*/
	FormatWorker []struct {
		Tp  string `json:"type"`
		Key string `json:"key"`
		Fr  string `json:"format"`
	}
	/*MasterWorker : contiene varios worker y se administran de forma individual con maps*/
	MasterWorker struct {
		Jobs     map[string]Job
		pathjson string
		config   ConfigWorker
		logs     map[string]*StLog
		workers  map[string]*Worker
		print    map[string]bool
	}
)

/*SendWork : Reconfigura un worker el tiempo de ejecucion*/
func (p *MasterWorker) SendWork(key string) (*Worker, error) {
	err := p.ReloadWork(key)
	if err != nil {
		return nil, err
	}
	return p.workers[key], nil
}

/*Loadmaster : carga todas las configuraciones del master workers*/
func (p *MasterWorker) Loadmaster() error {
	err := p.loadworks()
	if err != nil {
		return err
	}
	err = p.loaddir()
	if err != nil {
		return err
	}
	err = p.loadlog()
	if err != nil {
		return err
	}
	return nil
}

/*NewMasterWorker :  Crea un MasterWorker*/
func NewMasterWorker(Jobs map[string]Job, config string, indload bool) (MasterWorker, error) {
	master := MasterWorker{
		pathjson: config,
		Jobs:     Jobs,
	}
	err := master.LoadConfig(config)
	if err != nil {
		return master, err
	}
	if indload {
		err := master.Loadmaster()
		if err != nil {
			return master, err
		}
	}

	return master, nil
}

/*ReloadWork : recarga un workers para recargar todos sus parametros*/
func (p *MasterWorker) ReloadWork(key string) error {
	var TaskMap Worker
	var err error
	indresp := false
	if !p.ValidWork(key) {
		return fmt.Errorf("el work no existe")
	}
	err = p.ReloadConfig()
	if err != nil {
		return err
	}
	Block{
		Try: func() {
			for _, v := range p.config.Workers {
				if v.Key == key {
					work, err := NewWork(v.Fr, v.Tp, p.Jobs[v.Key])
					if err != nil {
						break
					}
					TaskMap = work
					indresp = true
					break
				}
			}
		},
		Catch: func(e Exception) {
			err = fmt.Errorf("error en procesar el workers")
		},
	}.Do()
	if !indresp {
		return fmt.Errorf("el work no existe")
	}
	p.workers[key] = &TaskMap
	return err
}

/*Finally : Finalizacion del proceso donde el indicativo reset es para reintentar secuencia*/
func (p *MasterWorker) Finally(key string) {
	if p.ValidWork(key) {
		p.workers[key].Finally()
	}
}

/*FinallyDet : Finalizacion del proceso donde el indicativo reset es para reintentar secuencia ademas guarda en los logs si termino bien o tuvo un error ademas tiene la opcion de
reconfigurar el proceso para la proxima ejecucion*/
func (p *MasterWorker) FinallyDet(key, msg string, indreload, inderr bool) error {
	if p.ValidWork(key) {
		format := fmt.Sprintf("[%s] %s\n", ToDateStr(time.Now()), msg)
		p.workers[key].Finally()
		p.print[key] = false
		if inderr {
			PrintRed(format)
			p.Error(key, msg)
		} else {
			PrintGreen(format)
			p.Debug(key, msg)
		}
		if indreload && !inderr {
			return p.ReloadWork(key)
		}
		return nil
	}
	return nil
}

/*StartGen : ejecuta una tarea en paralelo en forma general con un log de inicio de proceso*/
func (p *MasterWorker) StartDet(key, msg string) {
	if p.ValidWork(key) {
		format := fmt.Sprintf("[%s] %s\n", ToDateStr(time.Now()), msg)
		p.workers[key].StartGen()
		if p.workers[key].GetStart() && !p.print[key] {
			p.print[key] = true
			p.Debug(key, msg)
			PrintGreen(format)
		}
	}
}

/*StartGen : ejecuta una tarea en paralelo en forma general*/
func (p *MasterWorker) StartGen(key string) {
	if p.ValidWork(key) {
		p.workers[key].StartGen()
	}
}

/*SetActivo : modifica para que se desactive la tarea*/
func (p *MasterWorker) SetActivo(key string, act bool) {
	if p.ValidWork(key) {
		p.workers[key].SetActivo(act)
	}
}

/*ValidWork :  valida si existe un worker*/
func (p *MasterWorker) ValidWork(key string) bool {
	_, ok := p.workers[key]
	return ok
}

/*Tick : regresa el secuencial de tiempo*/
func (p *MasterWorker) Tick(key string) <-chan time.Time {
	if p.ValidWork(key) {
		return p.workers[key].Tick()
	}
	return nil
}

/*Valid : envia si el proceso termino*/
func (p *MasterWorker) Valid(key string) <-chan bool {
	if p.ValidWork(key) {
		return p.workers[key].Valid()
	}
	return nil
}

/*Err :  envia el error en un canal de ejecucion*/
func (p *MasterWorker) Err(key string) <-chan error {
	if p.ValidWork(key) {
		return p.workers[key].Err()
	}
	return nil
}

/*LoadWorkers : carga todos los worker y lo envie en un map*/
func (p *MasterWorker) LoadWorkers() (map[string]*Worker, error) {
	err := p.loadworks()
	return p.workers, err
}

/*loadworks : carga todos los works con los archivos de configuracion*/
func (p *MasterWorker) loadworks() error {
	TaskMap := make(map[string]*Worker)
	Prints := make(map[string]bool)
	var err error
	Block{
		Try: func() {
			for _, v := range p.config.Workers {
				work, err := NewWork(v.Fr, v.Tp, p.Jobs[v.Key])
				if err != nil {
					break
				}
				TaskMap[v.Key] = &work
				Prints[v.Key] = false
			}
		},
		Catch: func(e Exception) {
			err = fmt.Errorf("error en procesar el map workers")
		},
	}.Do()
	if err == nil {
		p.workers = TaskMap
		p.print = Prints
	}
	return err
}

/*LoadConfig : carga el archivo de configuracion de los worker como archivo json*/
func (p *MasterWorker) LoadConfig(PathJSON string) error {
	var (
		err     error
		config  ConfigWorker
		ptrArch *os.File
	)
	if !FileExt(PathJSON, "JSON") {
		return Msj.GetError("CN09")
	}
	p.pathjson = PathJSON
	PathJSON, err = TrimFile(PathJSON)
	if err != nil {
		return Msj.GetError("CN08")
	}
	ptrArch, err = os.Open(PathJSON)
	if err != nil {
		return Msj.GetError("CN08")
	}
	defer ptrArch.Close()
	decJSON := json.NewDecoder(ptrArch)
	err = decJSON.Decode(&config)
	if err != nil {
		return Msj.GetError("CN08")
	}
	p.config = config
	return nil
}

/*ReloadConfig : recarga el archivo de configuracion por cualquier cambio de los workers*/
func (p *MasterWorker) ReloadConfig() error {
	return p.LoadConfig(p.pathjson)
}

/*loadlog : carga los diferentes logs de los distintos workers para tener un proceso limpio*/
func (p *MasterWorker) loadlog() error {
	if p.config.Pathlog != "" {
		fileslog := StArchMa{
			{
				Path:   p.config.Pathlog + "/logs",
				IndDir: true,
			}, {
				Path:   p.config.Pathlog + "/logs/error",
				IndDir: true,
			}, {
				Path:   p.config.Pathlog + "/logs/debug",
				IndDir: true,
			},
		}
		err := fileslog.Create()
		if err != nil {
			return err
		}
		LogsMap := make(map[string]*StLog)
		for key := range p.workers {
			if p.config.Debug {
				LogsMap[key+"d"] = &StLog{
					Dir:    p.config.Pathlog + "/logs/debug",
					Prefix: "DEBUG",
					Name:   key,
					Fe:     time.Now(),
				}
				err = LogsMap[key+"d"].Init()
				if err != nil {
					return err
				}
			}

			LogsMap[key+"e"] = &StLog{
				Dir:    p.config.Pathlog + "/logs/error",
				Prefix: "ERROR",
				Name:   key,
				Fe:     time.Now(),
			}
			err = LogsMap[key+"e"].Init()
			if err != nil {
				return err
			}
		}
		p.logs = LogsMap
	}
	return nil
}

/*loaddir : carga los  directorios y archivos de configuracion del proyecto*/
func (p *MasterWorker) loaddir() error {
	if len(p.config.ConfigDir) > 0 {
		err := p.config.ConfigDir.Create()
		if err != nil {
			return err
		}
	}
	return nil
}

/*Printf : Ingresa un texto en los logs asignados. */
func (p *MasterWorker) Printf(err bool, key, format string, args ...interface{}) error {
	if p.ValidWork(key) {
		tp := ReturnIf(err, key+"e", key+"d").(string)
		p.logs[tp].Printf(format, args...)
	}
	return nil
}

/*Debug : ingresa un texto en modo debug*/
func (p *MasterWorker) Debug(key, format string, args ...interface{}) error {
	return p.Printf(false, key, format, args...)
}

/*Debug : ingresa un texto en modo error*/
func (p *MasterWorker) Error(key, format string, args ...interface{}) error {
	return p.Printf(true, key, format, args...)
}
