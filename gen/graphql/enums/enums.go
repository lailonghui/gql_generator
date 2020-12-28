package enums

import (
	"github.com/vektah/gqlparser/v2/ast"
	"gqlgen-generator/gen/graphql/context"
)

// 返回所有graphql枚举
func AllEnum(ctx *context.GqlBuildContext) []*ast.Definition {
	var defs []*ast.Definition = []*ast.Definition{
		buildTableSelectColumn(ctx),
		//buildOrderBy(ctx),
	}
	return defs
}

func commonEnumValues(ctx *context.GqlBuildContext) (defs []*ast.EnumValueDefinition) {
	defs = make([]*ast.EnumValueDefinition, 0)
	for _, tableInfo := range ctx.TableInfo.Columns {
		def := &ast.EnumValueDefinition{
			Description: tableInfo.Columndesc,
			Name:        tableInfo.Columnname,
		}
		defs = append(defs, def)
	}
	return
}
