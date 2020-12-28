package inputs

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/v2/ast"
	"gqlgen-generator/gen/graphql/context"
	"gqlgen-generator/gen/graphql/util"
	"strings"
)

func buildTableAggregateOrderBy(ctx *context.GqlBuildContext) *ast.Definition {
	def := &ast.Definition{}
	def.Kind = ast.InputObject
	def.Description = fmt.Sprintf(`order by aggregate values of table "%s"`, ctx.TableName)
	def.Name = TableAggregateOrderByName(ctx)
	// build fields
	fieldList := []*ast.FieldDefinition{
		&ast.FieldDefinition{
			Name: "avg",
			Type: util.NewType(TableAvgOrderByName(ctx)),
		},
		&ast.FieldDefinition{
			Name: "count",
			Type: util.NewType("OrderBy"),
		},
		&ast.FieldDefinition{
			Name: "max",
			Type: util.NewType(TableMaxOrderByName(ctx)),
		},
		&ast.FieldDefinition{
			Name: "min",
			Type: util.NewType(TableMinOrderByName(ctx)),
		},
		&ast.FieldDefinition{
			Name: "stddev",
			Type: util.NewType(TableStddevOrderByName(ctx)),
		},
		&ast.FieldDefinition{
			Name: "stddev_pop",
			Type: util.NewType(TableStddevPopOrderByName(ctx)),
		},
		&ast.FieldDefinition{
			Name: "stddev_samp",
			Type: util.NewType(TableStddevSampOrderByName(ctx)),
		},
		&ast.FieldDefinition{
			Name: "sum",
			Type: util.NewType(TableSumOrderByName(ctx)),
		},
		&ast.FieldDefinition{
			Name: "var_pop",
			Type: util.NewType(TableVarPopOrderByName(ctx)),
		},
		&ast.FieldDefinition{
			Name: "var_samp",
			Type: util.NewType(TableVarSampOrderByName(ctx)),
		},
		&ast.FieldDefinition{
			Name: "variance",
			Type: util.NewType(TableVarianceOrderByName(ctx)),
		},
	}

	def.Fields = fieldList
	return def
}

// 获取聚合排序类型
func getAggregateOrderByType(ctx *context.GqlBuildContext, aggregateType string) *ast.Type {
	return util.NewType(strcase.ToCamel(strings.Join([]string{ctx.TableName, aggregateType, "order_by"}, "_")))
}

func TableAggregateOrderByName(ctx *context.GqlBuildContext) string {
	return strcase.ToCamel(ctx.TableName + "_aggregate_order_by")
}
