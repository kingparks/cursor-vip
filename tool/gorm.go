package tool

import (
	"database/sql"
	"gorm.io/gorm"
)

// RawMap sql查询返回[]map[string]string类型
func RawMap(db *gorm.DB, sqlQuery string, sqlValues ...interface{}) (result []map[string]string, err error) {
	rows, err := db.Raw(sqlQuery, sqlValues...).Rows()
	if err != nil {
		return
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return
	}
	values := make([]sql.RawBytes, len(cols))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return
		}
		var value string
		resultC := map[string]string{}
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			resultC[cols[i]] = value
		}
		result = append(result, resultC)
	}
	return
}
