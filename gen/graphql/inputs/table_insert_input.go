package inputs

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/v2/ast"
	"gqlgen-generator/gen/graphql/context"
	"gqlgen-generator/gen/graphql/scalars"
	"gqlgen-generator/gen/graphql/util"
)

func buildTableInsertInput(ctx *context.GqlBuildContext) *ast.Definition {
	def := &ast.Definition{}
	def.Kind = ast.InputObject
	def.Description = fmt.Sprintf(`input type for inserting data into table "%s"`, ctx.TableName)
	def.Name = TableInsertInputName(ctx)
	// build fields
	fieldList := make([]*ast.FieldDefinition, 0)
	for _, columnInfo := range ctx.Columns {
		fieldList = append(fieldList, &ast.FieldDefinition{
			Name:         columnInfo.Columnname,
			Arguments:    nil,
			DefaultValue: nil,
			Type:         util.NewType(scalars.CalculateNamedType(columnInfo)),
			Directives:   nil,
			Position:     nil,
		})
	}
	def.Fields = fieldList
	return def
}

func TableInsertInputName(ctx *context.GqlBuildContext) string {
	return strcase.ToCamel(ctx.TableName + "_insert_input")
}
