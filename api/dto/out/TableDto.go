package out

import (
	"sheets-database/domain/tables"
)

type TableDto struct {
	TableName string
	Rows      []RowDto
}

func TableDtoFromDomain(domain tables.Table) TableDto {
	dtoRows := []RowDto{}
	for _, domainRow := range domain.GetRows()   {
		dtoRows = append(dtoRows, RowDtoFromDomain(domainRow))
	}
	return TableDto{domain.GetTableName(), dtoRows}
}

type RowDto struct {
	Id     string
	Values []string
}

func RowDtoFromDomain(domain tables.Row) RowDto {
	return RowDto{domain.GetId(), domain.GetValues()}
}

type FullTableDto struct {
	TableName string
	Rows      []map[string]string
}

func FullTableDtoFromDomain(domain tables.FullTable) FullTableDto {
	return FullTableDto{
		domain.GetTableName(),
		domain.GetRows(),
	}
}
