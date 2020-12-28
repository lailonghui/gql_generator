package enums

import (
	"github.com/vektah/gqlparser/v2/ast"
	"gqlgen-generator/gen/graphql/context"
)

func buildOrderBy(ctx *context.GqlBuildContext) *ast.Definition {
	orderByEnumDefinition := &ast.Definition{
		Kind:        ast.Enum,
		Description: "可选排序",
		Name:        OrderByName(ctx),
		EnumValues: []*ast.EnumValueDefinition{
			&ast.EnumValueDefinition{
				Description: "asc, nulls last",
				Name:        "asc",
			},
			&ast.EnumValueDefinition{
				Description: "asc, nulls first",
				Name:        "asc_nulls_first",
			},
			&ast.EnumValueDefinition{
				Description: "asc, nulls last",
				Name:        "asc_nulls_last",
			},
			&ast.EnumValueDefinition{
				Description: "desc, nulls last",
				Name:        "desc",
			},
			&ast.EnumValueDefinition{
				Description: "desc, nulls first",
				Name:        "desc_nulls_first",
			},
			&ast.EnumValueDefinition{
				Description: "desc, nulls last",
				Name:        "desc_nulls_last",
			},
		},
	}
	return orderByEnumDefinition
}

func OrderByName(ctx *context.GqlBuildContext) string {
	return "OrderBy"
}
