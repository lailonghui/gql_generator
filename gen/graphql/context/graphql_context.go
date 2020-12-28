package context

import (
	"github.com/vektah/gqlparser/v2/ast"
	"gqlgen-generator/gen/table"
)

// gql构建上下文
type GqlBuildContext struct {
	*table.TableInfo
	*ast.SchemaDocument
}
