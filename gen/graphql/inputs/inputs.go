package inputs

import (
	"github.com/vektah/gqlparser/v2/ast"
	"gqlgen-generator/gen/graphql/context"
)

// 返回所有graphql input
func AllInputs(ctx *context.GqlBuildContext) []*ast.Definition {
	var defs []*ast.Definition = []*ast.Definition{
		buildTableOrderBy(ctx),
		buildBoolExp(ctx),
		buildTableInsertInput(ctx),
		buildTableIncInput(ctx),
		buildTableSetInput(ctx),
		//buildTableAggregateOrderBy(ctx),
		//buildTableAvgOrderBy(ctx),
		//buildTableMaxOrderBy(ctx),
		//buildTableMinOrderBy(ctx),
		//buildTableStddevOrderBy(ctx),
		//buildTableStddevPopOrderBy(ctx),
		//buildTableStddevSampOrderBy(ctx),
		//buildTableSumOrderBy(ctx),
		//buildTableVarPopOrderBy(ctx),
		//buildTableVarSampOrderBy(ctx),
		//buildTableVarianceOrderBy(ctx),
	}
	return defs
}
