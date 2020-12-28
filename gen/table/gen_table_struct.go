package table

import (
	"errors"
	"gqlgen-generator/db"
	"regexp"
)

var CleanRegexp, _ = regexp.Compile(`\(\d+\)`)

type TableInfo struct {
	// 表名
	TableName string
	// 表描述
	TableDesc string
	// 列信息
	Columns []ColumnInfo
}

type ColumnInfo struct {
	// 列名
	Columnname string
	// 列描述
	Columndesc string
	// 是否可以为空
	Nullable bool
	// 是否是主键
	Primary bool
	// 列类型
	Columntype string
}

func (t *TableInfo) PrimaryKeyColumn() *ColumnInfo {
	for i, _ := range t.Columns {
		if t.Columns[i].Primary {
			return &t.Columns[i]
		}
	}
	return nil
}

//解析表信息
func ResolverTableInfos(tables []string) (tableInfos []TableInfo, err error) {
	for _, table := range tables {
		exist := isTableExist(table)
		if !exist {
			err = errors.New("table:" + table + " not exist")
			return
		}
		tableInfo := TableInfo{
			TableName: table,
		}
		tableDesc := getTableDesc(table)
		tableInfo.TableDesc = tableDesc
		tableInfo.Columns = getTableColumns(table)
		tableInfos = append(tableInfos, tableInfo)
	}
	return
}

// 判断表是否存在
func isTableExist(table string) bool {
	baseSql := `
SELECT COUNT
	( * ) 
FROM
	information_schema.tables 
WHERE
	table_schema = CURRENT_SCHEMA ( ) 
	AND TABLE_NAME = ? 
	AND table_type = 'BASE TABLE'
`
	var count int
	db.DB.Raw(baseSql, table).Scan(&count)
	if count > 0 {
		return true
	}
	return false
}

//获取表描述信息
func getTableDesc(table string) string {
	baseSql := `
SELECT CAST
	( obj_description ( relfilenode, 'pg_class' ) AS VARCHAR ) AS COMMENT 
FROM
	pg_class C 
WHERE
	relname = ?`
	var tableDesc string
	db.DB.Raw(baseSql, table).Scan(&tableDesc)
	return tableDesc
}

// 获取表的主键名称
func getTablePrimaryKeyName(table string) string {
	baseSql := `
SELECT
	pg_attribute.attname AS colname 
FROM
	pg_constraint
	INNER JOIN pg_class ON pg_constraint.conrelid = pg_class.oid
	INNER JOIN pg_attribute ON pg_attribute.attrelid = pg_class.oid 
	AND pg_attribute.attnum = pg_constraint.conkey [ 1 ]
	INNER JOIN pg_type ON pg_type.oid = pg_attribute.atttypid 
WHERE
	pg_class.relname = ? 
	AND pg_constraint.contype = 'p'
`
	var primaryColumnName string
	db.DB.Raw(baseSql, table).Scan(&primaryColumnName)
	return primaryColumnName
}

// 获取列信息
func getTableColumns(table string) (columnInfos []ColumnInfo) {
	baseSql := `
SELECT
	A.attname AS columnName,
	concat_ws ( '', T.typname, SUBSTRING ( format_type ( A.atttypid, A.atttypmod ) FROM '\(.*\)' ) ) AS columnType,
	d.description AS columndesc ,
	(case  A.attnotnull when 't' then false else true end) as nullable
FROM
	pg_class C,
	pg_attribute A,
	pg_type T,
	pg_description d 
WHERE
	C.relname =? 
	AND A.attnum > 0 
	AND A.attrelid = C.oid 
	AND A.atttypid = T.oid 
	AND d.objoid = A.attrelid 
	AND d.objsubid = A.attnum
`
	db.DB.Raw(baseSql, table).Scan(&columnInfos)
	primaryKeyColumn := getTablePrimaryKeyName(table)
	for i, _ := range columnInfos {
		if columnInfos[i].Columnname == primaryKeyColumn {
			columnInfos[i].Primary = true
		}
		columnInfos[i].Columntype = cleanColumnType(columnInfos[i].Columntype)

	}
	return
}

func cleanColumnType(columnType string) string {
	s := CleanRegexp.Split(columnType, 2)
	return s[0]

}
