package dto

import (
	"strconv"
	"sheets-database/domain"
)

type GetAllData struct {
	SpreadSheetId string                `json:"spreadsheetId"`
	Properties    SpreadSheetProperties `json:"properties"`
	Sheets        []Sheet               `json:"sheets"`
}

func (get GetAllData) ToDomain() []domain.Table {
	var tables []domain.Table;
	for i := 0; i < len(get.Sheets); i++ {
		sheet := get.Sheets[i]
		tables = append(tables, sheet.toDomain())
	}
	return tables
}

type SpreadSheetProperties struct {
	Title string `json:"title"`
}

type Sheet struct {
	Properties SheetProperties `json:"properties"`
	Data       []Data          `json:"data"`
}

func (s Sheet) toDomain() domain.Table {
	var rows []domain.Row
	var rowData []RowData = s.Data[0].RowData
	for i := 0; i < len(rowData); i++ {
		workingData := rowData[i]
		rows = append(rows,
			domain.Row{
				*workingData.Values[0].UserEnteredValue.GetValue(),
				workingData.toDomain(),
				})
	}
	return domain.Table{
		s.Properties.Title,
		rows,
	}
}

type SheetProperties struct {
	SheetId int    `json:"sheetId"`
	Title   string `json:"title"`
}

//seems to always have only one element in the list
type Data struct {
	RowData []RowData `json "rowData"`
}

type RowData struct {
	Values []RowValue `json:"values"`
}

func (d RowData) toDomain() []string {
	var dataValues []string
	for i:=1;i<len(d.Values) ;i++  {
		dataValues = append(dataValues, *d.Values[i].UserEnteredValue.GetValue())
	}
	return dataValues
}

type RowValue struct {
	Note             string           `json:"note"`
	UserEnteredValue UserEnteredValue `json:"userEnteredValue"`
}

//todo test
type UserEnteredValue struct {
	BoolValue   *bool    `json:boolValue, omitempty`
	NumberValue *float64 `json:numberValue, omitempty`
	StringValue *string  `json:stringValue, omitempty`
}

func (userEnteredValue UserEnteredValue) GetValue() *string {
	var enteredValueString string
	if userEnteredValue.StringValue != nil {
		enteredValueString = *userEnteredValue.StringValue
	} else if userEnteredValue.NumberValue != nil {
		enteredValueString = strconv.FormatFloat(*userEnteredValue.NumberValue, 'f', -1, 64)
	} else if userEnteredValue.BoolValue != nil{
		enteredValueString = strconv.FormatBool(*userEnteredValue.BoolValue)
	} else {
		enteredValueString = "NULL"//todo definable
	}
	return &enteredValueString
}
