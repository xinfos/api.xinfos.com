package model

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

type Row []interface{}

//BatchInsertRawSQL 批量创建
func BatchInsertRawSQL(db *gorm.DB, table string, columns []string, rows []Row) error {
	if len(rows) == 0 {
		return nil
	}

	sql := fmt.Sprintf("INSERT INTO %s(%s) VALUES ", table, strings.Join(columns, ","))
	length := len(columns)

	var valueHolders []string
	for i := 0; i < length; i++ {
		valueHolders = append(valueHolders, "?")
	}
	placeHolder := fmt.Sprintf("(%s)", strings.Join(valueHolders, ","))

	var values []string
	var args []interface{}
	for _, row := range rows {
		if len(row) != length {
			return fmt.Errorf("Row (%v) does not match columns (%v)", row, columns)
		}
		values = append(values, placeHolder)
		args = append(args, row...)
	}
	sql += strings.Join(values, ",")
	return db.Exec(sql, args...).Error
}

//QueryArrayToString 将查询条件数组转换成查询字符串
func QueryArrayToString(query []string) string {
	queryStr := ""
	if len(query) > 0 {
		queryStr = strings.Join(query, " AND ")
	}
	return queryStr
}
