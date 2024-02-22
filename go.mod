module github.com/factotum/moneymaker/user-service

go 1.19

require (
	github.com/go-chi/chi/v5 v5.0.10
	github.com/go-chi/cors v1.2.1
	github.com/golang-migrate/migrate/v4 v4.16.2
	github.com/google/uuid v1.3.1
	github.com/jaydamon/http-toolbox v0.0.0-20230114132444-809dfa8092f7
	github.com/jaydamon/moneymakergocloak v0.0.0-20230916184810-73ab7be6968c
	github.com/lib/pq v1.10.9
	github.com/stretchr/testify v1.8.4
)

require (
	github.com/Nerzal/gocloak/v12 v12.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-resty/resty/v2 v2.7.0 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rabbitmq/amqp091-go v1.8.1 // indirect
	github.com/segmentio/ksuid v1.0.4 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/net v0.15.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/jaydamon/moneymakergocloak => ../moneymakergocloak
)