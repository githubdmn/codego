

### Initial example instructions

go build -o hello-app  # Compiles the entire module
./hello-app           # Output: Hello, Alice!

# Or run directly:
go run main.go

go get github.com/some/package  # Adds dependency to go.mod
go mod tidy                     # Clean up unused dependencies

go test ./...


hello/
├── go.mod
├── main.go
└── greeting/
    ├── greeting.go
        └── greeting_test.go


