# list all receipts
@help:
  just -l

# run go mod vendor to fetch all dependencies for local build
@vendor:
  go mod vendor

# run go mod tidy to update dependencies for go modules
@tidy:
  go mod tidy

# run all tests with ginkgo
@test:
  ginkgo run -r -cover -coverprofile=coverage.out

# build a binary executable for the project, with race detection
@build:
  go build -race -o ./short-url cmd/short-url/main.go

# a shortcut to connect to the mysql database
@mysql:
  mysql -h 127.0.0.1 -P 3306 -u root -ppassword db

# a shortcut to connect to the redis database
@redis:
  redis-cli

# run lint
@lint:
  golangci-lint run

# generate RESTful API documentation with Swagger 2.0
@swag-init:
  swag init
