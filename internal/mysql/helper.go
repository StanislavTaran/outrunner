package mysql

import (
	"database/sql"
	"encoding/hex"
	"strconv"
)

func rowsToJSON(rows *sql.Rows) ([]map[string]interface{}, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, 0)

	if err != nil {
		return nil, err
	}

	count := len(columns)
	values := make([]interface{}, count)
	scanArgs := make([]interface{}, count)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		func() {
			masterData := make(map[string]interface{})
			for i, v := range values {

				// TODO need improvements
				switch colTypes[i].DatabaseTypeName() {
				case "CHAR", "VARCHAR", "TEXT", "ENUM", "SET":
					x := v.([]byte)
					masterData[columns[i]] = string(x)
				case "INTEGER", "INT", "TINYINT", "SMALLINT", "FLOAT", "DOUBLE", "DECIMAL", "NUMERIC":
					x := v.([]byte)
					masterData[columns[i]], _ = strconv.ParseFloat(string(x), 64)
				case "DATE", "TIME", "DATETIME", "TIMESTAMP", "YEAR":
					masterData[columns[i]] = v
				case "BINARY", "VARBINARY", "BLOB":
					x := v.([]byte)
					res := hex.EncodeToString(x)
					masterData[columns[i]] = res
				default:
					masterData[columns[i]] = v
				}
			}
			result = append(result, masterData)
		}()
	}
	return result, nil
}
