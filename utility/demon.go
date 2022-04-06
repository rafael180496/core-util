package utility

import (
	"fmt"
	"time"
)

type (
	/*Worker : Orquestador de procesos paralelo*/
	Worker struct {
		start     bool
		finalid   bool
		fexec     time.Time
		ticker    *time.Ticker
		hr        HrTime
		job       Job
		valid     chan bool
		err       chan error
		activo    bool
		indticker bool
	}
	/*Job : alias para un proceso*/
	Job func(chan bool, chan error)
)

/*Finally : Finalizacion del proceso donde el indicativo reset es para reintentar secuencia*/
func (p *Worker) Finally() {
	p.start = false
}

/*GetStart : Envia si el proceso esta en ejecucion*/
func (p *Worker) GetStart() bool {
	return p.start
}

/*SetActivo : modifica para que se desactive la tarea*/
func (p *Worker) SetActivo(act bool) {
	p.activo = act
}

/*StartGen : ejecuta una tarea en paralelo en forma general*/
func (p *Worker) StartGen() {
	if p.indticker {
		p.Start()
	} else {
		p.StartTime()
	}
}

/*Start : Ejecuta el proceso paralelo*/
func (p *Worker) Start() {
	if !p.start && p.activo {
		p.start = true
		p.valid = make(chan bool)
		p.err = make(chan error)
		go p.job(p.valid, p.err)
	}
}

/*StartTime : ejecuta un proceso paralelo mediante una hora especifica*/
func (p *Worker) StartTime() {
	if !DateEquals(p.fexec, time.Now()) && p.finalid {
		p.finalid = false
	}
	if p.hr.EqualNow() && !p.finalid {
		p.fexec = time.Now()
		p.finalid = true
		p.Start()
	}
}

/*Tick : regresa el secuencial de tiempo*/
func (p *Worker) Tick() <-chan time.Time {
	return p.ticker.C
}

/*Valid : envia si el proceso termino*/
func (p *Worker) Valid() <-chan bool {
	return p.valid
}

/*Err :  envia el error en un canal de ejecucion*/
func (p *Worker) Err() <-chan error {
	return p.err
}

/*NewWork : crea una tarea con formatos y tipo de de secuencia
1 - Ticker
2 - Time
Formatos:
05 hr
54 min
12:16 time
*/
func NewWork(format, tp string, job Job) (Worker, error) {
	switch tp {
	case "hr", "min":
		tiker, err := NewTicker(format, tp)
		if err != nil {
			return Worker{}, err
		}
		wrol := NewWorkTicker(tiker, job)
		return wrol, nil
	case "time":
		hrtime, err := NewHrTimeStr(format)
		if err != nil {
			return Worker{}, err
		}
		wrol, err := NewWorkTime(hrtime.hr, hrtime.min, job)
		if err != nil {
			return Worker{}, err
		}
		return wrol, nil
	default:
		return Worker{}, fmt.Errorf("error en tipo")
	}
}

/*NewWorkTicker : crea una tarea mediante un ticker para ponerlo en un servicio*/
func NewWorkTicker(p *time.Ticker, job Job) Worker {
	return Worker{
		finalid:   false,
		job:       job,
		ticker:    p,
		start:     false,
		activo:    true,
		indticker: true,
	}
}

/*NewWorkTime : crea una tarea mediante una hora exacta para ponerlo en un servicio*/
func NewWorkTime(hr, min int, job Job) (Worker, error) {
	hrTime, err := NewHrTime(hr, min)
	if err != nil {
		return Worker{}, err
	}
	return Worker{
		finalid:   false,
		job:       job,
		hr:        hrTime,
		start:     false,
		activo:    true,
		indticker: false,
	}, nil
}
