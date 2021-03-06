package infra

import (
	"sheets-database/domain"
	"google.golang.org/api/sheets/v4"
	"sheets-database/domain/tables"
)

const MAJOR_DIMENTION_ROWS = "ROWS"
const USER_INPUT_OPTION = "USER_ENTERED"

type RestSheetsService struct {
	authService domain.AuthenticationService
}

func CreateRestSheetService(authService domain.AuthenticationService) domain.SheetsService {
	return RestSheetsService{authService}
}

func (r RestSheetsService) GetAllData(sheetId string) []tables.Table {
	_, err := createSheetsClient(r.authService)
	domain.LogIfPresent(err)
	return nil
}

func (r RestSheetsService) InsertRowsIntoTable(sheetId string, table tables.Table) error {
	sheetClient, err := createSheetsClient(r.authService)
	domain.LogWithMessageIfPresent("problem with sheet client connection", err)

	_, sheetErr := sheetClient.Spreadsheets.Values.
		Append(sheetId, table.GetTableName(), serializeValueRange(table)).
		ValueInputOption(USER_INPUT_OPTION).
		Do()

	return sheetErr
}

//todo error handling
func (r RestSheetsService) GetAllDataForTable(sheetId string, tableName string) (tables.Table, error) {
	sheetClient, err := createSheetsClient(r.authService)
	domain.LogWithMessageIfPresent("problem with sheet client connection", err)
	values, googleError := sheetClient.Spreadsheets.Values.Get(sheetId, tableName).MajorDimension(MAJOR_DIMENTION_ROWS).Do()
	domain.LogWithMessageIfPresent("google sheets error", googleError)
	return deserializeValueRangeToDomain(values, tableName), nil
}

func deserializeValueRangeToDomain(valueRange *sheets.ValueRange, tableName string) tables.Table {
	var tableRows []tables.Row
	for _, typelessRow := range valueRange.Values {
		typedRow := typeCastRow(typelessRow)
		id := typedRow[0]
		values := make([]string, len(typedRow)-1) //dont need the first value as its already in the id
		copy(values, typedRow[1:])
		domainRow := tables.CreateRow(id, values)
		tableRows = append(tableRows, domainRow)
	}
	return tables.CreateTable(tableName, tableRows)
}

func serializeValueRange(table tables.Table) *sheets.ValueRange {
	outValues := [][]interface{}{}
	for _, row := range table.GetRows() {
		outValues = append(outValues, serializeRow(row))
	}

	valueRange := sheets.ValueRange{
		MajorDimension: MAJOR_DIMENTION_ROWS,
		Range:          table.GetTableName(),
		Values:         outValues,
	}

	return &valueRange
}

func typeCastRow(row []interface{}) []string {
	newRow := []string{}
	for _, value := range row {
		newVal := value.(string)
		newRow = append(newRow, newVal)
	}
	return newRow
}

func serializeRow(row tables.Row) []interface{} {
	out := []interface{}{}
	out = append(out, row.GetId())
	for _, rowVal := range row.GetValues() {
		out = append(out, rowVal)
	}
	return out
}
