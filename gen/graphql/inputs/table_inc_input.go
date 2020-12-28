package inputs

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/v2/ast"
	"gqlgen-generator/gen/graphql/context"
	"gqlgen-generator/gen/graphql/scalars"
	"gqlgen-generator/gen/graphql/util"
)

func buildTableIncInput(ctx *context.GqlBuildContext) *ast.Definition {
	def := &ast.Definition{}
	def.Kind = ast.InputObject
	def.Description = fmt.Sprintf(`input type for incrementing integer column in table "%s"`, ctx.TableName)
	def.Name = TableIncInputName(ctx)
	// build fields
	fieldList := make([]*ast.FieldDefinition, 0)
	for _, columnInfo := range ctx.Columns {
		switch scalars.CalculateNamedType(columnInfo) {
		case scalars.SCALAR_BIGINT, scalars.SCALAR_INT, scalars.SCALAR_FLOAT, scalars.SCALAR_NUMERIC:
			// 支持数值类型
			fieldList = append(fieldList, &ast.FieldDefinition{
				Name:         columnInfo.Columnname,
				Arguments:    nil,
				DefaultValue: nil,
				Type:         util.NewType(scalars.CalculateNamedType(columnInfo)),
				Directives:   nil,
				Position:     nil,
			})
		default:
			// 不支持其他类型
		}
	}
	def.Fields = fieldList
	return def
}

func TableIncInputName(ctx *context.GqlBuildContext) string {
	return strcase.ToCamel(ctx.TableName + "_inc_input")
}
