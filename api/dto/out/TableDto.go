package out

type ListTableDto struct {
	TableName string
	Rows []ListRowDto
}

type ListRowDto struct {
	Id string
	Values []string
}

type MapTableDto struct {
	TableName string
	Rows []map[string]string
}


