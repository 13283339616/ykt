module five

go 1.15

require (
	github.com/13283339616/apollo v1.0.2
	github.com/13283339616/eureka v1.0.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/panjf2000/ants/v2 v2.4.3
	github.com/ugorji/go v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20201217014255-9d1352758620 // indirect
	golang.org/x/sys v0.0.0-20201214210602-f9fddec55a1e // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	controller v0.0.0
	entity v0.0.0
	vo v0.0.0

)

replace (
	controller => ./controller
	entity => ./entity
	vo => ./vo
)
