package repository

import (
	"context"
	"fmt"
	"github.com/bulutcan99/go_ipfs_chain_builder/internal/aggregate"
	"github.com/bulutcan99/go_ipfs_chain_builder/model"
	config_mysql "github.com/bulutcan99/go_ipfs_chain_builder/pkg/config/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type IAggregateRepo interface {
	GetUsersWithColumnTypes() ([]aggregate.AggregatedData, error)
}

type AggregateRepo struct {
	db      *sqlx.DB
	context context.Context
}

func NewAggregateRepo(db *config_mysql.MYSQL) *AggregateRepo {
	return &AggregateRepo{
		db:      db.Client,
		context: db.Context,
	}
}

func (r *AggregateRepo) GetUsersWithColumnTypes() ([]aggregate.AggregatedData, error) {
	var aggregatedColumnInfo []aggregate.AggregatedData

	const query = `
		SELECT COLUMN_NAME, DATA_TYPE 
		FROM information_schema.columns 
		WHERE TABLE_SCHEMA = 'ipfs' AND TABLE_NAME = 'users';`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query for column information")
	}
	defer rows.Close()

	for rows.Next() {
		var columnName, dataType string

		if err := rows.Scan(&columnName, &dataType); err != nil {
			return nil, errors.Wrap(err, "failed to scan column information")
		}
		fmt.Println("columnName:", columnName, "dataType:", dataType)
		columnQuery := fmt.Sprintf("SELECT %s FROM users", columnName)
		userRows, err := r.db.Query(columnQuery)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to execute query for column %s", columnName)
		}
		defer userRows.Close()

		for userRows.Next() {
			var value interface{}

			switch dataType {
			case "int":
				var intValue int
				if err := userRows.Scan(&intValue); err != nil {
					return nil, errors.Wrap(err, "failed to scan user row value")
				}
				value = intValue
			case "varchar":
				var textValue string
				if err := userRows.Scan(&textValue); err != nil {
					return nil, errors.Wrap(err, "failed to scan user row value")
				}
				value = textValue
			default:
				return nil, errors.Errorf("unsupported data type: %s", dataType)
			}

			aggregateInfo := aggregate.AggregatedData{
				ColumnInfo: model.DbInformation{
					ColumnName: columnName,
					DataType:   dataType,
				},
				UserValue: value,
			}
			fmt.Println("aggregateInfo:", aggregateInfo)
			aggregatedColumnInfo = append(aggregatedColumnInfo, aggregateInfo)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "error during iteration")
	}

	return aggregatedColumnInfo, nil
}
