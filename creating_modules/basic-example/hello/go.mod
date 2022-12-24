module hello


go 1.17

require (
	basic_example v1.0.0
)

replace (
	basic_example v1.0.0 => ../basic_example
)