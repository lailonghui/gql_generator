package graphql

import (
	"bytes"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/formatter"
	"gqlgen-generator/gen/graphql/context"
	"gqlgen-generator/gen/graphql/enums"
	"gqlgen-generator/gen/graphql/inputs"
	"gqlgen-generator/gen/graphql/mutation"
	querys2 "gqlgen-generator/gen/graphql/querys"
	"gqlgen-generator/gen/graphql/scalars"
	"gqlgen-generator/gen/graphql/types"
	"gqlgen-generator/gen/graphql/util"
	"gqlgen-generator/gen/table"
)

// 生成graphql文本
func GenerateGraphqlText(tableInfo *table.TableInfo) string {
	gqlBuildContext := &context.GqlBuildContext{
		TableInfo: tableInfo,
	}
	schemaDocument := buildGraphqlSchema(gqlBuildContext)
	var buf bytes.Buffer
	f := formatter.NewFormatter(&buf)
	f.FormatSchemaDocument(schemaDocument)
	return buf.String()
}

// 生成只含有model的graphql文本
func GenerateModelGraphqlText(tableInfo *table.TableInfo) string {
	ctx := &context.GqlBuildContext{
		TableInfo: tableInfo,
	}
	schemaDocument := &ast.SchemaDocument{}
	ctx.SchemaDocument = schemaDocument
	schemaDocument.Definitions = make([]*ast.Definition, 0)
	schemaDocument.Definitions = append(schemaDocument.Definitions, types.BuildTableModel(ctx))
	schemaDocument.Definitions = append(schemaDocument.Definitions, &ast.Definition{
		Kind: ast.Object,
		Name: "Query",
		Fields: []*ast.FieldDefinition{&ast.FieldDefinition{
			Name: "t",
			Type: util.NewType(scalars.SCALAR_INT),
		}},
	})
	var buf bytes.Buffer
	f := formatter.NewFormatter(&buf)
	f.FormatSchemaDocument(schemaDocument)
	return buf.String()
}

// 生成graphql schema对象
func buildGraphqlSchema(ctx *context.GqlBuildContext) (schemaDocument *ast.SchemaDocument) {
	schemaDocument = &ast.SchemaDocument{}
	ctx.SchemaDocument = schemaDocument
	schemaDocument.Definitions = make([]*ast.Definition, 0)
	schemaDocument.Definitions = append(schemaDocument.Definitions, types.AllTypes(ctx)...)
	schemaDocument.Definitions = append(schemaDocument.Definitions, inputs.AllInputs(ctx)...)
	schemaDocument.Definitions = append(schemaDocument.Definitions, enums.AllEnum(ctx)...)

	schemaDocument.Extensions = append([]*ast.Definition{}, querys2.BuildQuery(ctx))
	schemaDocument.Extensions = append(schemaDocument.Extensions, mutation.BuildMutation(ctx))
	return
}
