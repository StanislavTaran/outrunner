package mysql

import (
	"database/sql"
	"fmt"
	"strconv"
)

func rowsToJSON(rows *sql.Rows) ([]map[string]interface{}, error) {
	columns, err := rows.Columns()
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
				x := v.([]byte)
				if nx, ok := strconv.ParseFloat(string(x), 64); ok == nil {
					masterData[columns[i]] = nx
				} else if b, ok := strconv.ParseBool(string(x)); ok == nil {
					masterData[columns[i]] = b
				} else if "string" == fmt.Sprintf("%T", string(x)) {
					masterData[columns[i]] = string(x)
				} else {
					fmt.Printf("Failed on if for type %T of %v\n", x, x)
				}
			}
			result = append(result, masterData)
		}()
	}
	return result, nil
}
