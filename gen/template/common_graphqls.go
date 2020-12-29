package template

var COMMON_GRAPHQLS = `
"""
expression to compare columns of type _jsonb. All fields are combined with logical 'AND'.
"""
input JsonbComparisonExp {
	_eq: Jsonb
	_gt: Jsonb
	_gte: Jsonb
	_in: [Jsonb!]
	_is_null: Boolean
	_lt: Jsonb
	_lte: Jsonb
	_neq: Jsonb
	_nin: [Jsonb!]
}
"""
expression to compare columns of type bigint. All fields are combined with logical 'AND'.
"""
input BigintComparisonExp {
	_eq: Bigint
	_gt: Bigint
	_gte: Bigint
	_in: [Bigint!]
	_is_null: Boolean
	_lt: Bigint
	_lte: Bigint
	_neq: Bigint
	_nin: [Bigint!]
}
"""
expression to compare columns of type Boolean. All fields are combined with logical 'AND'.
"""
input BooleanComparisonExp {
	_eq: Boolean
	_gt: Boolean
	_gte: Boolean
	_in: [Boolean!]
	_is_null: Boolean
	_lt: Boolean
	_lte: Boolean
	_neq: Boolean
	_nin: [Boolean!]
}
"""
expression to compare columns of type Int. All fields are combined with logical 'AND'.
"""
input IntComparisonExp {
	_eq: Int
	_gt: Int
	_gte: Int
	_in: [Int!]
	_is_null: Boolean
	_lt: Int
	_lte: Int
	_neq: Int
	_nin: [Int!]
}
"""
expression to compare columns of type Float. All fields are combined with logical 'AND'.
"""
input FloatComparisonExp{
	_eq: Float
	_gt: Float
	_gte: Float
	_in: [Float!]
	_is_null: Boolean
	_lt: Float
	_lte: Float
	_neq: Float
	_nin: [Float!]
}
"""
column ordering options
"""
enum OrderBy {
	"""
	in the ascending order, nulls last
	"""
	asc
	"""
	in the ascending order, nulls first
	"""
	asc_nulls_first
	"""
	in the ascending order, nulls last
	"""
	asc_nulls_last
	"""
	in the descending order, nulls first
	"""
	desc
	"""
	in the descending order, nulls first
	"""
	desc_nulls_first
	"""
	in the descending order, nulls last
	"""
	desc_nulls_last
}
"""
expression to compare columns of type String. All fields are combined with logical 'AND'.
"""
input StringComparisonExp {
	_eq: String
	_gt: String
	_gte: String
	_ilike: String
	_in: [String!]
	_is_null: Boolean
	_like: String
	_lt: String
	_lte: String
	_neq: String
	_nilike: String
	_nin: [String!]
	_nlike: String
	_nsimilar: String
	_similar: String
}
"""
expression to compare columns of type timestamptz. All fields are combined with logical 'AND'.
"""
input TimestamptzComparisonExp {
	_eq: Timestamptz
	_gt: Timestamptz
	_gte: Timestamptz
	_in: [Timestamptz!]
	_is_null: Boolean
	_lt: Timestamptz
	_lte: Timestamptz
	_neq: Timestamptz
	_nin: [Timestamptz!]
}
"""
expression to compare columns of type numeric. All fields are combined with logical 'AND'.
"""
input NumericComparisonExp {
	_eq: Numeric
	_gt: Numeric
	_gte: Numeric
	_in: [Numeric!]
	_is_null: Boolean
	_lt: Numeric
	_lte: Numeric
	_neq: Numeric
	_nin: [Numeric!]
}
"""
expression to compare columns of type point. All fields are combined with logical 'AND'.
"""
input PointComparisonExp {
	_eq: Point
	_gt: Point
	_gte: Point
	_in: [Point!]
	_is_null: Boolean
	_lt: Point
	_lte: Point
	_neq: Point
	_nin: [Point!]
}
scalar Jsonb
scalar Bigint
scalar Timestamptz
scalar Point
scalar Numeric

`
