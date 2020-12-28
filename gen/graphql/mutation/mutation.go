package mutation

import (
	"github.com/vektah/gqlparser/v2/ast"
	"gqlgen-generator/gen/graphql/context"
	"gqlgen-generator/gen/graphql/inputs"
	"gqlgen-generator/gen/graphql/scalars"
	"gqlgen-generator/gen/graphql/types"
	"gqlgen-generator/gen/graphql/util"
)

func BuildMutation(ctx *context.GqlBuildContext) (def *ast.Definition) {
	def = &ast.Definition{
		Kind:    ast.Object,
		Name:    "Mutation",
		BuiltIn: false,
	}
	fields := make([]*ast.FieldDefinition, 0)

	// 删除记录
	deleteArguments := []*ast.ArgumentDefinition{
		util.NewNotNullArgument(inputs.TableBoolExpName(ctx), "where", ""),
	}
	deleteMutation := &ast.FieldDefinition{
		Description: "删除记录",
		Name:        "delete_" + ctx.TableName,
		Arguments:   deleteArguments,
		Type:        util.NewType(types.TableMutationResponseName(ctx)),
		Directives:  nil,
		Position:    nil,
	}
	// 根据主键删除记录
	deleteByPkArguments := []*ast.ArgumentDefinition{
		util.NewNotNullArgument(scalars.GetPrimaryKeyType(*ctx.TableInfo), "id", ""),
	}
	deleteByPkMutation := &ast.FieldDefinition{
		Description: "根据主键删除记录",
		Name:        "delete_" + ctx.TableName + "_by_pk",
		Arguments:   deleteByPkArguments,
		Type:        util.NewType(types.TableModelName(ctx.TableName)),
		Directives:  nil,
		Position:    nil,
	}
	// 插入记录
	insertArguments := []*ast.ArgumentDefinition{
		util.NewNotNullListAndItemNotNullArgument(inputs.TableInsertInputName(ctx), "objects", ""),
	}
	insertMutation := &ast.FieldDefinition{
		Description: "插入",
		Name:        "insert_" + ctx.TableName,
		Arguments:   insertArguments,
		Type:        util.NewType(types.TableMutationResponseName(ctx)),
		Directives:  nil,
		Position:    nil,
	}
	// 插入一条记录
	insertOneRecordArguments := []*ast.ArgumentDefinition{
		util.NewNotNullArgument(inputs.TableInsertInputName(ctx), "objects", ""),
	}
	insertOneRecordMutation := &ast.FieldDefinition{
		Description: "插入一条记录",
		Name:        "insert_" + ctx.TableName + "_one",
		Arguments:   insertOneRecordArguments,
		Type:        util.NewType(types.TableModelName(ctx.TableName)),
		Directives:  nil,
		Position:    nil,
	}
	// 更新记录
	updateArguments := []*ast.ArgumentDefinition{
		util.NewArgument(inputs.TableIncInputName(ctx), "_inc", ""),
		util.NewArgument(inputs.TableSetInputName(ctx), "_set", ""),
		util.NewNotNullArgument(inputs.TableBoolExpName(ctx), "where", ""),
	}
	updateMutation := &ast.FieldDefinition{
		Description: "更新",
		Name:        "update_" + ctx.TableName,
		Arguments:   updateArguments,
		Type:        util.NewType(types.TableMutationResponseName(ctx)),
		Directives:  nil,
		Position:    nil,
	}
	// 根据主键更新记录
	updateByPkArguments := []*ast.ArgumentDefinition{
		util.NewArgument(inputs.TableIncInputName(ctx), "_inc", ""),
		util.NewArgument(inputs.TableSetInputName(ctx), "_set", ""),
		util.NewNotNullArgument(scalars.GetPrimaryKeyType(*ctx.TableInfo), "id", ""),
	}
	updateByPkMutation := &ast.FieldDefinition{
		Description: "更新",
		Name:        "update_" + ctx.TableName + "_by_pk",
		Arguments:   updateByPkArguments,
		Type:        util.NewType(types.TableModelName(ctx.TableName)),
		Directives:  nil,
		Position:    nil,
	}

	fields = append(fields, deleteMutation, deleteByPkMutation, insertMutation, insertOneRecordMutation, updateMutation, updateByPkMutation)

	def.Fields = fields
	return
}
