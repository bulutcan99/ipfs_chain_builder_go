package model

type DbInformation struct {
	ColumnName string `db:"COLUMN_NAME"`
	DataType   string `db:"DATA_TYPE"`
}
