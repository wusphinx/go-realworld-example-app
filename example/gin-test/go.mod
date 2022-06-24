module github.com/wusphinx/go-realworld-example-app/example/gin-test

go 1.16

replace github.com/wusphinx/go-realworld-example-app/example/libs => ../libs

require (
	github.com/gin-gonic/gin v1.8.1
	github.com/stretchr/testify v1.7.5
	github.com/wusphinx/go-realworld-example-app/example/libs v0.0.0-00010101000000-000000000000
)
