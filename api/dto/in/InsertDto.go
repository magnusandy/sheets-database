package in

import (
	"sheets-database/domain/tables"
	"sheets-database/domain/metadata"
	"github.com/google/uuid"
)

type InsertDto struct {
	SheetId string
	TableName string
	Data []map[string]string//todo maybe easier as interface?
	//todo force insert (one fail shouldnt make the rest not go) default one fail all stop
}

func (i InsertDto) ToDomain() tables.FullTable {
	withIds := i.attachIds()
	return tables.CreateFullTable(withIds.TableName, withIds.Data)
}

//todo should this be in the domain?
func (i InsertDto) attachIds() InsertDto {
	for _, rowMap := range i.Data {
		rowMap[metadata.ID_COLUMN] = uuid.New().String()
	}
	return i
}