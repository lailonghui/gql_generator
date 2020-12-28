package inputs

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/v2/ast"
	"gqlgen-generator/gen/graphql/context"
	"gqlgen-generator/gen/graphql/util"
)

func buildTableMinOrderBy(ctx *context.GqlBuildContext) *ast.Definition {
	def := &ast.Definition{}
	def.Kind = ast.InputObject
	def.Description = fmt.Sprintf(`order by min() on columns of table "%s"`, ctx.TableName)
	def.Name = TableMinOrderByName(ctx)
	// build fields
	fieldList := make([]*ast.FieldDefinition, 0)
	for _, columnInfo := range ctx.Columns {

		fieldList = append(fieldList, &ast.FieldDefinition{
			Name:         columnInfo.Columnname,
			Arguments:    nil,
			DefaultValue: nil,
			Type:         util.NewType("OrderBy"),
			Directives:   nil,
			Position:     nil,
		})

	}

	def.Fields = fieldList
	return def
}

func TableMinOrderByName(ctx *context.GqlBuildContext) string {
	return strcase.ToCamel(ctx.TableName + "_min_order_by")
}
