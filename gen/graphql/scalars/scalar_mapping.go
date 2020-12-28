package scalars

import (
	"gqlgen-generator/gen/table"
)

var ScalarMapping map[string]string

const (
	SCALAR_BIGINT      = "Bigint"
	SCALAR_INT         = "Int"
	SCALAR_STRING      = "String"
	SCALAR_FLOAT       = "Float"
	SCALAR_NUMERIC     = "Numeric"
	SCALAR_TIMESTAMPTZ = "Timestamptz"
	SCALAR_POINT       = "Point"
	SCALAR_BOOLEAN     = "Boolean"
)

// 默认的gql类型
var DefaultGqlTypeNamed = SCALAR_STRING

func init() {
	ScalarMapping = make(map[string]string)
	ScalarMapping["int8"] = SCALAR_BIGINT
	ScalarMapping["int4"] = SCALAR_INT
	ScalarMapping["text"] = SCALAR_STRING
	ScalarMapping["timestamptz"] = SCALAR_TIMESTAMPTZ
	ScalarMapping["point"] = SCALAR_POINT
	ScalarMapping["float4"] = SCALAR_FLOAT
	ScalarMapping["bool"] = SCALAR_BOOLEAN
	ScalarMapping["numeric"] = SCALAR_NUMERIC
}

// 根据数据库字段类型获取gql字段类型
func CalculateNamedType(columnInfo table.ColumnInfo) string {

	gqlFieldTypeNamed, ok := ScalarMapping[columnInfo.Columntype]
	if !ok {
		return DefaultGqlTypeNamed
	}
	return gqlFieldTypeNamed
}

// 获取主键的gql字段类型
func GetPrimaryKeyType(table table.TableInfo) string {
	for _, column := range table.Columns {
		if column.Primary {
			return CalculateNamedType(column)
		}
	}
	return ""
}
