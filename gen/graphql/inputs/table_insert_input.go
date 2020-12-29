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
		if !columnInfo.Nullable {
			continue
		}
	/*	if columnInfo.Primary {
			continue
		}
		if i == 1 {
			tableNameAndId := ctx.TableName + "_id"
			if strings.Contains(tableNameAndId, columnInfo.Columnname) {
				continue
			}
		}*/
		t := util.NewType(scalars.CalculateNamedType(columnInfo))
		//if !columnInfo.Nullable {
		//	t = util.NewNotNullType(scalars.CalculateNamedType(columnInfo))
		//}
		fieldList = append(fieldList, &ast.FieldDefinition{
			Name:         columnInfo.Columnname,
			Arguments:    nil,
			DefaultValue: nil,
			Type:         t,
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
