package template

var GQLGEN_YML = `
# Where are all the schema files located? globs are supported eg  src/**/*.graphqls
schema:
  - graph/graphqls/*.graphqls

# Where should the generated server code go?
exec:
  filename: graph/generated/generated.go
  package: generated

# Uncomment to enable federation
# federation:
#   filename: graph/generated/federation.go
#   package: generated

# Where should any generated models go?
model:
  filename: graph/model/models_gen.go
  package: model

# Where should the resolver implementations go?
resolver:
  layout: follow-schema
  dir: graph/resolver
  package: resolver

# struct_tag: json

# Optional: turn on to use []Thing instead of []*Thing
# omit_slice_element_pointers: false

# Optional: set to speed up generation time by not performing a final validation pass.
# skip_validation: true

# gqlgen will search for any type names in the schema in these go packages
# if they match it will use them, otherwise it will generate them.
autobind:
  - "{{.Conf.ProjectName}}/{{.Conf.ModuleDir}}/model"
  - "{{.Conf.ProjectName}}/pkg/graphql/model"

# This section declares type mapping between the GraphQL and go type systems
#
# The first line in each type will be used as defaults for resolver arguments and
# modelgen, the others will be allowed when binding to fields. Configure them to
# your liking
models:
  ID:
    model:
      - github.com/99designs/gqlgen/graphql.ID
      - github.com/99designs/gqlgen/graphql.Int
      - github.com/99designs/gqlgen/graphql.Int64
      - github.com/99designs/gqlgen/graphql.Int32
  Int:
    model:
      - github.com/99designs/gqlgen/graphql.Int
      - github.com/99designs/gqlgen/graphql.Int64
      - github.com/99designs/gqlgen/graphql.Int32
  Bigint:
    model:
      - github.com/99designs/gqlgen/graphql.Int64
  Numeric:
    model: github.com/99designs/gqlgen/graphql.Float
  Timestamptz:
    model: {{.Conf.ProjectName}}/pkg/graphql/scalar.Timestamptz


`
