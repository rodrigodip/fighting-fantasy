# Clean arch structure propose:

/project-root
│── cmd/
│   └── app/
│       └── main.go
│
│── internal/
│   ├── domain/
    │   │   ├── user.go
│   │   └── order.go
│   │
│   ├── application/
│   │   ├── usecase/
│   │   │   ├── user_usecase.go       # DTOs podem estar aqui
│   │   │   └── order_usecase.go
│   │   ├── service/
│   │   │   ├── auth_service.go
│   │   │   └── notification_service.go
│   │
│   ├── api/                     # Antes era interface/
│   │   ├── http/
│   │   │   ├── user_handler.go
│   │   │   └── order_handler.go
│   │   └── grpc/
│   │
│   └── infrastructure/
│       ├── repository/
│       │   ├── user_repo.go
│       │   └── order_repo.go
│       ├── db/
│       ├── logger/
│       └── externalapi/
│
│── pkg/
│   ├── config/
│   ├── errors/
│   └── middleware/
│
│── go.mod
│── go.sum
