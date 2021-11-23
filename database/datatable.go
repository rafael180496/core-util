package database

import (
	"fmt"
	"strings"

	utl "github.com/rafael180496/core-util/utility"
)

type (
	/*DataTable : maneja un crud completo y genera script automaticos*/
	DataTable struct {
		table string
		rows  []StData
		index []string
	}
)

/*NewDataTable : Crea un datable correctamente*/
func NewDataTable(table string, rows []StData, index []string) DataTable {
	var data DataTable
	data.SetTable(table)
	data.AddIndexs(index...)
	data.AddRows(rows...)
	return data
}

/*ValidRow : valida si tiene filas llenadas*/
func (p *DataTable) ValidRow() bool {
	return utl.ReturnIf(len(p.rows) > 0, true, false).(bool)
}

/*GetCols : Obtiene el nombre de las columnas de cada fila*/
func (p *DataTable) GetCols() ([]string, error) {
	var cols []string
	item, err := p.GetRow(1)
	if err != nil {
		return cols, err
	}
	cols = item.KeyColum()
	if len(cols) <= 0 {
		return cols, fmt.Errorf("they have no loaded columns")
	}
	return cols, nil
}

/*GetIndex : Obtiene el nombre de los indices*/
func (p *DataTable) GetIndex() []string {
	return p.index
}

/*GetTable : Obtiene el nombre de la tabla*/
func (p *DataTable) GetTable() string {
	return p.table
}

/*SetTable : Modifica el nombre de la tabla*/
func (p *DataTable) SetTable(table string) {
	p.table = strings.ToUpper(utl.Trim(table))
}

/*LenRows : Obtienen la cantidad de fila*/
func (p *DataTable) LenRows() int {
	return len(p.rows)
}

/*LenIndex : Obtienen la cantidad de llaves primarias para update o delete*/
func (p *DataTable) LenIndex() int {
	return len(p.index)
}

/*AddIndexs : agrega varias claves primarias*/
func (p *DataTable) AddIndexs(cols ...string) error {
	for _, col := range cols {
		err := p.AddIndex(col)
		if err != nil {
			return err
		}
	}
	return nil
}

/*AddRows : Agrega varias fila */
func (p *DataTable) AddRows(rows ...StData) {
	for _, item := range rows {
		p.AddRow(item)
	}
}

/*AddRow : Agrega una fila */
func (p *DataTable) AddRow(row StData) {
	p.rows = append(p.rows, row.UpperKey())
}

/*AddIndex : agrega una llave para los delete o update*/
func (p *DataTable) AddIndex(col string) error {
	col = strings.ToUpper(col)
	item, err := p.GetRow(1)
	if err != nil {
		return err
	}
	if !item.ValidColum(col) {
		return fmt.Errorf("the column does not exist")
	}
	p.index = append(p.index, col)
	return nil
}

/*GetRow : Obtiene una fila X por el numero de fila*/
func (p *DataTable) GetRow(row int) (StData, error) {
	if p.LenRows() < row {
		return StData{}, fmt.Errorf("the row number exceeds the total number of rows")
	}
	return p.rows[(row - 1)], nil
}

/*GetRows : Obtiene todas las fila*/
func (p *DataTable) GetRows() []StData {
	return p.rows
}

/*GenSQL : genera acciones de base de datos mediante los siguientes comando INSERT,UPDATE,DELETE*/
func (p *DataTable) GenSQL(accion string) ([]StQuery, error) {
	switch accion {
	case INSERT:
		return p.GenInserts()
	case DELETE:
		return p.GenDeletes()
	case UPDATE:
		return p.GenUpdates()
	default:
		return nil, fmt.Errorf("invalid action datatable")
	}
}

/*GenInserts : genera insert masivos para modificaciones de base de datos*/
func (p *DataTable) GenInserts() ([]StQuery, error) {
	var queries []StQuery
	clone := *p
	sqltemp, err := sqldinamic(clone, INSERT)
	if err != nil {
		return queries, err
	}
	for _, quirie := range p.GetRows() {
		queries = append(queries, StQuery{
			Querie: sqltemp,
			Args:   quirie,
		})
	}
	return queries, nil
}

/*GenDeletes : genera los delete masivos para modificaciones de base de datos*/
func (p *DataTable) GenDeletes() ([]StQuery, error) {
	var queries []StQuery
	clone := *p
	sqltemp, err := sqldinamic(clone, DELETE)
	if err != nil {
		return queries, err
	}
	cols, _ := p.GetCols()
	colsnew := utl.FilterExcl(cols, clone.GetIndex())
	for _, quirie := range p.GetRows() {
		queries = append(queries, StQuery{
			Querie: sqltemp,
			Args:   quirie.Filter(colsnew...),
		})
	}
	return queries, nil
}

/*GenUpdates : genera los update  masivos para modificaciones de base de datos de ante mano tener que colocar indices*/
func (p *DataTable) GenUpdates() ([]StQuery, error) {
	var queries []StQuery
	clone := *p
	sqltemp, err := sqldinamic(clone, UPDATE)
	if err != nil {
		return queries, err
	}
	for _, quirie := range p.GetRows() {
		queries = append(queries, StQuery{
			Querie: sqltemp,
			Args:   quirie,
		})
	}
	return queries, nil
}

/*sqldinamic : genera los sql temporales para los crud*/
func sqldinamic(data DataTable, acc string) (string, error) {
	table := utl.Trim(data.GetTable())
	if table == "" {
		return "", fmt.Errorf("they do not have a loaded table")
	}
	cols, err := data.GetCols()
	if err != nil {
		return "", err
	}
	if !utl.ValidDuplidArrayStr(cols) {
		return "", fmt.Errorf("duplicate column")
	}
	if utl.InStr(acc, UPDATE, DELETE) && data.LenIndex() <= 0 {
		return "", fmt.Errorf("duplicate column")
	}
	switch acc {
	case INSERT:
		sqltmp := sqlinsert(table, cols)
		return sqltmp, nil
	case UPDATE:
		sqltmp := sqlupdate(table, cols, data.GetIndex())
		return sqltmp, nil
	case DELETE:
		sqltmp := sqldelete(table, data.GetIndex())
		return sqltmp, nil
	default:
		return "", nil

	}
}
func sqldelete(table string, indices []string) string {
	sqltmp := fmt.Sprintf("DELETE FROM  %s", table)
	sqltmp = sqlConditional(sqltmp, indices)
	return sqltmp
}

func sqlupdate(table string, cols []string, indices []string) string {
	sqltmp := fmt.Sprintf("UPDATE %s SET", table)
	values := utl.FilterExcl(cols, indices)
	for i, item := range values {
		ind := (len(values) - 1)
		sqltmp = fmt.Sprintf("%s %s = :%s%s", sqltmp, item, item, utl.ReturnIf(i == ind, "", " ,").(string))
	}
	sqltmp = sqlConditional(sqltmp, indices)
	return sqltmp
}

func sqlConditional(sqltmp string, indices []string) string {
	sqltmp = fmt.Sprintf("%s WHERE ", sqltmp)
	for i, item := range indices {
		ind := (len(indices) - 1)
		sqltmp = fmt.Sprintf("%s %s = :%s%s", sqltmp, item, item, utl.ReturnIf(i == ind, "", " AND").(string))
	}
	return sqltmp
}

func sqlinsert(table string, cols []string) string {
	sqltmp := fmt.Sprintf("INSERT INTO %s (", table)
	for i, item := range cols {
		if i == (len(cols) - 1) {
			sqltmp = fmt.Sprintf("%s%s) VALUES(", sqltmp, item)
		} else {
			sqltmp = fmt.Sprintf("%s%s,", sqltmp, item)
		}
	}
	for i, item := range cols {
		ind := (len(cols) - 1)
		sqltmp = fmt.Sprintf("%s:%s%s", sqltmp, item, utl.ReturnIf(i == ind, ")", ",").(string))
	}
	return sqltmp
}
