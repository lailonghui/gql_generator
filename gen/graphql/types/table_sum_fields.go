package types

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/v2/ast"
	"gqlgen-generator/gen/graphql/context"
	"gqlgen-generator/gen/graphql/scalars"
	"gqlgen-generator/gen/graphql/util"
)

func buildTableSumFields(ctx *context.GqlBuildContext) *ast.Definition {
	def := &ast.Definition{}
	def.Kind = ast.Object
	def.Description = fmt.Sprintf(`aggregate sum on columns of table "%s"`, ctx.TableName)
	def.Name = TableSumFieldsName(ctx)
	// build fields
	fieldList := make([]*ast.FieldDefinition, 0)
	for _, columnInfo := range ctx.Columns {
		t := scalars.CalculateNamedType(columnInfo)
		switch t {
		case scalars.SCALAR_BIGINT, scalars.SCALAR_INT, scalars.SCALAR_FLOAT, scalars.SCALAR_NUMERIC:
			// 支持数值类型
			fieldList = append(fieldList, &ast.FieldDefinition{
				Name:         columnInfo.Columnname,
				Arguments:    nil,
				DefaultValue: nil,
				Type:         util.NewType(t),
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

func TableSumFieldsName(ctx *context.GqlBuildContext) string {
	return strcase.ToCamel(ctx.TableName + "_sum_fields")
}
