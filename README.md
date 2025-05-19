

### Initial example instructions

go mod init example.com/hello

go build -o out/hello-app  # Compiles the entire module
./out/hello-app           # Output: Hello, Alice!

# Or run directly:
go run main.go

go get github.com/some/package  # Adds dependency to go.mod
go mod tidy                     # Clean up unused dependencies

go test ./...


project/
├── go.mod
├── main.go
└── greeting/
    ├── greeting.go
    └── greeting_test.go


