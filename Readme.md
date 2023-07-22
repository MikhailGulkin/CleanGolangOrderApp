# Order Service

Order Service is a simple Web API.

## Principles and patterns that I use:

- Clean Architecture
- Experiments with DDD
- CQRS (All but one method use half of the CQRS; `GetAllOrdersByUserID` uses the cache)
- Outbox Pattern
- Saga Pattern(Choreography-based)
- Mediator with Query and Command Bus(`MediatR` inspired)

To use Saga pattern need external `CustomerService` or you can simulate events with RabbitMQ admin panel
## Endpoints
#### GET /healthcheck/
```bash
curl -X "GET" "http://localhost:8000/healthcheck/"
```
```json
{"status": "OK"}
```
To create order need some product, address and client.
#### POST /products/
```bash
curl -X "POST" \
  "http://localhost/api/v1/products/" \
  -d '{
    "productID": "ec14f9be-44e3-4ec0-97d3-47fdac6c55c9",
    "price": 100,
    "discount": 10,
    "Description": "string",
    "Name": "string"
}`
```
`Status 204 No Content`

#### POST /address/
```bash
curl -X "POST" \
  "http://localhost/api/v1/address/" \
  -d '{
    "addressID": "7b2b575e-65e7-427c-bb1b-1fe609fc1f6b",
    "buildingNumber": 44,
    "streetName": "string",
    "city": "string",
    "country": "string"
}`
```
`Status 204 No Content`

#### POST /users/
```bash
curl -X "POST" \
  "http://localhost/api/v1/users/" \
  -d '{
    "userID": "4b2b575e-65e7-427c-bb1b-1fe609fc1f6b",
    "username": "string",
    "addressID": "7b2b575e-65e7-427c-bb1b-1fe609fc1f6b"
}`
```
`Status 204 No Content`

#### POST /orders/
```bash
curl -X "POST" \
  "http://localhost/api/v1/orders/" \
  -d '{
    "orderID": "1b2b434e-44e8-424c-bb1b-1fe609fc1f7d",
    "paymentMethod": "string",
    "productsIds": ["ec14f9be-44e3-4ec0-97d3-47fdac6c55c9"],
    "userID": "4b2b575e-65e7-427c-bb1b-1fe609fc1f6b",
    "deliveryAddress": "7b2b575e-65e7-427c-bb1b-1fe609fc1f6b"
}`
```
After creating an order, need to imitate a response from a third-party service
#### RabbitMQ
- CustomerSaga Queue Payload
```bash
{
    "orderID": "1b2b434e-44e8-424c-bb1b-1fe609fc1f7d",
    "orderType": "Approved"
}
```
After publish payload can get orders:
#### GET /orders/
- From DB
```bash
curl -X "GET" "http://localhost:8000/api/v1/orders?limit=10&order=asc"
```
```json
{
    "orders": [
        {
            "orderStatus": "New",
            "paymentMethod": "Online",
            "client": {
                "ClientID": "4b2b575e-65e7-427c-bb1b-1fe609fc1f6b",
                "ClientName": "test"
            },
            "address": "Russia, Moscow, Pushkina, 44",
            "OrderID": "1b2b434e-44e8-424c-bb1b-1fe609fc1f7d",
            "serialNumber": 1,
            "products": [
                {
                    "productID": "ec14f9be-44e3-4ec0-97d3-47fdac6c55c9",
                    "name": "he3444443h1",
                    "actualPrice": 90
                }
            ],
            "createdAt": "2023-07-23T00:27:52.641012+03:00",
            "totalPrice": 90
        }
    ],
    "count": 1,
    "limit": 10,
    "order": "asc"
}
```

#### GET /orders/user/:id
- From Redis Cache
```bash
curl -X "GET" "http://localhost:8000/api/v1/orders/user/4b2b575e-65e7-427c-bb1b-1fe609fc1f6b?limit=10&order=desc"
```
```json
{
    "orders": [
        {
            "orderStatus": "New",
            "paymentMethod": "Online",
            "client": {
                "ClientID": "4b2b575e-65e7-427c-bb1b-1fe609fc1f6b",
                "ClientName": "test"
            },
            "address": "Russia, Moscow, Pushkina, 44",
            "OrderID": "1b2b434e-44e8-424c-bb1b-1fe609fc1f7d",
            "serialNumber": 1,
            "products": [
                {
                    "productID": "ec14f9be-44e3-4ec0-97d3-47fdac6c55c9",
                    "name": "he3444443h1",
                    "actualPrice": 90
                }
            ],
            "createdAt": "2023-07-23T00:32:00.414866435+03:00",
            "totalPrice": 90
        }
    ],
    "count": 1,
    "limit": 10,
    "order": "desc"
}
```
## Dependencies
### Infrastructure
- [Postgres](https://www.postgresql.org/docs/current/index.html) — Database
- [RabbitMQ](https://www.rabbitmq.com/) — The queue used to publish and handle events
- [Redis](https://redis.io/) — For Query cache
- [Docker](https://docs.docker.com/) — For deployment

### Grafana stack
- [Grafana](https://grafana.com/docs/grafana/latest/) — Web view for logs and prometheus info
- [Loki](https://grafana.com/docs/loki/latest/) — A platform to store and query logs
- [Vector.dev](https://vector.dev) — A tool to collect logs and send them to Loki
- [Prometheus](https://prometheus.io/) — A tool to alerting series info

### Golang libs
- [Gin](https://gin-gonic.com/) - Web framework
- [Gorm](https://gorm.io/index.html) - ORM for working with database 
- [go-redis](https://redis.uptrace.dev/) - Cache database for query request
- [cron](https://github.com/robfig/cron) - To implement relay from Outbox pattern
- [fx](https://uber-go.github.io/fx/) - A dependency injection tool for easy initialization and delivery of dependencies
- [zap](https://pkg.go.dev/go.uber.org/zap) - For better logging config
- [amqp091-go](https://github.com/rabbitmq/amqp091-go) - For working with RabbitMQ

### Setup
The first thing to do is to clone the repository:
```bash
git clone https://github.com/MikhailGulkin/TheCleanestGolangOrderApp
cd TheCleanestGolangOrderApp
```
After that create `.env` and `configs/app/prod.toml`. Examples for use lie next to these files.
Run following command
```bash
make up-prod
```

### TODO
- [ ] Refactoring auto-tests
- [ ] Configure CI
- [ ] Create simple `CustomerService`