package in

import "sheets-database/domain/metadata"

type SelectDto struct {
	SheetId string
	TableName string
	Format metadata.ResultFormat //nullable
	//todo limit
	//todo offset
	//todo mapOrList for values
}