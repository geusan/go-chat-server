# Chat server

Chat server with go clean arch

## Libraries
- echo server
- ORM: gorm
- Docs: swagger
- testing: mockery
- Socket: gorilla socket

## Commands

Swag
```bash
# swag (Threre is an issue finding gorm.Model dependency)
swag init --parseDependency --parseInternal

# connect http://localhost:8000/swagger/index.html
```

Mockery
```bash
# Create all of mocks (If i comment with mockery)
go generate ./...
```

DB
```bash
# TODO: migration manager(HOW TO?)
go run cli/migration.go
```


## TO DO
- [x] Add middleware tests
- [ ] Make some class for Exception and DTO
