package database

import (
	"fmt"
)

type (
	/*StExt : extractores para la extructura merge y datatable*/
	StExt struct {
		SQLIn        StQuery
		TableNameOut string
		Index        []string
	}
	/*StMerge :  estructura para crear merge para servicio*/
	StMerge struct {
		CnxIn     StConect
		CnxOut    StConect
		ItemsExt  []StExt
		InDelIn   bool
		InDelOut  bool
		DelsqlIn  []StQuery
		DelsqlOut []StQuery
		//AccMerge : procesa la accion para hacer el merge de la base de datos
		AccMerge func(CnxIn, CnxOut StConect) error
	}
)

/*LoadDataIn : carga los datos de la base de datos de entrada*/
func (p *StMerge) LoadDataIn() ([]DataTable, error) {
	cnx := p.CnxIn
	defer cnx.Close()
	var (
		err    error
		result []DataTable
	)
	if p.InDelIn {
		err = cnx.Exec(p.DelsqlIn, true)
		if err != nil {
			return result, err
		}
	}
	if len(p.ItemsExt) <= 0 {
		return result, fmt.Errorf("los extractores estan vacio")
	}
	for _, v := range p.ItemsExt {
		data, err := cnx.QueryMap(v.SQLIn, 0, true, false)
		if err != nil {
			return result, err
		}
		datatable := NewDataTable(v.TableNameOut,
			data, v.Index)
		result = append(result, datatable)
	}

	return result, err
}

/*Process :  proceso de merge en los etl*/
func (p *StMerge) Process() error {
	data, err := p.LoadDataIn()
	if err != nil {
		return err
	}
	cnx := p.CnxOut
	if p.InDelOut {
		err = cnx.Exec(p.DelsqlOut, true)
		if err != nil {
			return err
		}
	}
	for _, v := range data {
		err = cnx.ExecDatatable(v, INSERT, true)
		if err != nil {
			return err
		}
	}
	cnx.Close()
	err = p.AccMerge(p.CnxIn, p.CnxOut)
	if err != nil {
		return err
	}

	return nil
}
