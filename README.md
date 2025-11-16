CLEANGO-TEMPLATE
===

A clean, minimal, and production-ready Go project template built with a Clean Architecture mindset.
Designed to help teams and developers start building APIs/Services faster with a well-organized structure.

## Installation & Dependencies
- Initialize a Go module (if not already done):

  ```
  go mod init <module-name>
  ``` 
- Download and install dependencies:

  ```
  go mod tidy
  ```
- Verify Go version:
  ```
  go version
  ```  

## Use
- Run the service
  ```
  go run ./cmd/main.go
  ```
- for test
  ```
  go test ./...
  ```
## Directory Hierarchy
```
|—— api
|    |—— route.go
|—— cmd
|    |—— main.go
|—— go.mod
|—— go.sum
|—— internal
|    |—— configs
|        |—— env.go
|    |—— domain
|        |—— model
|            |—— user_model.go
|    |—— handler
|        |—— user_handler.go
|    |—— repository
|        |—— user_repository.go
|—— pkg
|    |—— enum
|        |—— user_enum.go
|    |—— utils
|        |—— utils.go
|    |—— worker
|        |—— worker.go
```
- cmd/ — The application's entry point (e.g., main.go)
- api/ — Contains the router/routing setup (e.g., integration with frameworks such as Echo, Fiber, Gin, or net/http)
- internal/configs/ — Configuration files and environment/config loaders
- internal/domain/ — The domain layer (entities, domain models, and repository/business rule interfaces)
- internal/handler/ — Adapters for handling HTTP requests (controllers/handlers)
- internal/repository/ — Database connections and persistence implementations
- pkg/ — Reusable code (enums, utility functions, background workers, etc.)

Note: If you want to add a Usecase/Service layer, place it between the domain and handler layers to contain business logic that is independent of transport or persistence.


## References
- [มารู้จักกับ Clean Architecture กันดีกว่า!](https://dev.to/prodigy9/clean-architecture-1o9p)
- [Uncle-bobs-clean-architecture](https://medium.com/pragmatic-programming/uncle-bobs-clean-architecture-3884c6f5ffea)
