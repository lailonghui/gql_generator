package querys

import (
	"github.com/vektah/gqlparser/v2/ast"
	"gqlgen-generator/gen/graphql/context"
	"gqlgen-generator/gen/graphql/enums"
	"gqlgen-generator/gen/graphql/inputs"
	"gqlgen-generator/gen/graphql/scalars"
	"gqlgen-generator/gen/graphql/types"
	"gqlgen-generator/gen/graphql/util"
)

func BuildQuery(ctx *context.GqlBuildContext) (def *ast.Definition) {
	def = &ast.Definition{
		Kind:    ast.Object,
		Name:    "Query",
		BuiltIn: false,
	}
	fields := make([]*ast.FieldDefinition, 0)

	// 列表查询
	listQueryArguments := []*ast.ArgumentDefinition{
		util.NewListAndItemNotNullArgument(enums.TableSelectColumnName(ctx), "distinct_on", ""),
		util.NewArgument(scalars.SCALAR_INT, "limit", ""),
		util.NewArgument(scalars.SCALAR_INT, "offset", ""),
		util.NewListAndItemNotNullArgument(inputs.TableOrderByName(ctx), "order_by", ""),
		util.NewArgument(inputs.TableBoolExpName(ctx), "where", ""),
	}
	listQuery := &ast.FieldDefinition{
		Description: "列表查询",
		Name:        ctx.TableName,
		Arguments:   listQueryArguments,
		Type:        ast.NonNullListType(ast.NonNullNamedType(types.TableModelName(ctx.TableName), nil), nil),
		Directives:  nil,
		Position:    nil,
	}
	// 聚合查询
	aggergateArguments := []*ast.ArgumentDefinition{
		util.NewListAndItemNotNullArgument(enums.TableSelectColumnName(ctx), "distinct_on", ""),
		util.NewArgument(scalars.SCALAR_INT, "limit", ""),
		util.NewArgument(scalars.SCALAR_INT, "offset", ""),
		util.NewListAndItemNotNullArgument(inputs.TableOrderByName(ctx), "order_by", ""),
		util.NewArgument(inputs.TableBoolExpName(ctx), "where", ""),
	}
	aggergateQuery := &ast.FieldDefinition{
		Description: "聚合查询",
		Name:        ctx.TableName + "_aggregate",
		Arguments:   aggergateArguments,
		Type:        util.NewNotNullType(types.TableAggregateName(ctx)),
		Directives:  nil,
		Position:    nil,
	}
	// 主键查询
	pkArguments := []*ast.ArgumentDefinition{
		util.NewNotNullArgument(scalars.GetPrimaryKeyType(*ctx.TableInfo), "id", ""),
	}
	pkQuery := &ast.FieldDefinition{
		Description: "主键查询",
		Name:        ctx.TableName + "_by_pk",
		Arguments:   pkArguments,
		Type:        util.NewNotNullType(types.TableModelName(ctx.TableName)),
		Directives:  nil,
		Position:    nil,
	}
	fields = append(fields, listQuery, aggergateQuery, pkQuery)

	def.Fields = fields
	return
}
