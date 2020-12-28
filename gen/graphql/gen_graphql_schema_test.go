package graphql

import (
	"fmt"
	"gqlgen-generator/gen/table"
	"testing"
)

func Test_generateGraphqlText(t *testing.T) {
	tableInfo := &table.TableInfo{
		TableName: "enterprise",
		TableDesc: "企业",
		Columns: []table.ColumnInfo{{
			Columnname: "id",
			Columndesc: "ID",
			Nullable:   false,
			Primary:    true,
			Columntype: "int8",
		}, {
			Columnname: "enterprise_id",
			Columndesc: "企业ID",
			Nullable:   false,
			Primary:    false,
			Columntype: "text",
		}, {
			Columnname: "create_at",
			Columndesc: "创建时间",
			Nullable:   true,
			Primary:    false,
			Columntype: "timestamptz",
		}},
	}
	fmt.Println(GenerateGraphqlText(tableInfo))
}
