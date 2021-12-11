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
			result = appendEntity(result, []map[string]interface{}{masterData})
		}()
	}
	return result, nil
}
func appendEntity(slice []map[string]interface{}, data []map[string]interface{}) []map[string]interface{} {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]map[string]interface{}, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
