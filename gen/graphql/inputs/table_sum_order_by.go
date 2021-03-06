package inputs

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/v2/ast"
	"gqlgen-generator/gen/graphql/context"
	"gqlgen-generator/gen/graphql/scalars"
	"gqlgen-generator/gen/graphql/util"
)

func buildTableSumOrderBy(ctx *context.GqlBuildContext) *ast.Definition {
	def := &ast.Definition{}
	def.Kind = ast.InputObject
	def.Description = fmt.Sprintf(`order by stddev_samp() on columns of table "%s"`, ctx.TableName)
	def.Name = TableSumOrderByName(ctx)
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
				Type:         util.NewType("OrderBy"),
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

func TableSumOrderByName(ctx *context.GqlBuildContext) string {
	return strcase.ToCamel(ctx.TableName + "_sum_order_by")
}
