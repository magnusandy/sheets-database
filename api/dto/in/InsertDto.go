package in

import "sheets-database/domain/tables"

type InsertDto struct {
	SheetId string
	TableName string
	Data []map[string]string//todo maybe easier as interface?
	//todo force insert (one fail shouldnt make the rest not go) default one fail all stop
}

func (i InsertDto) ToDomain() tables.FullTable {
	return tables.CreateFullTable(i.TableName, i.Data)
}