export GOSUMDB=off
go clean -modcache
go mod tidy
go mod download github.com/my-org-for-test/go-producer
go build
go run main.go
