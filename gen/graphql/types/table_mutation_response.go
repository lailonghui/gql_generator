package types

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/v2/ast"
	"gqlgen-generator/gen/graphql/context"
	"gqlgen-generator/gen/graphql/scalars"
	"gqlgen-generator/gen/graphql/util"
)

func buildTableMutationResponse(ctx *context.GqlBuildContext) *ast.Definition {
	def := &ast.Definition{}
	def.Kind = ast.Object
	def.Description = fmt.Sprintf(`response of any mutation on the table "%s"`, ctx.TableName)
	def.Name = TableMutationResponseName(ctx)
	// build fields
	fieldList := []*ast.FieldDefinition{
		&ast.FieldDefinition{
			Name: "affected_rows",
			Type: util.NewNotNullType(scalars.SCALAR_INT),
		},
		&ast.FieldDefinition{
			Name: "returning",
			Type: util.NewNotNullListAndItemNotNullType(TableModelName(ctx.TableName)),
		},
	}

	def.Fields = fieldList
	return def
}

func TableMutationResponseName(ctx *context.GqlBuildContext) string {
	return strcase.ToCamel(ctx.TableName + "_mutation_response")
}
