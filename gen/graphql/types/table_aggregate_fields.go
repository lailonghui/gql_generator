package types

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/v2/ast"
	"gqlgen-generator/gen/graphql/context"
	"gqlgen-generator/gen/graphql/scalars"
	"gqlgen-generator/gen/graphql/util"
)

func buildTableAggregateFields(ctx *context.GqlBuildContext) *ast.Definition {
	def := &ast.Definition{}
	def.Kind = ast.Object
	def.Description = fmt.Sprintf(`aggregate fields of "%s"`, ctx.TableName)
	def.Name = TableAggregateFieldsName(ctx)
	// build fields
	fieldList := []*ast.FieldDefinition{
		&ast.FieldDefinition{
			Name: "avg",
			Type: util.NewType(TableAvgFieldsName(ctx)),
		},
		&ast.FieldDefinition{
			Name: "count",
			Type: util.NewType(scalars.SCALAR_INT),
		},
		&ast.FieldDefinition{
			Name: "max",
			Type: util.NewType(TableMaxFieldsName(ctx)),
		},
		&ast.FieldDefinition{
			Name: "min",
			Type: util.NewType(TableMinFieldsName(ctx)),
		},
		&ast.FieldDefinition{
			Name: "stddev",
			Type: util.NewType(TableStddevFieldsName(ctx)),
		},
		&ast.FieldDefinition{
			Name: "stddev_pop",
			Type: util.NewType(TableStddevPopFieldsName(ctx)),
		},
		&ast.FieldDefinition{
			Name: "stddev_samp",
			Type: util.NewType(TableStddevSampFieldsName(ctx)),
		},
		&ast.FieldDefinition{
			Name: "sum",
			Type: util.NewType(TableSumFieldsName(ctx)),
		},
		&ast.FieldDefinition{
			Name: "var_pop",
			Type: util.NewType(TableVarPopFieldsName(ctx)),
		},
		&ast.FieldDefinition{
			Name: "var_samp",
			Type: util.NewType(TableVarSampFieldsName(ctx)),
		},
		&ast.FieldDefinition{
			Name: "variance",
			Type: util.NewType(TableVarianceFieldsName(ctx)),
		},
	}

	def.Fields = fieldList
	return def
}

func TableAggregateFieldsName(ctx *context.GqlBuildContext) string {
	return strcase.ToCamel(ctx.TableName + "_aggregate_fields")
}
