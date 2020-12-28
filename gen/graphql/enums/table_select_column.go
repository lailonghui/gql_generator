package enums

import (
	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/v2/ast"
	"gqlgen-generator/gen/graphql/context"
)

func buildTableSelectColumn(ctx *context.GqlBuildContext) (def *ast.Definition) {
	return &ast.Definition{
		Kind:        ast.Enum,
		Description: "可选select",
		Name:        TableSelectColumnName(ctx),
		EnumValues:  commonEnumValues(ctx),
		Position:    nil,
		BuiltIn:     false,
	}
}

func TableSelectColumnName(ctx *context.GqlBuildContext) string {
	return strcase.ToCamel(ctx.TableInfo.TableName + "_select_column")
}
