package inputs

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/v2/ast"
	"gqlgen-generator/gen/graphql/context"
	"gqlgen-generator/gen/graphql/util"
)

func buildTableMaxOrderBy(ctx *context.GqlBuildContext) *ast.Definition {
	def := &ast.Definition{}
	def.Kind = ast.InputObject
	def.Description = fmt.Sprintf(`order by max() on columns of table "%s"`, ctx.TableName)
	def.Name = TableMaxOrderByName(ctx)
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

func TableMaxOrderByName(ctx *context.GqlBuildContext) string {
	return strcase.ToCamel(ctx.TableName + "_max_order_by")
}
