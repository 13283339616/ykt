module controller

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	vo v0.0.0
	entity v0.0.0
)

replace (
	entity => ../entity
	vo => ../vo
)
