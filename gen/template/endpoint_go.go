package template

var ENDPOINT_GO = `
package {{.Conf.ModuleName}}

import (
	"{{.Conf.ProjectName}}/{{.Conf.ModuleDir}}/graph/generated"
	"{{.Conf.ProjectName}}/{{.Conf.ModuleDir}}/graph/resolver"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
)

func endpoint() *handler.Server {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))
	return srv
}

func GinEndpoint() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	srv := endpoint()

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

`
