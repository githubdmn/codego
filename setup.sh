go mod init github.com/githubdmn/codego

mkdir -p cmd/api
mkdir -p internal/api/handler
mkdir -p internal/api/middleware
mkdir -p internal/api/route
mkdir -p internal/config
mkdir -p internal/model
mkdir -p internal/service
mkdir -p pkg/database
mkdir docs

# Gin framework
go get github.com/gin-gonic/gin

# Ent ORM
#go get entgo.io/ent/dialect@latest
go get -d entgo.io/ent/cmd/ent
go install entgo.io/ent/cmd/ent@latest

# Swagger
go get github.com/swaggo/swag/cmd/swag
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/files

# Config management
go get github.com/spf13/viper

# Logger
go get go.uber.org/zap

go run entgo.io/ent/cmd/ent init User