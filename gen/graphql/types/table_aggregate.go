package types

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/v2/ast"
	"gqlgen-generator/gen/graphql/context"
	"gqlgen-generator/gen/graphql/util"
)

func buildTableAggregate(ctx *context.GqlBuildContext) *ast.Definition {
	def := &ast.Definition{}
	def.Kind = ast.Object
	def.Description = fmt.Sprintf(`aggregated selection of "%s"`, ctx.TableName)
	def.Name = TableAggregateName(ctx)
	// build fields
	fieldList := []*ast.FieldDefinition{
		&ast.FieldDefinition{
			Name: "aggregate",
			Type: util.NewType(TableAggregateFieldsName(ctx)),
		},
	}

	def.Fields = fieldList
	return def
}

func TableAggregateName(ctx *context.GqlBuildContext) string {
	return strcase.ToCamel(ctx.TableName + "_aggregate")
}
