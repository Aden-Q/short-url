# list all receipts
@help:
  just -l

# run go mod vendor to fetch all dependencies for local build
@vendor:
  go mod vendor

# run all tests with ginkgo
@test:
  ginkgo run -r

# build a binary executable for the project
@build:
  go build -o short-url

# a shortcut to connect to the mysql database
@mysql:
  mysql -h 127.0.0.1 -P 3306 -u root -p db
