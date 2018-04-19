package domain

type Table struct {
	TableName string
	Rows []Row
}

type Row struct {
	Id string
	Values []string
}