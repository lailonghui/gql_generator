package inputs

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/v2/ast"
	"gqlgen-generator/gen/graphql/context"
	"gqlgen-generator/gen/graphql/scalars"
	"gqlgen-generator/gen/graphql/util"
	"gqlgen-generator/gen/table"
)

func buildBoolExp(ctx *context.GqlBuildContext) *ast.Definition {
	def := &ast.Definition{}
	def.Kind = ast.InputObject
	def.Description = fmt.Sprintf(`Boolean expression to filter rows from the table "blacklist_operation_record". All fields are combined with a logical '%s'.`, ctx.TableName)
	def.Name = TableBoolExpName(ctx)
	// build fields
	fieldList := make([]*ast.FieldDefinition, 0)
	fieldList = append(fieldList, &ast.FieldDefinition{
		Name:       "_and",
		Type:       util.NewListType(TableBoolExpName(ctx)),
		Directives: nil,
		Position:   nil,
	})
	fieldList = append(fieldList, &ast.FieldDefinition{
		Name:       "_not",
		Type:       util.NewType(TableBoolExpName(ctx)),
		Directives: nil,
		Position:   nil,
	})
	fieldList = append(fieldList, &ast.FieldDefinition{
		Name:       "_or",
		Type:       util.NewListType(TableBoolExpName(ctx)),
		Directives: nil,
		Position:   nil,
	})
	for _, columnInfo := range ctx.Columns {
		fieldList = append(fieldList, &ast.FieldDefinition{
			Name:         columnInfo.Columnname,
			Arguments:    nil,
			DefaultValue: nil,
			Type:         getBoolExpFieldType(columnInfo),
			Directives:   nil,
			Position:     nil,
		})
	}
	def.Fields = fieldList
	return def
}

func getBoolExpFieldType(info table.ColumnInfo) *ast.Type {
	return util.NewType(scalars.CalculateNamedType(info) + "ComparisonExp")
}

func TableBoolExpName(ctx *context.GqlBuildContext) string {
	return strcase.ToCamel(ctx.TableName + "_bool_exp")
}
