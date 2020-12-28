package types

import (
	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/v2/ast"
	"gqlgen-generator/gen/graphql/context"
	"gqlgen-generator/gen/graphql/scalars"
)

func BuildTableModel(ctx *context.GqlBuildContext) (def *ast.Definition) {
	def = &ast.Definition{}
	def.Kind = ast.Object
	def.Description = ctx.TableDesc
	def.Name = TableModelName(ctx.TableName)
	// build fields
	fieldList := make([]*ast.FieldDefinition, 0)
	for _, columnInfo := range ctx.Columns {
		fieldList = append(fieldList, &ast.FieldDefinition{
			Description:  columnInfo.Columndesc,
			Name:         columnInfo.Columnname,
			Arguments:    nil,
			DefaultValue: nil,
			Type: &ast.Type{
				NamedType: scalars.CalculateNamedType(columnInfo),
				NonNull:   !columnInfo.Nullable,
				Position:  nil,
			},
			Directives: nil,
			Position:   nil,
		})
	}
	def.Fields = fieldList
	return
}

func TableModelName(tableName string) string {
	return strcase.ToCamel(tableName)
}
