package types

import (
	"github.com/vektah/gqlparser/v2/ast"
	"gqlgen-generator/gen/graphql/context"
)

// 返回所有graphql type
func AllTypes(ctx *context.GqlBuildContext) []*ast.Definition {
	var defs []*ast.Definition = []*ast.Definition{
		BuildTableModel(ctx),
		buildTableAggregate(ctx),
		buildTableAggregateFields(ctx),
		buildTableAvgFields(ctx),
		buildTableMaxFields(ctx),
		buildTableMinFields(ctx),
		buildTableStddevFields(ctx),
		buildTableStddevPopFields(ctx),
		buildTableStddevSampFields(ctx),
		buildTableSumFields(ctx),
		buildTableVarPopFields(ctx),
		buildTableVarSampFields(ctx),
		buildTableVarianceFields(ctx),
		buildTableMutationResponse(ctx),
	}
	return defs
}
