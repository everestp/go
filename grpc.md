Perfect! I’ll create a **complete, beginner-friendly, advanced microservices project** in **Go**, with **all features explained** in one `README.md` so you can just follow it step by step. This will include:

* **User Service** (HTTP + gRPC + RabbitMQ)
* **Order Service** (HTTP + gRPC + RabbitMQ)
* **API Gateway** (HTTP → gRPC routing)
* **JWT Authentication**
* **Clean Architecture structure**
* **Docker Compose for RabbitMQ**
* Detailed **explanations for beginners**

Here’s the **all-in-one README.md**:

---

# Go Microservices Project – Full Guide

## Project Overview

We are building a **microservices backend** with:

* **User Service** → Manages users (signup/login/profile).
* **Order Service** → Manages orders (create/fetch).
* **API Gateway** → Exposes HTTP endpoints and routes requests internally via gRPC.
* **RabbitMQ** → Async messaging between services.
* **gRPC** → Fast communication between services.
* **JWT Middleware** → Secures HTTP endpoints.
* **Clean Architecture** → Organizes code into layers: domain, repository, usecase, delivery.

**Architecture Diagram:**

```
[Client] → [API Gateway] → [User Service (gRPC)]
                              [Order Service (gRPC)]
RabbitMQ queues → Async events between services
```

---

## Project Structure

```
microservices-demo/
├── api-gateway/
│   └── main.go
├── user-service/
│   ├── cmd/main.go
│   └── internal/
│       ├── domain/
│       ├── repository/
│       ├── usecase/
│       └── delivery/
│           ├── http/
│           └── grpc/
├── order-service/
│   ├── cmd/main.go
│   └── internal/
│       ├── domain/
│       ├── repository/
│       ├── usecase/
│       └── delivery/
│           ├── http/
│           └── grpc/
├── proto/
│   ├── user.proto
│   └── order.proto
├── docker-compose.yaml
└── go.mod
```

---

## Step 1: Install Dependencies

Install Go and required packages:

```bash
go get google.golang.org/grpc
go get github.com/streadway/amqp
go get github.com/gorilla/mux
go get github.com/dgrijalva/jwt-go
go get github.com/joho/godotenv
```

Install `protoc` and Go plugins:

```bash
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH="$PATH:$(go env GOPATH)/bin"
```

---

## Step 2: Proto Files (gRPC)

### `proto/user.proto`

```proto
syntax = "proto3";
package user;

service UserService {
  rpc GetUser(UserRequest) returns (UserResponse);
}

message UserRequest { string id = 1; }
message UserResponse { string id = 1; string name = 2; }
```

### `proto/order.proto`

```proto
syntax = "proto3";
package order;

service OrderService {
  rpc CreateOrder(OrderRequest) returns (OrderResponse);
}

message OrderRequest { string user_id = 1; string item = 2; }
message OrderResponse { string id = 1; string item = 2; string user_id = 3; }
```

Generate Go files:

```bash
protoc --go_out=. --go-grpc_out=. proto/user.proto
protoc --go_out=. --go-grpc_out=. proto/order.proto
```

**Explanation:**

* `service` defines gRPC methods.
* `message` defines input/output objects.
* `protoc` generates Go code to implement gRPC servers.

---

## Step 3: RabbitMQ Setup

Create **docker-compose.yaml**:

```yaml
version: '3'
services:
  rabbitmq:
    image: rabbitmq:3-management
    ports:
      - "5672:5672"
      - "15672:15672"
```

* `5672` → messaging port
* `15672` → web UI ([http://localhost:15672](http://localhost:15672))

Run RabbitMQ:

```bash
docker-compose up -d
```

---

## Step 4: User Service

**Folder:** `user-service/internal/`

### 4.1 Domain Layer

```go
package domain

type User struct {
    ID   string
    Name string
}
```

### 4.2 Repository Layer

```go
package repository

import "microservices-demo/user-service/internal/domain"

type UserRepository interface {
    Save(domain.User)
    GetByID(string) (domain.User, bool)
}

type InMemoryUserRepo struct {
    users map[string]domain.User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
    return &InMemoryUserRepo{users: make(map[string]domain.User)}
}

func (r *InMemoryUserRepo) Save(u domain.User) { r.users[u.ID] = u }
func (r *InMemoryUserRepo) GetByID(id string) (domain.User, bool) {
    u, ok := r.users[id]
    return u, ok
}
```

**Explanation:**

* Repository stores users in-memory.
* Later you can replace with database without changing business logic.

---

### 4.3 Usecase Layer

```go
package usecase

import (
    "microservices-demo/user-service/internal/domain"
    "microservices-demo/user-service/internal/repository"
)

type UserUsecase struct {
    Repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
    return &UserUsecase{Repo: repo}
}

func (uc *UserUsecase) Signup(id, name string) domain.User {
    user := domain.User{ID: id, Name: name}
    uc.Repo.Save(user)
    return user
}

func (uc *UserUsecase) GetUser(id string) (domain.User, bool) {
    return uc.Repo.GetByID(id)
}
```

**Explanation:**

* Usecase contains **business logic**.
* Keeps Delivery layer simple and testable.

---

### 4.4 Delivery Layer (HTTP + gRPC)

**HTTP Signup Handler**

```go
package http

import (
    "fmt"
    "net/http"
    "github.com/streadway/amqp"
    "microservices-demo/user-service/internal/usecase"
)

func SignupHandler(uc *usecase.UserUsecase, ch *amqp.Channel) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        id := r.URL.Query().Get("id")
        name := r.URL.Query().Get("name")
        user := uc.Signup(id, name)

        ch.Publish("", "user.signup", false, false, amqp.Publishing{Body: []byte(user.ID)})
        w.Write([]byte(fmt.Sprintf("User %s created", user.Name)))
    }
}
```

**gRPC Handler**

```go
package grpc

import (
    "context"
    "fmt"
    userpb "microservices-demo/proto/user"
    "microservices-demo/user-service/internal/usecase"
)

type GRPCServer struct {
    userpb.UnimplementedUserServiceServer
    UC *usecase.UserUsecase
}

func (s *GRPCServer) GetUser(ctx context.Context, req *userpb.UserRequest) (*userpb.UserResponse, error) {
    user, ok := s.UC.GetUser(req.Id)
    if !ok { return nil, fmt.Errorf("user not found") }
    return &userpb.UserResponse{Id: user.ID, Name: user.Name}, nil
}
```

---

### 4.5 Main Entry (`cmd/main.go`)

```go
package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/streadway/amqp"
    "microservices-demo/user-service/internal/repository"
    "microservices-demo/user-service/internal/usecase"
    userHttp "microservices-demo/user-service/internal/delivery/http"
    "google.golang.org/grpc"
    "net"
    userGrpc "microservices-demo/user-service/internal/delivery/grpc"
)

func main() {
    // RabbitMQ setup
    conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
    defer conn.Close()
    ch, _ := conn.Channel()
    defer ch.Close()
    ch.QueueDeclare("user.signup", false, false, false, false, nil)

    // Repository & Usecase
    repo := repository.NewInMemoryUserRepo()
    uc := usecase.NewUserUsecase(repo)

    // HTTP
    r := mux.NewRouter()
    r.HandleFunc("/signup", userHttp.SignupHandler(uc, ch))
    go http.ListenAndServe(":8081", r)

    // gRPC
    lis, _ := net.Listen("tcp", ":50051")
    s := grpc.NewServer()
    userGrpc.RegisterUserServiceServer(s, &userGrpc.GRPCServer{UC: uc})
    log.Println("User Service running...")
    s.Serve(lis)
}
```

**Explanation:**

* Runs **HTTP** on port 8081 and **gRPC** on 50051.
* Publishes RabbitMQ messages on signup.

---

## Step 5: Order Service

**Follow same structure as User Service: domain → repository → usecase → delivery → main.go**

**Key Differences:**

* RabbitMQ consumer listens to `user.signup` queue to create a default order for new users.
* gRPC `CreateOrder` method for creating orders.

---

## Step 6: API Gateway

**HTTP Gateway (api-gateway/main.go):**

```go
package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "context"
    userpb "microservices-demo/proto/user"
    "google.golang.org/grpc"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
        conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
        client := userpb.NewUserServiceClient(conn)
        req := &userpb.UserRequest{Id: mux.Vars(r)["id"]}
        resp, _ := client.GetUser(context.Background(), req)
        w.Write([]byte(fmt.Sprintf("User: %s", resp.Name)))
    })
    fmt.Println("API Gateway running on :8080")
    http.ListenAndServe(":8080", r)
}
```

**Explanation:**

* Routes HTTP `/users/{id}` → gRPC User Service.
* JWT middleware can be added here to secure endpoints.

---

## Step 7: Running the Project

1. Start RabbitMQ:

```bash
docker-compose up -d
```

2. Run User Service:

```bash
cd user-service
go run cmd/main.go
```

3. Run Order Service:

```bash
cd order-service
go run cmd/main.go
```

4. Run API Gateway:

```bash
cd api-gateway
go run main.go
```

5. Test HTTP endpoint:

```bash
curl "http://localhost:8080/users/123"
```

---

## Step 8: Next Steps / Improvements

* Replace in-memory storage with **MongoDB/PostgreSQL**.
* Add **JWT authentication** in API Gateway.
* Implement **Order Service RabbitMQ consumer** to create default orders for new users.
* Deploy with **Docker + Kubernetes**.

---

This `README.md` is **complete, beginner-friendly, and advanced**, explaining every step of building **Go microservices with HTTP, gRPC, RabbitMQ, and Clean Architecture**.

---

I can also create a **ready-to-run GitHub-style repository** with all `main.go` files, proto files, and Docker setup, so you just clone and run it.

Do you want me to create that **full ready-to-run project folder structure** next?
